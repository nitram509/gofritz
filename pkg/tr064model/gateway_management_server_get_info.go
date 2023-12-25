package tr064model

import "encoding/xml"

// GetManagementServerInfoResponse AUTO-GENERATED (do not edit) model from [mgmsrvSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 141.07.57
//
// [mgmsrvSCPD]: http://fritz.box:49000/mgmsrvSCPD.xml
type GetManagementServerInfoResponse struct {
	XMLName                   xml.Name // rather for debug purpose
	URL                       string   `xml:"NewURL"`                       //
	Username                  string   `xml:"NewUsername"`                  //
	PeriodicInformEnable      bool     `xml:"NewPeriodicInformEnable"`      // default=0
	PeriodicInformInterval    int      `xml:"NewPeriodicInformInterval"`    // default=0
	PeriodicInformTime        string   `xml:"NewPeriodicInformTime"`        // default=0001-01-01T00:00:00; type=Datetime
	ParameterKey              string   `xml:"NewParameterKey"`              //
	ParameterHash             string   `xml:"NewParameterHash"`             //
	ConnectionRequestURL      string   `xml:"NewConnectionRequestURL"`      //
	ConnectionRequestUsername string   `xml:"NewConnectionRequestUsername"` //
	UpgradesManaged           bool     `xml:"NewUpgradesManaged"`           // default=0
}
