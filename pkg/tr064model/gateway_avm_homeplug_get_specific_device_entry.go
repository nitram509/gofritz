package tr064model

// GetSpecificDeviceEntryResponse auto generated model from [x_homeplugSCPD],
// based on SOAP action 'GetSpecificDeviceEntry', Fritz!Box-System-Version 164.07.57
//
// [x_homeplugSCPD]: http://fritz.box:49000/x_homeplugSCPD.xml
type GetSpecificDeviceEntryResponse struct {
	Active           bool   `xml:"NewActive"`           // default=0
	Name             string `xml:"NewName"`             //
	Model            string `xml:"NewModel"`            //
	UpdateAvailable  bool   `xml:"NewUpdateAvailable"`  // default=0
	UpdateSuccessful string `xml:"NewUpdateSuccessful"` // oneOf=[unknown,failed,succeeded]
}
