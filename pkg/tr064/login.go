package tr064

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type XAvmDeCreateUrlSIDResponse struct {
	U      string `xml:"u,attr"`
	UrlSID string `xml:"NewX_AVM-DE_UrlSID"`
}

func Login() *soap.SoapSession {
	action := "X_AVM-DE_CreateUrlSID"
	uri := "urn:dslforum-org:service:DeviceConfig:1"
	req, err := soap.SoapRequest("fritz.box", action, uri, "", "/upnp/control/deviceconfig", nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if response.StatusCode != 401 {
		log.Fatal("Expect to be unauthorized. Did you have enabled security or is TR064 protocol disabled? \n" /*+ res.Body.()*/)
	}

	username := os.Getenv("FB_USERNAME")
	password := os.Getenv("FB_PASSWORD")
	authHeader := response.Header.Get("WWW-Authenticate")
	method := response.Request.Method
	url := response.Request.URL.String()
	digest := soap.CreateAuthenticationDigest(username, password, authHeader, method, url)

	req, err = soap.SoapRequest("fritz.box", action, uri, digest, "/upnp/control/deviceconfig", nil)
	if err != nil {
		panic(err)
	}
	response, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	if response.StatusCode != 200 {
		log.Fatal("Session create didn't work. \n" /*+ response.String()*/)
	}

	buf, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var envResp soap.SoapResponse
	err = xml.Unmarshal(buf, &envResp)

	sid := strings.TrimLeft(envResp.Body.XAvmDeCreateUrlSIDResponse.UrlSID, "sid=")
	return &soap.SoapSession{
		Sid:        sid,
		AuthDigest: digest,
		AuthHeader: authHeader,
	}
}
