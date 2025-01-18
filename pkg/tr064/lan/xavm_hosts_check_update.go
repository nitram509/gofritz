package lan

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// XavmHostsCheckUpdate AUTO-GENERATED (do not edit) code from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_HostsCheckUpdate', Fritz!Box-System-Version 164.08.00
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
func XavmHostsCheckUpdate(session *soap.SoapSession) (tr064model.XavmHostsCheckUpdateResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/hosts").
		Uri("urn:dslforum-org:service:Hosts:1").
		Action("X_AVM-DE_HostsCheckUpdate").
		Do()
	if err != nil {
		return tr064model.XavmHostsCheckUpdateResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.XavmHostsCheckUpdateResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.XavmHostsCheckUpdateResponse{}, err
	}
	return result, nil
}
