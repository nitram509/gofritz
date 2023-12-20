package tr064model

import "encoding/xml"

// GetUserInterfaceInfoResponse AUTO-GENERATED (do not edit) model from [userifSCPD],
// based on SOAP action 'GetInfo', Fritz!Box-System-Version 164.07.57
//
// [userifSCPD]: http://fritz.box:49000/userifSCPD.xml
type GetUserInterfaceInfoResponse struct {
	XMLName                xml.Name // rather for debug purpose
	UpgradeAvailable       bool     `xml:"NewUpgradeAvailable"`              // default=0
	PasswordRequired       bool     `xml:"NewPasswordRequired"`              // default=0
	PasswordUserSelectable bool     `xml:"NewPasswordUserSelectable"`        // default=1
	WarrantyDate           string   `xml:"NewWarrantyDate"`                  // default=0001-01-01T00:00:00; type=Datetime
	Version                string   `xml:"NewX_AVM-DE_Version"`              //
	DownloadURL            string   `xml:"NewX_AVM-DE_DownloadURL"`          //
	InfoURL                string   `xml:"NewX_AVM-DE_InfoURL"`              //
	UpdateState            string   `xml:"NewX_AVM-DE_UpdateState"`          // oneOf=[Started,Stopped,Error,NoUpdate,UpdateAvailable,Unknown]
	BuildType              string   `xml:"NewX_AVM-DE_BuildType"`            // oneOf=[Release,Intern,Work,Personal,Modified,Inhaus,Labor_Beta,Labor_RC,Labor_DSL,Labor_Phone,Labor,Labor_Test,Labor_Plus]
	SetupAssistantStatus   bool     `xml:"NewX_AVM-DE_SetupAssistantStatus"` // default=0
}
