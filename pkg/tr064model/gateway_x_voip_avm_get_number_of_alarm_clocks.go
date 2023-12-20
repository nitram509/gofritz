package tr064model

// XavmGetNumberOfAlarmClocksResponse auto generated model from [x_voipSCPD],
// based on SOAP action 'X_AVM-DE_GetNumberOfAlarmClocks', Fritz!Box-System-Version 164.07.57
//
// [x_voipSCPD]: http://fritz.box:49000/x_voipSCPD.xml
type XavmGetNumberOfAlarmClocksResponse struct {
	NumberOfAlarmClocks int `xml:"NewX_AVM-DE_NumberOfAlarmClocks"` // default=3
}
