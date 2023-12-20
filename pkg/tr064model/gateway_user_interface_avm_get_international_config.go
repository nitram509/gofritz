package tr064model

// XavmGetInternationalConfigResponse auto generated model from [userifSCPD],
// based on SOAP action 'X_AVM-DE_GetInternationalConfig', Fritz!Box-System-Version 164.07.57
//
// [userifSCPD]: http://fritz.box:49000/userifSCPD.xml
type XavmGetInternationalConfigResponse struct {
	Language     string `xml:"NewX_AVM-DE_Language"`     //
	Country      string `xml:"NewX_AVM-DE_Country"`      //
	Annex        string `xml:"NewX_AVM-DE_Annex"`        //
	LanguageList string `xml:"NewX_AVM-DE_LanguageList"` //
	CountryList  string `xml:"NewX_AVM-DE_CountryList"`  //
	AnnexList    string `xml:"NewX_AVM-DE_AnnexList"`    //
}
