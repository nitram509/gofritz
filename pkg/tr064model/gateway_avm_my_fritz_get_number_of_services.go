package tr064model

// GetNumberOfServicesResponse auto generated model from [x_myfritzSCPD],
// based on SOAP action 'GetNumberOfServices', Fritz!Box-System-Version 164.07.57
//
// [x_myfritzSCPD]: http://fritz.box:49000/x_myfritzSCPD.xml
type GetNumberOfServicesResponse struct {
	NumberOfServices int `xml:"NewNumberOfServices"` // default=0
}
