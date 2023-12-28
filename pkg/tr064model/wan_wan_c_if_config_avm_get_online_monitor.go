package tr064model

import "encoding/xml"

// XavmGetOnlineMonitorResponse AUTO-GENERATED (do not edit) model from [wancommonifconfigSCPD],
// based on SOAP action 'X_AVM-DE_GetOnlineMonitor', Fritz!Box-System-Version 164.07.57
//
// [wancommonifconfigSCPD]: http://fritz.box:49000/wancommonifconfigSCPD.xml
type XavmGetOnlineMonitorResponse struct {
	XMLName               xml.Name // rather for debug purpose
	TotalNumberSyncGroups int      `xml:"NewTotalNumberSyncGroups"` // default=0
	SyncGroupName         string   `xml:"NewSyncGroupName"`         //
	SyncGroupMode         string   `xml:"NewSyncGroupMode"`         //
	Max_ds                int      `xml:"Newmax_ds"`                // default=0
	Max_us                int      `xml:"Newmax_us"`                // default=0
	Ds_current_bps        string   `xml:"Newds_current_bps"`        //
	Mc_current_bps        string   `xml:"Newmc_current_bps"`        //
	Us_current_bps        string   `xml:"Newus_current_bps"`        //
	Prio_realtime_bps     string   `xml:"Newprio_realtime_bps"`     //
	Prio_high_bps         string   `xml:"Newprio_high_bps"`         //
	Prio_default_bps      string   `xml:"Newprio_default_bps"`      //
	Prio_low_bps          string   `xml:"Newprio_low_bps"`          //
}
