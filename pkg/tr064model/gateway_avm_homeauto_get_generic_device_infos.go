package tr064model

// GetGenericDeviceInfosResponse auto generated model from [x_homeautoSCPD],
// based on SOAP action 'GetGenericDeviceInfos', Fritz!Box-System-Version 164.07.57
//
// [x_homeautoSCPD]: http://fritz.box:49000/x_homeautoSCPD.xml
type GetGenericDeviceInfosResponse struct {
	AIN                    string `xml:"NewAIN"`                    //
	DeviceId               int    `xml:"NewDeviceId"`               // default=0
	FunctionBitMask        int    `xml:"NewFunctionBitMask"`        // default=0
	FirmwareVersion        string `xml:"NewFirmwareVersion"`        //
	Manufacturer           string `xml:"NewManufacturer"`           //
	ProductName            string `xml:"NewProductName"`            //
	DeviceName             string `xml:"NewDeviceName"`             //
	Present                string `xml:"NewPresent"`                // oneOf=[DISCONNECTED,REGISTRERED,CONNECTED,UNKNOWN]
	MultimeterIsEnabled    string `xml:"NewMultimeterIsEnabled"`    // oneOf=[DISABLED,ENABLED,UNDEFINED]
	MultimeterIsValid      string `xml:"NewMultimeterIsValid"`      // oneOf=[INVALID,VALID,UNDEFINED]
	MultimeterPower        int    `xml:"NewMultimeterPower"`        // default=0
	MultimeterEnergy       int    `xml:"NewMultimeterEnergy"`       // default=0
	TemperatureIsEnabled   string `xml:"NewTemperatureIsEnabled"`   // oneOf=[DISABLED,ENABLED,UNDEFINED]
	TemperatureIsValid     string `xml:"NewTemperatureIsValid"`     // oneOf=[INVALID,VALID,UNDEFINED]
	TemperatureCelsius     int    `xml:"NewTemperatureCelsius"`     // default=0
	TemperatureOffset      int    `xml:"NewTemperatureOffset"`      // default=0
	SwitchIsEnabled        string `xml:"NewSwitchIsEnabled"`        // oneOf=[DISABLED,ENABLED,UNDEFINED]
	SwitchIsValid          string `xml:"NewSwitchIsValid"`          // oneOf=[INVALID,VALID,UNDEFINED]
	SwitchState            string `xml:"NewSwitchState"`            // oneOf=[OFF,ON,TOGGLE,UNDEFINED]
	SwitchMode             string `xml:"NewSwitchMode"`             // oneOf=[AUTO,MANUAL,UNDEFINED]
	SwitchLock             bool   `xml:"NewSwitchLock"`             // default=0
	HkrIsEnabled           string `xml:"NewHkrIsEnabled"`           // oneOf=[DISABLED,ENABLED,UNDEFINED]
	HkrIsValid             string `xml:"NewHkrIsValid"`             // oneOf=[INVALID,VALID,UNDEFINED]
	HkrIsTemperature       int    `xml:"NewHkrIsTemperature"`       // default=0
	HkrSetVentilStatus     string `xml:"NewHkrSetVentilStatus"`     // oneOf=[CLOSED,OPEN,TEMP]
	HkrSetTemperature      int    `xml:"NewHkrSetTemperature"`      // default=0
	HkrReduceVentilStatus  string `xml:"NewHkrReduceVentilStatus"`  // oneOf=[CLOSED,OPEN,TEMP]
	HkrReduceTemperature   int    `xml:"NewHkrReduceTemperature"`   // default=0
	HkrComfortVentilStatus string `xml:"NewHkrComfortVentilStatus"` // oneOf=[CLOSED,OPEN,TEMP]
	HkrComfortTemperature  int    `xml:"NewHkrComfortTemperature"`  // default=0
}
