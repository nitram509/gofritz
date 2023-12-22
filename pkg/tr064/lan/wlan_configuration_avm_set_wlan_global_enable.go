package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

// XavmSetWLANGlobalEnable AUTO-GENERATED (do not edit) code from [wlanconfigSCPD],
// based on SOAP action 'X_AVM-DE_SetWLANGlobalEnable', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
func XavmSetWLANGlobalEnable(session *soap.SoapSession, avmWlanGlobalEnable bool) (tr064model.XavmSetWLANGlobalEnableResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/wlanconfig3").
		Uri("urn:dslforum-org:service:WLANConfiguration:3").
		Action("X_AVM-DE_SetWLANGlobalEnable").
		AddBoolParam("NewX_AVM-DE_WLANGlobalEnable", avmWlanGlobalEnable).
		Do().Body.Data
	result := tr064model.XavmSetWLANGlobalEnableResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
