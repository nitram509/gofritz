package tr064model

// XavmDoUpdateResponse AUTO-GENERATED (do not edit) model from [userifSCPD],
// based on SOAP action 'X_AVM-DE_DoUpdate', Fritz!Box-System-Version 164.07.57
//
// [userifSCPD]: http://fritz.box:49000/userifSCPD.xml
type XavmDoUpdateResponse struct {
	UpgradeAvailable bool   `xml:"NewUpgradeAvailable"`     // default=0
	UpdateState      string `xml:"NewX_AVM-DE_UpdateState"` // oneOf=[Started,Stopped,Error,NoUpdate,UpdateAvailable,Unknown]
}
