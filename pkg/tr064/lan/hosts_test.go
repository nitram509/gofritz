package lan

import (
	"encoding/xml"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/nitram509/gofitz/pkg/soap"
	"testing"
)

func Test_parsing_XAvmGetSpecificHostEntryByIPResponse(t *testing.T) {
	var envelope soap.SoapResponse
	err := xml.Unmarshal(([]byte)(referenceXAvmGetSpecificHostEntryByIPResponse), &envelope)
	then.AssertThat(t, err, is.Nil())

	response := envelope.Body.XAvmGetSpecificHostEntryByIpResponse
	then.AssertThat(t, response.MACAddress, is.EqualTo("11:22:33:44:55:66"))
	then.AssertThat(t, response.Active, is.False())
	then.AssertThat(t, response.HostName, is.EqualTo("foobar"))
	then.AssertThat(t, response.InterfaceType, is.EqualTo("type"))
	then.AssertThat(t, response.XAvmPort, is.EqualTo("0"))
	then.AssertThat(t, response.XAvmSpeed, is.EqualTo("0"))
	then.AssertThat(t, response.XAvmUpdateAvailable, is.False())
	then.AssertThat(t, response.XAvmUpdateSuccessful, is.EqualTo("unknown"))
	then.AssertThat(t, response.XAvmInfoURL, is.EqualTo("info-url"))
	then.AssertThat(t, response.XAvmMACAddressList, is.EqualTo("mac"))
	then.AssertThat(t, response.XAvmModel, is.EqualTo("model"))
	then.AssertThat(t, response.XAvmURL, is.EqualTo("url"))
	then.AssertThat(t, response.XAvmGuest, is.False())
	then.AssertThat(t, response.XAvmRequestClient, is.False())
	then.AssertThat(t, response.XAvmVPN, is.False())
	then.AssertThat(t, response.XAvmWANAccess, is.EqualTo("granted"))
	then.AssertThat(t, response.XAvmDisallow, is.False())
	then.AssertThat(t, response.XAvmIsMeshable, is.True())
	then.AssertThat(t, response.XAvmPriority, is.False())
	then.AssertThat(t, response.XAvmFriendlyName, is.EqualTo("foobar"))
	then.AssertThat(t, response.XAvmFriendlyNameIsWriteable, is.True())
}

var referenceXAvmGetSpecificHostEntryByIPResponse = `
<?xml version="1.0"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
    <s:Body>
        <u:X_AVM-DE_GetSpecificHostEntryByIPResponse xmlns:u="urn:dslforum-org:service:Hosts:1">
			<NewMACAddress>11:22:33:44:55:66</NewMACAddress>
			<NewActive>0</NewActive>
			<NewHostName>foobar</NewHostName>
			<NewInterfaceType>type</NewInterfaceType>
			<NewX_AVM-DE_Port>0</NewX_AVM-DE_Port>
			<NewX_AVM-DE_Speed>0</NewX_AVM-DE_Speed>
			<NewX_AVM-DE_UpdateAvailable>0</NewX_AVM-DE_UpdateAvailable>
			<NewX_AVM-DE_UpdateSuccessful>unknown</NewX_AVM-DE_UpdateSuccessful>
			<NewX_AVM-DE_InfoURL>info-url</NewX_AVM-DE_InfoURL>
			<NewX_AVM-DE_MACAddressList>mac</NewX_AVM-DE_MACAddressList>
			<NewX_AVM-DE_Model>model</NewX_AVM-DE_Model>
			<NewX_AVM-DE_URL>url</NewX_AVM-DE_URL>
			<NewX_AVM-DE_Guest>0</NewX_AVM-DE_Guest>
			<NewX_AVM-DE_RequestClient>0</NewX_AVM-DE_RequestClient>
			<NewX_AVM-DE_VPN>0</NewX_AVM-DE_VPN>
			<NewX_AVM-DE_WANAccess>granted</NewX_AVM-DE_WANAccess>
			<NewX_AVM-DE_Disallow>0</NewX_AVM-DE_Disallow>
			<NewX_AVM-DE_IsMeshable>1</NewX_AVM-DE_IsMeshable>
			<NewX_AVM-DE_Priority>0</NewX_AVM-DE_Priority>
			<NewX_AVM-DE_FriendlyName>foobar</NewX_AVM-DE_FriendlyName>
			<NewX_AVM-DE_FriendlyNameIsWriteable>1</NewX_AVM-DE_FriendlyNameIsWriteable>
        </u:X_AVM-DE_GetSpecificHostEntryByIPResponse>
    </s:Body>
</s:Envelope>
`
