package scpd

type SystemVersion struct {
	HW          string `xml:"HW"`
	Major       string `xml:"Major"`
	Minor       string `xml:"Minor"`
	Patch       string `xml:"Patch"`
	Buildnumber string `xml:"Buildnumber"`
	Display     string `xml:"Display"`
}

type Icon struct {
	Mimetype string `xml:"mimetype"`
	Width    string `xml:"width"`
	Height   string `xml:"height"`
	Depth    string `xml:"depth"`
	Url      string `xml:"url"`
}

type abstractDevice struct {
	DeviceType       string `xml:"deviceType"`
	FriendlyName     string `xml:"friendlyName"`
	Manufacturer     string `xml:"manufacturer"`
	ManufacturerURL  string `xml:"manufacturerURL"`
	ModelDescription string `xml:"modelDescription"`
	ModelName        string `xml:"modelName"`
	ModelNumber      string `xml:"modelNumber"`
	ModelURL         string `xml:"modelURL"`
	UDN              string `xml:"UDN"`
}

type DeviceRoot struct {
	abstractDevice
	SerialNumber    string    `xml:"serialNumber"`
	OriginUDN       string    `xml:"originUDN"`
	IconList        []Icon    `xml:"iconList>icon"`
	ServiceList     []Service `xml:"serviceList>service"`
	DeviceList      []Device  `xml:"deviceList>device"`
	PresentationURL string    `xml:"presentationURL"`
}

type Device struct {
	abstractDevice
	UPC         string    `xml:"UPC"`
	ServiceList []Service `xml:"serviceList>service"`
	DeviceList  []Device  `xml:"deviceList>device"`
}

type Service struct {
	ServiceType string `xml:"serviceType"`
	ServiceId   string `xml:"serviceId"`
	ControlURL  string `xml:"controlURL"`
	EventSubURL string `xml:"eventSubURL"`
	SCPDURL     string `xml:"SCPDURL"`
}

type Argument struct {
	Name                 string `xml:"name"`
	Direction            string `xml:"direction"`
	RelatedStateVariable string `xml:"relatedStateVariable"`
}

type Action struct {
	Name         string     `xml:"name"`
	ArgumentList []Argument `xml:"argumentList>argument,omitempty"`
}

type StateVariable struct {
	Name             string   `xml:"name"`
	SendEvents       string   `xml:"sendEvents,attr"`
	DataType         string   `xml:"dataType"`
	DefaultValue     string   `xml:"defaultValue,omitempty"`
	AllowedValueList []string `xml:"allowedValueList>allowedValue,omitempty"`
}

type ServiceControlledProtocolDescriptions struct {
	SpecVersion struct {
		Major string `xml:"major"`
		Minor string `xml:"minor"`
	} `xml:"specVersion"`
	SystemVersion     SystemVersion   `xml:"systemVersion,omitempty"`
	Device            DeviceRoot      `xml:"device,omitempty"`
	ActionList        []Action        `xml:"actionList>action,omitempty"`
	ServiceStateTable []StateVariable `xml:"serviceStateTable>stateVariable,omitempty"`
}
