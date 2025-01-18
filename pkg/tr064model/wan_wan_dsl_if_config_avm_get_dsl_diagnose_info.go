package tr064model

import "encoding/xml"

// XavmGetDSLDiagnoseInfoResponse AUTO-GENERATED (do not edit) model from [wandslifconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetDSLDiagnoseInfo', Fritz!Box-System-Version 164.08.00
//
// [wandslifconfigSCPD]: http://fritz.box:49000/wandslifconfigSCPD.xml
type XavmGetDSLDiagnoseInfoResponse struct {
	XMLName             xml.Name // rather for debug purpose
	DSLDiagnoseState    string   `xml:"NewX_AVM-DE_DSLDiagnoseState"`    //
	CableNokDistance    int      `xml:"NewX_AVM-DE_CableNokDistance"`    // default=0
	DSLLastDiagnoseTime int      `xml:"NewX_AVM-DE_DSLLastDiagnoseTime"` // default=0
	DSLSignalLossTime   int      `xml:"NewX_AVM-DE_DSLSignalLossTime"`   // default=0
	DSLActive           bool     `xml:"NewX_AVM-DE_DSLActive"`           // default=0
	DSLSync             bool     `xml:"NewX_AVM-DE_DSLSync"`             // default=0
}
