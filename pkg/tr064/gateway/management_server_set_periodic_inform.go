package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetPeriodicInform AUTO-GENERATED (do not edit) code from [mgmsrvSCPD],
// based on SOAP action 'SetPeriodicInform', Fritz!Box-System-Version 164.07.57
//
// [mgmsrvSCPD]: http://fritz.box:49000/mgmsrvSCPD.xml
func SetPeriodicInform(session *soap.SoapSession) (tr064model.SetPeriodicInformResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/mgmsrv").
		Uri("urn:dslforum-org:service:ManagementServer:1").
		Action("SetPeriodicInform").
		Do().Body.Data
	result := tr064model.SetPeriodicInformResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
