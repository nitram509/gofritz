package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// GetSpecificFilelinkEntry AUTO-GENERATED (do not edit) code from [x_filelinksSCPD],
// based on SOAP action 'GetSpecificFilelinkEntry', Fritz!Box-System-Version 164.07.57
//
// [x_filelinksSCPD]: http://fritz.box:49000/x_filelinksSCPD.xml
func GetSpecificFilelinkEntry(session *soap.SoapSession, id string) (tr064model.GetSpecificFilelinkEntryResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_filelinks").
		Uri("urn:dslforum-org:service:X_AVM-DE_Filelinks:1").
		Action("GetSpecificFilelinkEntry").
		AddStringParam("NewID", id).
		Do()
	if err != nil {
		return tr064model.GetSpecificFilelinkEntryResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.GetSpecificFilelinkEntryResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.GetSpecificFilelinkEntryResponse{}, err
	}
	return result, nil
}
