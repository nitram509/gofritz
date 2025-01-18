package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// DeleteFilelinkEntry AUTO-GENERATED (do not edit) code from [x_filelinksSCPD],
// based on SOAP action 'DeleteFilelinkEntry', Fritz!Box-System-Version 164.08.00
//
// [x_filelinksSCPD]: http://fritz.box:49000/x_filelinksSCPD.xml
func DeleteFilelinkEntry(session *soap.SoapSession, id string) (tr064model.DeleteFilelinkEntryResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_filelinks").
		Uri("urn:dslforum-org:service:X_AVM-DE_Filelinks:1").
		Action("DeleteFilelinkEntry").
		AddStringParam("NewID", id).
		Do()
	if err != nil {
		return tr064model.DeleteFilelinkEntryResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.DeleteFilelinkEntryResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.DeleteFilelinkEntryResponse{}, err
	}
	return result, nil
}
