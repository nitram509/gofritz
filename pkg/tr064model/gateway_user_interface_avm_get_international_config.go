package tr064model

import "encoding/xml"

// XavmGetInternationalConfigResponse AUTO-GENERATED (do not edit) model from [userifSCPD],
// based on SOAP action 'X_AVM-DE_GetInternationalConfig', Fritz!Box-System-Version 141.07.57
//
// [userifSCPD]: http://fritz.box:49000/userifSCPD.xml
type XavmGetInternationalConfigResponse struct {
	XMLName      xml.Name // rather for debug purpose
	Language     string   `xml:"NewX_AVM-DE_Language"`     //
	Country      string   `xml:"NewX_AVM-DE_Country"`      //
	Annex        string   `xml:"NewX_AVM-DE_Annex"`        //
	LanguageList string   `xml:"NewX_AVM-DE_LanguageList"` //
	CountryList  string   `xml:"NewX_AVM-DE_CountryList"`  //
	AnnexList    string   `xml:"NewX_AVM-DE_AnnexList"`    //
}
