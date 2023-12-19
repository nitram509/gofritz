package tr064

import (
	"encoding/xml"
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

func Login() *SoapSession {
	action := "X_AVM-DE_CreateUrlSID"
	uri := "urn:dslforum-org:service:DeviceConfig:1"
	req, err := soapRequest("fritz.box", action, uri, "", "/upnp/control/deviceconfig", nil)
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
	digest := createAuthenticationDigest(username, password, authHeader, method, url)

	req, err = soapRequest("fritz.box", action, uri, digest, "/upnp/control/deviceconfig", nil)
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

	var envResp soapResponse
	err = xml.Unmarshal(buf, &envResp)

	sid := strings.TrimLeft(envResp.Body.XAvmDeCreateUrlSIDResponse.UrlSID, "sid=")
	return &SoapSession{
		sid:        sid,
		authDigest: digest,
		authHeader: authHeader,
	}
}
