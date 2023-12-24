package soap

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/nitram509/gofitz/pkg"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const httpMethod = "POST"

type GenericSoapResponse struct {
	XMLName       xml.Name `xml:"Envelope"`
	S             string   `xml:"s,attr"`
	EncodingStyle string   `xml:"encodingStyle,attr"`
	Body          struct {
		Data []byte `xml:",innerxml"`
	} `xml:"Body,omitempty"`
}

type soapErrorResponse struct {
	XMLName       xml.Name `xml:"Envelope"`
	S             string   `xml:"s,attr"`
	EncodingStyle string   `xml:"encodingStyle,attr"`
	Body          struct {
		Fault struct {
			Faultcode   string `xml:"faultcode"`
			Faultstring string `xml:"faultstring"`
			Detail      struct {
				UPnPError struct {
					ErrorCode        string `xml:"errorCode"`
					ErrorDescription string `xml:"errorDescription"`
				} `xml:"UPnPError"`
			} `xml:"detail"`
		} `xml:"Fault"`
	} `xml:"Body"`
}

func newHttpRequest(requestUrl string, digestAuth string, soapAction string, soapUri string, params []soapParam) *http.Request {
	sb := strings.Builder{}
	if params != nil {
		for _, param := range params {
			sb.Write([]byte(fmt.Sprintf("<%s>%s</%s>", param.name, param.value, param.name)))
		}
	}
	envelope := `<?xml version='1.0' encoding='utf-8'?>
<s:Envelope s:encodingStyle='http://schemas.xmlsoap.org/soap/encoding/' xmlns:s='http://schemas.xmlsoap.org/soap/envelope/'>
<s:Body>
<u:` + soapAction + ` xmlns:u='` + soapUri + `'>` + sb.String() + `</u:` + soapAction + `>
</s:Body>
</s:Envelope>`

	req, err := http.NewRequest(httpMethod, requestUrl, strings.NewReader(envelope))
	if err != nil {
		panic(err)
	}
	req.Header.Set("SOAPACTION", soapUri+"#"+soapAction)
	req.Header.Set("Accept", "text/xml")
	req.Header.Set("Content-Type", "text/xml")
	req.Header.Set("User-Agent", "gofritz/"+pkg.VERSION)
	if len(digestAuth) > 0 {
		req.Header.Set("Authorization", digestAuth)
	}
	return req
}

func (cmd *soapCmd) Do() GenericSoapResponse {
	fullUrl := cmd.session.BaseUrl() + cmd.reqPath
	client := &http.Client{}

	var response *http.Response
	var respData []byte
	var err error
	maxRetries := 1
	for unauthorized := true; unauthorized; maxRetries-- {
		digest := cmd.session.computeDigestResponse(httpMethod, fullUrl)
		req := newHttpRequest(fullUrl, digest, cmd.soapAction, cmd.soapUri, cmd.soapActionParams)
		response, err = client.Do(req)
		if err != nil {
			panic(err)
		}
		respData, err = io.ReadAll(response.Body)
		unauthorized = response.StatusCode == 401
		if unauthorized {
			cmd.session.reqAuthDigest = response.Header.Get("WWW-Authenticate")
		} // else, there is no indicator for how long the session is valid ...
		if unauthorized && maxRetries <= 0 {
			panic(errors.New(fmt.Sprintf("can't authenticate '%s' with user '%s'", cmd.session.BaseUrl(), cmd.session.username)))
		}
	}

	if response.StatusCode != 200 {
		if response.StatusCode == 500 {
			upnpError := soapErrorResponse{}
			err := xml.Unmarshal(respData, &upnpError)
			if err != nil {
				panic(err)
			}
			err = errors.New(fmt.Sprintf("%s, soap_action=%s, error_code=%s, error_description=%s",
				upnpError.Body.Fault.Faultstring,
				cmd.soapAction,
				upnpError.Body.Fault.Detail.UPnPError.ErrorCode,
				upnpError.Body.Fault.Detail.UPnPError.ErrorDescription,
			))
			panic(err)
		} else {
			panic(errors.New("Error calling '" + fullUrl + "', Status:" + response.Status))
		}
	}

	envResp := GenericSoapResponse{}
	err = xml.Unmarshal(respData, &envResp)
	if err != nil {
		panic(err)
	}
	return envResp
}

type soapParam struct {
	name  string
	value string
}

func (cmd *soapCmd) AddStringParam(name string, value string) SoapCommand {
	escapedValue := strings.Builder{}
	err := xml.EscapeText(&escapedValue, []byte(value))
	if err != nil {
		panic(err)
	}
	cmd.soapActionParams = append(cmd.soapActionParams, soapParam{
		name:  name,
		value: escapedValue.String(),
	})
	return cmd
}

func (cmd *soapCmd) AddIntParam(name string, value int) SoapCommand {
	cmd.AddStringParam(name, strconv.Itoa(value))
	return cmd
}

func (cmd *soapCmd) AddBoolParam(name string, value bool) SoapCommand {
	if value {
		cmd.AddStringParam(name, "1")
	} else {
		cmd.AddStringParam(name, "0")
	}
	return cmd
}

type SoapSession struct {
	hostname      string
	username      string
	password      string
	reqAuthDigest string
}

// NewSession for a given host name/ip (typical "fritz.box" or "192.168.178.1") with basic credentials
func NewSession(host string, username string, password string) *SoapSession {
	return &SoapSession{
		hostname: host,
		username: username,
		password: password,
	}
}

// BaseUrl gives the valid URL for the given host with port number and no trailing slash
// (typical http://fritz.box:49000)
func (ss *SoapSession) BaseUrl() string {
	return "http://" + ss.hostname + ":49000"
}

func (ss *SoapSession) computeDigestResponse(httpMethod string, fullUrl string) string {
	if len(ss.reqAuthDigest) == 0 {
		return ""
	}
	return createAuthenticationDigestResponse(ss.username, ss.password, ss.reqAuthDigest, httpMethod, fullUrl)
}

func NewSoapRequest(session *SoapSession) SoapCommand {
	return &soapCmd{
		session: session,
	}
}

type SoapCommand interface {
	Action(string) SoapCommand
	Uri(action string) SoapCommand
	ReqPath(reqPath string) SoapCommand
	AddStringParam(name string, value string) SoapCommand
	AddIntParam(name string, value int) SoapCommand
	AddBoolParam(name string, value bool) SoapCommand
	Do() GenericSoapResponse
}

type soapCmd struct {
	session          *SoapSession
	soapUri          string
	reqPath          string
	soapAction       string
	soapActionParams []soapParam
}

func (cmd *soapCmd) Action(action string) SoapCommand {
	cmd.soapAction = action
	return cmd
}

func (cmd *soapCmd) Uri(uri string) SoapCommand {
	cmd.soapUri = uri
	return cmd
}

func (cmd *soapCmd) ReqPath(reqPath string) SoapCommand {
	cmd.reqPath = reqPath
	return cmd
}
