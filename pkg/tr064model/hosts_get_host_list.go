package tr064model

type XAvmGetHostListResponse struct {
	Index                       int    `xml:"Index"`
	IPAddress                   string `xml:"IPAddress"`
	MACAddress                  string `xml:"MACAddress"`
	Active                      bool   `xml:"Active"`
	HostName                    string `xml:"HostName"`
	InterfaceType               string `xml:"InterfaceType"` // One of (“Ethernet”, “802.11”, "HomePlug", “”)
	XAvmPort                    int    `xml:"X_AVM-DE_Port"`
	XAvmSpeed                   string `xml:"X_AVM-DE_Speed"` // Shows the speed in Mbit/s
	XAvmUpdateAvailable         bool   `xml:"X_AVM-DE_UpdateAvailable"`
	XAvmUpdateSuccessful        string `xml:"X_AVM-DE_UpdateSuccessful"`
	XAvmInfoURL                 string `xml:"X_AVM-DE_InfoURL"`
	XAvmMACAddressList          string `xml:"X_AVM-DE_MACAddressList"` // Comma separated list with device MAC Addresses e.g. different interfaces. Only AVM devices have this list, other devices have an empty tag.
	XAvmModel                   string `xml:"X_AVM-DE_Model"`
	XAvmURL                     string `xml:"X_AVM-DE_URL"`
	XAvmGuest                   bool   `xml:"X_AVM-DE_Guest"`
	XAvmRequestClient           bool   `xml:"X_AVM-DE_RequestClient"`
	XAvmVPN                     bool   `xml:"X_AVM-DE_VPN"`
	XAvmWANAccess               string `xml:"X_AVM-DE_WANAccess"` // Shows if the lan device has WAN access “granted”, “denied”, “error” or “unknown”
	XAvmDisallow                bool   `xml:"X_AVM-DE_Disallow"`
	XAvmIsMeshable              bool   `xml:"X_AVM-DE_IsMeshable"`
	XAvmPriority                bool   `xml:"X_AVM-DE_Priority"`
	XAvmFriendlyName            string `xml:"X_AVM-DE_FriendlyName"`
	XAvmFriendlyNameIsWriteable bool   `xml:"X_AVM-DE_FriendlyNameIsWriteable"`
}
