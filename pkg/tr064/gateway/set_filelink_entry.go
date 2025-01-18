package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetFilelinkEntry AUTO-GENERATED (do not edit) code from [x_filelinksSCPD],
// based on SOAP action 'SetFilelinkEntry', Fritz!Box-System-Version 164.08.00
//
// [x_filelinksSCPD]: http://fritz.box:49000/x_filelinksSCPD.xml
func SetFilelinkEntry(session *soap.SoapSession, id string, accessCountLimit int, expire int) (tr064model.SetFilelinkEntryResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_filelinks").
		Uri("urn:dslforum-org:service:X_AVM-DE_Filelinks:1").
		Action("SetFilelinkEntry").
		AddStringParam("NewID", id).
		AddIntParam("NewAccessCountLimit", accessCountLimit).
		AddIntParam("NewExpire", expire).
		Do()
	if err != nil {
		return tr064model.SetFilelinkEntryResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetFilelinkEntryResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetFilelinkEntryResponse{}, err
	}
	return result, nil
}
