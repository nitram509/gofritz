package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetPeriodicInform AUTO-GENERATED (do not edit) code from [mgmsrvSCPD],
// based on SOAP action 'SetPeriodicInform', Fritz!Box-System-Version 164.07.57
//
// [mgmsrvSCPD]: http://fritz.box:49000/mgmsrvSCPD.xml
func SetPeriodicInform(session *soap.SoapSession, periodicInformEnable bool, periodicInformInterval int, periodicInformTime string) (tr064model.SetPeriodicInformResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/mgmsrv").
		Uri("urn:dslforum-org:service:ManagementServer:1").
		Action("SetPeriodicInform").
		AddBoolParam("NewPeriodicInformEnable", periodicInformEnable).
		AddIntParam("NewPeriodicInformInterval", periodicInformInterval).
		AddStringParam("NewPeriodicInformTime", periodicInformTime).
		Do()
	if err != nil {
		return tr064model.SetPeriodicInformResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetPeriodicInformResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetPeriodicInformResponse{}, err
	}
	return result, nil
}
