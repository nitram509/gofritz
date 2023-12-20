package tr064model

// GetSecurityPortResponse auto generated model from [deviceinfoSCPD],
// based on SOAP action 'GetSecurityPort', Fritz!Box-System-Version 164.07.57
//
// [deviceinfoSCPD]: http://fritz.box:49000/deviceinfoSCPD.xml
type GetSecurityPortResponse struct {
	SecurityPort int `xml:"NewSecurityPort"` // default=0
}
