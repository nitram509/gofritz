package tr064model

// GetNumberOfDeviceEntriesResponse auto generated model from [x_homeplugSCPD],
// based on SOAP action 'GetNumberOfDeviceEntries', Fritz!Box-System-Version 164.07.57
//
// [x_homeplugSCPD]: http://fritz.box:49000/x_homeplugSCPD.xml
type GetNumberOfDeviceEntriesResponse struct {
	NumberOfEntries int `xml:"NewNumberOfEntries"` // default=0
}
