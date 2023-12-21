package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// GetNumberOfFilelinkEntries AUTO-GENERATED (do not edit) code from [x_filelinksSCPD],
// based on SOAP action 'GetNumberOfFilelinkEntries', Fritz!Box-System-Version 164.07.57
//
// [x_filelinksSCPD]: http://fritz.box:49000/x_filelinksSCPD.xml
func GetNumberOfFilelinkEntries(session *soap.SoapSession) (tr064model.GetNumberOfFilelinkEntriesResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_filelinks").
		Uri("urn:dslforum-org:service:X_AVM-DE_Filelinks:1").
		Action("GetNumberOfFilelinkEntries").
		Do().Body.Data
	result := tr064model.GetNumberOfFilelinkEntriesResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
