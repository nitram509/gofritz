package tr064

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type XAvmGetSpecificHostEntryByIpCommand interface {
	Do() XAvmGetSpecificHostEntryByIpResponse
}

type xAvmGetSpecificHostEntryByIpCommand struct {
	actionCmd
}

type XAvmGetSpecificHostEntryByIpResponse struct {
	MACAddress                  string `xml:"NewMACAddress"`
	Active                      string `xml:"NewActive"`
	HostName                    string `xml:"NewHostName"`
	InterfaceType               string `xml:"NewInterfaceType"`
	XAvmPort                    string `xml:"NewX_AVM-DE_Port"`
	XAvmSpeed                   string `xml:"NewX_AVM-DE_Speed"`
	XAvmUpdateAvailable         bool   `xml:"NewX_AVM-DE_UpdateAvailable"`
	XAvmUpdateSuccessful        string `xml:"NewX_AVM-DE_UpdateSuccessful"`
	XAvmInfoURL                 string `xml:"NewX_AVM-DE_InfoURL"`
	XAvmMACAddressList          string `xml:"NewX_AVM-DE_MACAddressList"`
	XAvmModel                   string `xml:"NewX_AVM-DE_Model"`
	XAvmURL                     string `xml:"NewX_AVM-DE_URL"`
	XAvmGuest                   bool   `xml:"NewX_AVM-DE_Guest"`
	XAvmRequestClient           bool   `xml:"NewX_AVM-DE_RequestClient"`
	XAvmVPN                     bool   `xml:"NewX_AVM-DE_VPN"`
	XAvmWANAccess               string `xml:"NewX_AVM-DE_WANAccess"`
	XAvmDisallow                bool   `xml:"NewX_AVM-DE_Disallow"`
	XAvmIsMeshable              bool   `xml:"NewX_AVM-DE_IsMeshable"`
	XAvmPriority                bool   `xml:"NewX_AVM-DE_Priority"`
	XAvmFriendlyName            string `xml:"NewX_AVM-DE_FriendlyName"`
	XAvmFriendlyNameIsWriteable bool   `xml:"NewX_AVM-DE_FriendlyNameIsWriteable"`
}

func (ac *actionCmd) XAvmGetSpecificHostEntryByIp(ipAddress string) XAvmGetSpecificHostEntryByIpCommand {
	ac.addParam("NewIPAddress", ipAddress)
	return xAvmGetSpecificHostEntryByIpCommand{actionCmd: *ac}
}

func (cmd xAvmGetSpecificHostEntryByIpCommand) Do() XAvmGetSpecificHostEntryByIpResponse {

	//cmd.authenticator.createDigest()
	username := os.Getenv("FB_USERNAME")
	password := os.Getenv("FB_PASSWORD")
	url := "http://fritz.box:49000" + cmd.soapCommand.reqPath
	digest := createAuthenticationDigest(username, password, cmd.authenticator.getAuthHeader(), "POST", url)

	req, err := soapRequest("fritz.box",
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
	envResp := soapResponse{}
	err = xml.Unmarshal(resp, &envResp)
	if err != nil {
		panic(err)
	}
	return envResp.Body.XAvmGetSpecificHostEntryByIpResponse
}

func XAvmGetSpecificHostEntryByIp(soap SoapSession, ipAddress string) string {
	response := NewSoapRequest().
		WithAuthenticator(soap).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_GetSpecificHostEntryByIp").
		XAvmGetSpecificHostEntryByIp(ipAddress).
		Do()
	return response.HostName
}
