package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetIPInterfaceNumberOfEntries AUTO-GENERATED (do not edit) code from [lanhostconfigmgmSCPD],
// based on SOAP action 'GetIPInterfaceNumberOfEntries', Fritz!Box-System-Version 141.07.57
//
// [lanhostconfigmgmSCPD]: http://fritz.box:49000/lanhostconfigmgmSCPD.xml
func GetIPInterfaceNumberOfEntries(session *soap.SoapSession) (tr064model.GetIPInterfaceNumberOfEntriesResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/lanhostconfigmgm").
		Uri("urn:dslforum-org:service:LANHostConfigManagement:1").
		Action("GetIPInterfaceNumberOfEntries").
		Do().Body.Data
	result := tr064model.GetIPInterfaceNumberOfEntriesResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
