package tr064

import (
	"bytes"
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type SoapSession struct {
	sid        string
	authDigest string
}

type envelopeParameter struct {
	Action string
	Uri    string
}

type envelopeResponse struct {
	XMLName       xml.Name `xml:"Envelope"`
	S             string   `xml:"s,attr"`
	EncodingStyle string   `xml:"encodingStyle,attr"`
	Body          struct {
		XAvmDeCreateUrlSIDResponse struct {
			U      string `xml:"u,attr"`
			UrlSID string `xml:"NewX_AVM-DE_UrlSID"`
		} `xml:"X_AVM-DE_CreateUrlSIDResponse"`
	} `xml:"Body"`
}

func soapRequest(hostname string, soapAction string, soapUri string, digestAuth string, requestPath string) (*http.Request, error) {
	soapRequestBody, err := template.New("soapEnvelope").Parse(`<?xml version='1.0' encoding='utf-8'?>
<s:Envelope s:encodingStyle='http://schemas.xmlsoap.org/soap/encoding/' xmlns:s='http://schemas.xmlsoap.org/soap/envelope/'>
<s:Body>
<u:{{ .Action }} xmlns:u='{{ .Uri }}'></u:{{ .Action }}>
</s:Body>
</s:Envelope>`)
	if err != nil {
		log.Fatal("Can't parse SOAP envelope template")
	}
	buf := new(bytes.Buffer)
	err = soapRequestBody.Execute(buf, envelopeParameter{
		Action: soapAction,
		Uri:    soapUri,
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
