package tr064model

import (
	"encoding/xml"
)

type SoapResponse struct {
	XMLName       xml.Name `xml:"Envelope"`
	S             string   `xml:"s,attr"`
	EncodingStyle string   `xml:"encodingStyle,attr"`
	Body          struct {
		Data                                 []byte                               `xml:",innerxml"`
		XAvmDeCreateUrlSIDResponse           XAvmDeCreateUrlSIDResponse           `xml:"X_AVM-DE_CreateUrlSIDResponse,omitempty"`
		XAvmGetSpecificHostEntryByIpResponse XAvmGetSpecificHostEntryByIpResponse `xml:"X_AVM-DE_GetSpecificHostEntryByIPResponse,omitempty"`
		XAvmGetHostListPathResponse          XAvmGetHostListPathResponse          `xml:"X_AVM-DE_GetHostListPathResponse,omitempty"`
	} `xml:"Body,omitempty"`
}
