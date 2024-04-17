package gateway

import (
	"encoding/xml"

	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// SetDVBCEnable AUTO-GENERATED (do not edit) code from [x_mediaSCPD],
// based on SOAP action 'SetDVBCEnable', Fritz!Box-System-Version 141.07.57
//
// [x_mediaSCPD]: http://fritz.box:49000/x_mediaSCPD.xml
func SetDVBCEnable(session *soap.SoapSession, dvbcEnabled bool) (tr064model.SetDVBCEnableResponse, error) {
	fbAction, err := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_media").
		Uri("urn:dslforum-org:service:X_AVM-DE_Media:1").
		Action("SetDVBCEnable").
		AddBoolParam("NewDVBCEnabled", dvbcEnabled).
		Do()
	if err != nil {
		return tr064model.SetDVBCEnableResponse{}, err
	}
	bodyData := fbAction.Body.Data
	result := tr064model.SetDVBCEnableResponse{}
	err = xml.Unmarshal(bodyData, &result)
	if err != nil {
		return tr064model.SetDVBCEnableResponse{}, err
	}
	return result, nil
}
