package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// SetDVBCEnable AUTO-GENERATED (do not edit) code from [x_mediaSCPD],
// based on SOAP action 'SetDVBCEnable', Fritz!Box-System-Version 141.07.57
//
// [x_mediaSCPD]: http://fritz.box:49000/x_mediaSCPD.xml
func SetDVBCEnable(session *soap.SoapSession, dvbcEnabled bool) (tr064model.SetDVBCEnableResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_media").
		Uri("urn:dslforum-org:service:X_AVM-DE_Media:1").
		Action("SetDVBCEnable").
		AddBoolParam("NewDVBCEnabled", dvbcEnabled).
		Do().Body.Data
	result := tr064model.SetDVBCEnableResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
