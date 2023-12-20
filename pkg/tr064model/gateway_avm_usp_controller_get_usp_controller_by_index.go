package tr064model

// GetUSPControllerByIndexResponse auto generated model from [x_uspcontrollerSCPD],
// based on SOAP action 'GetUSPControllerByIndex', Fritz!Box-System-Version 164.07.57
//
// [x_uspcontrollerSCPD]: http://fritz.box:49000/x_uspcontrollerSCPD.xml
type GetUSPControllerByIndexResponse struct {
	Enable                bool   `xml:"NewEnable"`                // default=0
	EndpointID            string `xml:"NewEndpointID"`            //
	MTP                   string `xml:"NewMTP"`                   // default=websocket; oneOf=[websocket,mqtt]
	Hostname              string `xml:"NewHostname"`              //
	Path                  string `xml:"NewPath"`                  //
	Port                  int    `xml:"NewPort"`                  // default=0
	UseTLS                bool   `xml:"NewUseTLS"`                // default=0
	AccessRightSmarthome  bool   `xml:"NewAccessRightSmarthome"`  // default=0
	AccessRightMesh       bool   `xml:"NewAccessRightMesh"`       // default=0
	AccessRightInternet   bool   `xml:"NewAccessRightInternet"`   // default=0
	AccessRightSystem     bool   `xml:"NewAccessRightSystem"`     // default=0
	AccessRightController bool   `xml:"NewAccessRightController"` // default=0
	AccessRightWiFi       bool   `xml:"NewAccessRightWiFi"`       // default=0
	AccessRightVoIP       bool   `xml:"NewAccessRightVoIP"`       // default=0
	Username              string `xml:"NewUsername"`              //
}
