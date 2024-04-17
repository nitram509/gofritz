package gateway

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/soap"
	"github.com/nitram509/gofritz/pkg/tr064model"
)

// StationSearch AUTO-GENERATED (do not edit) code from [x_mediaSCPD],
// based on SOAP action 'StationSearch', Fritz!Box-System-Version 141.07.57
//
// [x_mediaSCPD]: http://fritz.box:49000/x_mediaSCPD.xml
func StationSearch(session *soap.SoapSession, stationSearchMode string) (tr064model.StationSearchResponse, error) {
	bodyData := soap.NewSoapRequest(session).
		ReqPath("/upnp/control/x_media").
		Uri("urn:dslforum-org:service:X_AVM-DE_Media:1").
		Action("StationSearch").
		AddStringParam("NewStationSearchMode", stationSearchMode).
		Do().Body.Data
	result := tr064model.StationSearchResponse{}
	err := xml.Unmarshal(bodyData, &result)
	return result, err
}
