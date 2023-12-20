package tr064model

// XavmGetChangeCounterResponse AUTO-GENERATED (do not edit) model from [hostsSCPD],
// based on SOAP action 'X_AVM-DE_GetChangeCounter', Fritz!Box-System-Version 164.07.57
//
// [hostsSCPD]: http://fritz.box:49000/hostsSCPD.xml
type XavmGetChangeCounterResponse struct {
	ChangeCounter int `xml:"NewX_AVM-DE_ChangeCounter"` // default=0
}
