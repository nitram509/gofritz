package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetUpgradeManagement AUTO-GENERATED (do not edit) code from [mgmsrvSCPD],
// based on SOAP action 'SetUpgradeManagement', Fritz!Box-System-Version 164.07.57
//
// [mgmsrvSCPD]: http://fritz.box:49000/mgmsrvSCPD.xml
func SetUpgradeManagement(session *soap.SoapSession, upgradesManaged bool) (tr064model.SetUpgradeManagementResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/mgmsrv").
		Uri("urn:dslforum-org:service:ManagementServer:1").
		Action("SetUpgradeManagement").
		AddBoolParam("NewUpgradesManaged", upgradesManaged).
		Do().Body.Data
	result := tr064model.SetUpgradeManagementResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
