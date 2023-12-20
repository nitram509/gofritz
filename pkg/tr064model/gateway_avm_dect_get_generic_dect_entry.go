package tr064model

// GetGenericDectEntryResponse auto generated model from [x_dectSCPD],
// based on SOAP action 'GetGenericDectEntry', Fritz!Box-System-Version 164.07.57
//
// [x_dectSCPD]: http://fritz.box:49000/x_dectSCPD.xml
type GetGenericDectEntryResponse struct {
	ID               string `xml:"NewID"`               //
	Active           bool   `xml:"NewActive"`           // default=0
	Name             string `xml:"NewName"`             //
	Model            string `xml:"NewModel"`            //
	UpdateAvailable  bool   `xml:"NewUpdateAvailable"`  // default=0
	UpdateSuccessful string `xml:"NewUpdateSuccessful"` // oneOf=[unknown,failed,succeeded]
	UpdateInfo       string `xml:"NewUpdateInfo"`       //
}
