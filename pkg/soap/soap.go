package soap

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/nitram509/gofitz/pkg"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type SoapSession struct {
	Sid        string
	AuthDigest string
	AuthHeader string
}

type envelopeParameter struct {
	Action string
	Uri    string
	Params string
}

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

func SoapRequest(hostname string, soapAction string, soapUri string, digestAuth string, requestPath string, params []soapParam) (*http.Request, error) {
	soapRequestBody, err := template.New("soapEnvelope").Parse(`<?xml version='1.0' encoding='utf-8'?>
<s:Envelope s:encodingStyle='http://schemas.xmlsoap.org/soap/encoding/' xmlns:s='http://schemas.xmlsoap.org/soap/envelope/'>
<s:Body>
<u:{{ .Action }} xmlns:u='{{ .Uri }}'>{{ .Params }}</u:{{ .Action }}>
</s:Body>
</s:Envelope>`)
	if err != nil {
		log.Fatal("Can't parse SOAP envelope template")
	}
	buf := new(bytes.Buffer)
	sb := strings.Builder{}
	if params != nil {
		for _, param := range params {
			sb.Write([]byte(fmt.Sprintf("<%s>%s</%s>", param.name, param.value, param.name)))
		}
	}
	err = soapRequestBody.Execute(buf, envelopeParameter{
		Action: soapAction,
		Uri:    soapUri,
		Params: sb.String(),
	})
	if err != nil {
		log.Fatal("Can't parse SOAP Envelope Template with parameters: \n" + err.Error())
	}
	envelope := buf.String()
	requestUrl := "http://" + hostname + ":49000" + requestPath

	req, err := http.NewRequest("POST", requestUrl, strings.NewReader(envelope))
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
	return req, nil
}

type SoapAuthenticator interface {
	getAuthHeader() string
}

func (ss SoapSession) getAuthHeader() string {
	return ss.AuthHeader
}

type ActionCommand interface {
	AddParam(name string, value string) ActionCommand
	Do() GenericSoapResponse
}

type soapParam struct {
	name  string
	value string
}

type actionCmd struct {
	soapCommand      soapCmd
	soapActionParams []soapParam
	soapAction       string
	authenticator    SoapAuthenticator
}

func (cmd *actionCmd) AddParam(name string, value string) ActionCommand {
	cmd.soapActionParams = append(cmd.soapActionParams, soapParam{
		name:  name,
		value: value,
	})
	return cmd
}

type SoapCommand interface {
	Action(string) ActionCommand
	Uri(action string) SoapCommand
	ReqPath(reqPath string) SoapCommand
	WithAuthenticator(authenticator SoapAuthenticator) SoapCommand
}

func (c soapCmd) Action(action string) ActionCommand {
	return &actionCmd{
		soapCommand:   c,
		soapAction:    action,
		authenticator: c.authenticator,
	}
}

func (cmd *actionCmd) Do() GenericSoapResponse {

	//cmd.authenticator.createDigest()
	username := os.Getenv("FB_USERNAME")
	password := os.Getenv("FB_PASSWORD")
	url := "http://fritz.box:49000" + cmd.soapCommand.reqPath
	digest := CreateAuthenticationDigest(username, password, cmd.authenticator.getAuthHeader(), "POST", url)

	req, err := SoapRequest("fritz.box",
		cmd.soapAction,
		cmd.soapCommand.uri,
		digest,
		cmd.soapCommand.reqPath,
		cmd.soapActionParams)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	resp, err := io.ReadAll(response.Body)

	if response.StatusCode != 200 {
		if response.StatusCode == 500 {
			upnpError := soapErrorResponse{}
			err := xml.Unmarshal(resp, &upnpError)
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
			panic(errors.New("Error calling '" + url + "', Status:" + response.Status))
		}
	}

	if err != nil {
		panic(err)
	}
	envResp := GenericSoapResponse{}
	err = xml.Unmarshal(resp, &envResp)
	if err != nil {
		panic(err)
	}
	return envResp
}

func (c soapCmd) Uri(uri string) SoapCommand {
	c.uri = uri
	return c
}

func (c soapCmd) ReqPath(reqPath string) SoapCommand {
	c.reqPath = reqPath
	return c
}

func (c soapCmd) WithAuthenticator(authenticator SoapAuthenticator) SoapCommand {
	c.authenticator = authenticator
	return c
}

type soapCmd struct {
	uri           string
	reqPath       string
	authenticator SoapAuthenticator
}

func NewSoapRequest() SoapCommand {
	return soapCmd{}
}
