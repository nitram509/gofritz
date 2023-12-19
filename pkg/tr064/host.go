package tr064

import (
	"errors"
	"io"
	"net/http"
)

func (ss *SoapSession) GetSpecificHostEntryByIp(ipAddress string) string {
	action := "X_AVM-DE_GetSpecificHostEntryByIp"
	uri := "urn:dslforum-org:service:Hosts:1"
	reqPath := "/upnp/control/hosts"
	req, err := soapRequest("fritz.box", action, uri, ss.authDigest, reqPath)
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
		panic(errors.New("can't call GetSpecificHostEntryByIp()"))
	}

	if err != nil {
		panic(err)
	}

	return string(resp)
}
