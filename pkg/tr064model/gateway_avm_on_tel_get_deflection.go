package tr064model

// GetDeflectionResponse auto generated model from [x_contactSCPD],
// based on SOAP action 'GetDeflection', Fritz!Box-System-Version 164.07.57
//
// [x_contactSCPD]: http://fritz.box:49000/x_contactSCPD.xml
type GetDeflectionResponse struct {
	Enable             bool   `xml:"NewEnable"`             // default=0
	Type               string `xml:"NewType"`               // default=unknown; oneOf=[unknown,toAny,toPOTS,toVoIP,toMSN,fromAnonymous,fromAll,fromNumber,fromVIP,fromNotInPhonebook,fromPB,fon1,fon2,fon3,fon4]
	Number             string `xml:"NewNumber"`             //
	DeflectionToNumber string `xml:"NewDeflectionToNumber"` //
	Mode               string `xml:"NewMode"`               // default=eUnknown; oneOf=[eUnknown,eImmediately,eShortDelayed,eLongDelayed,eBusy,eParallelCall,eNoSignal,eVIP,eDelayed,eDelayedOrBusy,eBellBlockade,eDirectCall,eOff]
	Outgoing           string `xml:"NewOutgoing"`           //
	PhonebookID        int    `xml:"NewPhonebookID"`        // default=0
}
