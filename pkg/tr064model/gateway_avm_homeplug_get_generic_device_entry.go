package tr064model

// GetGenericDeviceEntryResponse AUTO-GENERATED (do not edit) model from [x_homeplugSCPD],
// based on SOAP action 'GetGenericDeviceEntry', Fritz!Box-System-Version 164.07.57
//
// [x_homeplugSCPD]: http://fritz.box:49000/x_homeplugSCPD.xml
type GetGenericDeviceEntryResponse struct {
	MACAddress       string `xml:"NewMACAddress"`       //
	Active           bool   `xml:"NewActive"`           // default=0
	Name             string `xml:"NewName"`             //
	Model            string `xml:"NewModel"`            //
	UpdateAvailable  bool   `xml:"NewUpdateAvailable"`  // default=0
	UpdateSuccessful string `xml:"NewUpdateSuccessful"` // oneOf=[unknown,failed,succeeded]
}
