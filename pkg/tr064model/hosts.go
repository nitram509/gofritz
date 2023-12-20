package tr064model

// XAvmGetSpecificHostEntryByIpResponse auto generated model from [hostsSCPD],
// based on service spec 'X_AVM-DE_GetSpecificHostEntryByIP'
//
// [hostsSCPD]: http://fritz.box/hostsSCPD.xml
type XAvmGetSpecificHostEntryByIpResponse struct {
	MACAddress                  string `xml:"NewMACAddress"`
	Active                      bool   `xml:"NewActive"`
	HostName                    string `xml:"NewHostName"`
	InterfaceType               string `xml:"NewInterfaceType"`
	XAvmPort                    string `xml:"NewX_AVM-DE_Port"`
	XAvmSpeed                   string `xml:"NewX_AVM-DE_Speed"`
	XAvmUpdateAvailable         bool   `xml:"NewX_AVM-DE_UpdateAvailable"`
	XAvmUpdateSuccessful        string `xml:"NewX_AVM-DE_UpdateSuccessful"`
	XAvmInfoURL                 string `xml:"NewX_AVM-DE_InfoURL"`
	XAvmMACAddressList          string `xml:"NewX_AVM-DE_MACAddressList"`
	XAvmModel                   string `xml:"NewX_AVM-DE_Model"`
	XAvmURL                     string `xml:"NewX_AVM-DE_URL"`
	XAvmGuest                   bool   `xml:"NewX_AVM-DE_Guest"`
	XAvmRequestClient           bool   `xml:"NewX_AVM-DE_RequestClient"`
	XAvmVPN                     bool   `xml:"NewX_AVM-DE_VPN"`
	XAvmWANAccess               string `xml:"NewX_AVM-DE_WANAccess"`
	XAvmDisallow                bool   `xml:"NewX_AVM-DE_Disallow"`
	XAvmIsMeshable              bool   `xml:"NewX_AVM-DE_IsMeshable"`
	XAvmPriority                bool   `xml:"NewX_AVM-DE_Priority"`
	XAvmFriendlyName            string `xml:"NewX_AVM-DE_FriendlyName"`
	XAvmFriendlyNameIsWriteable bool   `xml:"NewX_AVM-DE_FriendlyNameIsWriteable"`
}

type XAvmGetHostListPathResponse struct {
	XAvmHostListPath string `xml:"NewX_AVM-DE_HostListPath"`
}
