package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetUpgradeManagement AUTO-GENERATED (do not edit) code from [mgmsrvSCPD],
// based on SOAP action 'SetUpgradeManagement', Fritz!Box-System-Version 164.08.00
//
// [mgmsrvSCPD]: http://fritz.box:49000/mgmsrvSCPD.xml
func SetUpgradeManagement(session *soap.SoapSession, upgradesManaged bool) (tr064model.SetUpgradeManagementResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/mgmsrv").
		Uri("urn:dslforum-org:service:ManagementServer:1").
		Action("SetUpgradeManagement").
		AddBoolParam("NewUpgradesManaged", upgradesManaged).
		Do()
	if err != nil {
		return tr064model.SetUpgradeManagementResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetUpgradeManagementResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetUpgradeManagementResponse{}, err
	}
	return result, nil
}
