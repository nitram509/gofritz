package tr064model

// GetDefaultWEPKeyIndexResponse AUTO-GENERATED (do not edit) model from [wlanconfigSCPD],
// based on SOAP action 'GetDefaultWEPKeyIndex', Fritz!Box-System-Version 164.07.57
//
// [wlanconfigSCPD]: http://fritz.box:49000/wlanconfigSCPD.xml
type GetDefaultWEPKeyIndexResponse struct {
	DefaultWEPKeyIndex int `xml:"NewDefaultWEPKeyIndex"` // default=0
}
