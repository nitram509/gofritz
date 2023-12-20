package main

import "testing"

func Test_toCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"DeviceInfo", args{s: "DeviceInfo"}, "device_info"},
		{"DeviceConfig", args{s: "DeviceConfig"}, "device_config"},
		{"Layer3Forwarding", args{s: "Layer3Forwarding"}, "layer3_forwarding"},
		{"LANConfigSecurity", args{s: "LANConfigSecurity"}, "lan_config_security"},
		{"ManagementServer", args{s: "ManagementServer"}, "management_server"},
		{"Time", args{s: "Time"}, "time"},
		{"UserInterface", args{s: "UserInterface"}, "user_interface"},
		{"X_AVM-DE_Storage", args{s: "X_AVM-DE_Storage"}, "avm_storage"},
		{"X_AVM-DE_WebDAV", args{s: "X_AVM-DE_WebDAV"}, "avm_webdav"},
		{"X_AVM-DE_UPnP", args{s: "X_AVM-DE_UPnP"}, "avm_upnp"},
		{"X_AVM-DE_Speedtest", args{s: "X_AVM-DE_Speedtest"}, "avm_speedtest"},
		{"X_AVM-DE_RemoteAccess", args{s: "X_AVM-DE_RemoteAccess"}, "avm_remote_access"},
		{"X_AVM-DE_MyFritz", args{s: "X_AVM-DE_MyFritz"}, "avm_my_fritz"},
		{"X_VoIP", args{s: "X_VoIP"}, "x_voip"},
		{"X_AVM-DE_OnTel", args{s: "X_AVM-DE_OnTel"}, "avm_on_tel"},
		{"X_AVM-DE_Dect", args{s: "X_AVM-DE_Dect"}, "avm_dect"},
		{"X_AVM-DE_TAM", args{s: "X_AVM-DE_TAM"}, "avm_tam"},
		{"X_AVM-DE_AppSetup", args{s: "X_AVM-DE_AppSetup"}, "avm_app_setup"},
		{"X_AVM-DE_Homeauto", args{s: "X_AVM-DE_Homeauto"}, "avm_homeauto"},
		{"X_AVM-DE_Homeplug", args{s: "X_AVM-DE_Homeplug"}, "avm_homeplug"},
		{"X_AVM-DE_Filelinks", args{s: "X_AVM-DE_Filelinks"}, "avm_filelinks"},
		{"X_AVM-DE_Auth", args{s: "X_AVM-DE_Auth"}, "avm_auth"},
		{"X_AVM-DE_HostFilter", args{s: "X_AVM-DE_HostFilter"}, "avm_host_filter"},
		{"X_AVM-DE_USPController", args{s: "X_AVM-DE_USPController"}, "avm_usp_controller"},
		{"WLANConfiguration", args{s: "WLANConfiguration"}, "wlan_configuration"},
		{"WLANConfiguration", args{s: "WLANConfiguration"}, "wlan_configuration"},
		{"WLANConfiguration", args{s: "WLANConfiguration"}, "wlan_configuration"},
		{"LanDeviceHosts", args{s: "LanDeviceHosts"}, "lan_device_hosts"},
		{"LANEthernetIfCfg", args{s: "LANEthernetIfCfg"}, "lan_ethernet_if_cfg"},
		{"LANHCfgMgm", args{s: "LANHCfgMgm"}, "lan_h_cfg_mgm"},
		{"WANCIfConfig", args{s: "WANCIfConfig"}, "wan_c_if_config"},
		{"WANDSLIfConfig", args{s: "WANDSLIfConfig"}, "wan_dsl_if_config"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := serviceId2SnakeCase(tt.args.s); got != tt.want {
				t.Errorf("serviceId2SnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
