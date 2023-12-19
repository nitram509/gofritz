package tr064

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/nitram509/gofitz/pkg"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type SoapSession struct {
	sid        string
	authDigest string
	authHeader string
}

type envelopeParameter struct {
	Action string
	Uri    string
	Params string
}

type soapResponse struct {
	XMLName       xml.Name `xml:"Envelope"`
	S             string   `xml:"s,attr"`
	EncodingStyle string   `xml:"encodingStyle,attr"`
	Body          struct {
		XAvmDeCreateUrlSIDResponse           XAvmDeCreateUrlSIDResponse           `xml:"X_AVM-DE_CreateUrlSIDResponse,omitempty"`
		XAvmGetSpecificHostEntryByIpResponse XAvmGetSpecificHostEntryByIpResponse `xml:"X_AVM-DE_GetSpecificHostEntryByIPResponse,omitempty"`
	} `xml:"Body"`
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

type ConvertibleBoolean bool

func (bit *ConvertibleBoolean) UnmarshalJSON(data []byte) error {
	asString := string(data)
	if asString == "1" || asString == "true" {
		*bit = true
	} else if asString == "0" || asString == "false" {
		*bit = false
	} else {
		return errors.New(fmt.Sprintf("Boolean unmarshal error: invalid input %s", asString))
	}
	return nil
}

func soapRequest(hostname string, soapAction string, soapUri string, digestAuth string, requestPath string, params []soapParam) (*http.Request, error) {
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
	req.Header.Set("Content-Type", "text/xml")
	req.Header.Set("User-Agent", "gofritz/"+pkg.VERSION)
	if len(digestAuth) > 0 {
		req.Header.Set("Authorization", digestAuth)
	}
	return req, nil
}

/*
Soap Error Example
-----------------------

<?xml version="1.0"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
<s:Body>
<s:Fault>
<faultcode>s:Client</faultcode>
<faultstring>UPnPError</faultstring>
<detail>
<UPnPError xmlns="urn:dslforum-org:control-1-0">
<errorCode>402</errorCode>
<errorDescription>Invalid Args</errorDescription>
</UPnPError>
</detail>
</s:Fault>
</s:Body>
</s:Envelope>
*/

/*
<?xml version="1.0"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
<s:Body>
<s:Fault>
<faultcode>s:Client</faultcode>
<faultstring>UPnPError</faultstring>
<detail>
<UPnPError xmlns="urn:dslforum-org:control-1-0">
<errorCode>714</errorCode>
<errorDescription>NoSuchEntryInArray</errorDescription>
</UPnPError>
</detail>
</s:Fault>
</s:Body>
</s:Envelope>
*/

type SoapAuthenticator interface {
	getAuthHeader() string
}

func (ss SoapSession) getAuthHeader() string {
	return ss.authHeader
}

type ActionCommand interface {
	XAvmGetSpecificHostEntryByIp(ipAddress string) XAvmGetSpecificHostEntryByIpCommand
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

func (ac *actionCmd) addParam(name string, value string) {
	ac.soapActionParams = append(ac.soapActionParams, soapParam{
		name:  name,
		value: value,
	})
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
