package tr064

import (
	"encoding/xml"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"testing"
)

func Test_errorResponse(t *testing.T) {
	var resp soapErrorResponse
	err := xml.Unmarshal(([]byte)(referenceSoapErrorResponse), &resp)
	then.AssertThat(t, err, is.Nil())

	then.AssertThat(t, resp.Body.Fault.Faultstring, is.EqualTo("UPnPError"))
	then.AssertThat(t, resp.Body.Fault.Faultcode, is.EqualTo("s:Client"))
	then.AssertThat(t, resp.Body.Fault.Detail.UPnPError.ErrorCode, is.EqualTo("714"))
	then.AssertThat(t, resp.Body.Fault.Detail.UPnPError.ErrorDescription, is.EqualTo("NoSuchEntryInArray"))
}

var referenceSoapErrorResponse = `
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
`
