package tr064model

// GetStatisticsTotalResponse auto generated model from [wandslifconfigSCPD],
// based on SOAP action 'GetStatisticsTotal', Fritz!Box-System-Version 164.07.57
//
// [wandslifconfigSCPD]: http://fritz.box:49000/wandslifconfigSCPD.xml
type GetStatisticsTotalResponse struct {
	ReceiveBlocks       int `xml:"NewReceiveBlocks"`       // default=0
	TransmitBlocks      int `xml:"NewTransmitBlocks"`      // default=0
	CellDelin           int `xml:"NewCellDelin"`           // default=0
	LinkRetrain         int `xml:"NewLinkRetrain"`         // default=0
	InitErrors          int `xml:"NewInitErrors"`          // default=0
	InitTimeouts        int `xml:"NewInitTimeouts"`        // default=0
	LossOfFraming       int `xml:"NewLossOfFraming"`       // default=0
	ErroredSecs         int `xml:"NewErroredSecs"`         // default=0
	SeverelyErroredSecs int `xml:"NewSeverelyErroredSecs"` // default=0
	FECErrors           int `xml:"NewFECErrors"`           // default=0
	ATUCFECErrors       int `xml:"NewATUCFECErrors"`       // default=0
	HECErrors           int `xml:"NewHECErrors"`           // default=0
	ATUCHECErrors       int `xml:"NewATUCHECErrors"`       // default=0
	CRCErrors           int `xml:"NewCRCErrors"`           // default=0
	ATUCCRCErrors       int `xml:"NewATUCCRCErrors"`       // default=0
}
