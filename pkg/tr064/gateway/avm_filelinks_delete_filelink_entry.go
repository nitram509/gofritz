package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// DeleteFilelinkEntry AUTO-GENERATED (do not edit) code from [x_filelinksSCPD],
// based on SOAP action 'DeleteFilelinkEntry', Fritz!Box-System-Version 164.07.57
//
// [x_filelinksSCPD]: http://fritz.box:49000/x_filelinksSCPD.xml
func DeleteFilelinkEntry(session *soap.SoapSession) (tr064model.DeleteFilelinkEntryResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_filelinks").
		Uri("urn:dslforum-org:service:X_AVM-DE_Filelinks:1").
		Action("DeleteFilelinkEntry").
		Do().Body.Data
	result := tr064model.DeleteFilelinkEntryResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
