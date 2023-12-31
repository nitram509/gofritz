package scpd

import (
	"encoding/xml"
	"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"testing"
)

func Test_parsing_tr64desc_xml_response(t *testing.T) {
	var spec ServiceControlledProtocolDescriptions
	err := xml.Unmarshal(([]byte)(tr64descXmlResponse), &spec)
	then.AssertThat(t, err, is.Nil())

	then.AssertThat(t, spec.SpecVersion.Major, is.EqualTo("1"))
	then.AssertThat(t, spec.SpecVersion.Minor, is.EqualTo("0"))

	then.AssertThat(t, spec.SystemVersion.HW, is.EqualTo("236"))
	then.AssertThat(t, spec.SystemVersion.Major, is.EqualTo("164"))
	then.AssertThat(t, spec.SystemVersion.Minor, is.EqualTo("7"))
	then.AssertThat(t, spec.SystemVersion.Patch, is.EqualTo("57"))
	then.AssertThat(t, spec.SystemVersion.Buildnumber, is.EqualTo("107823"))
	then.AssertThat(t, spec.SystemVersion.Display, is.EqualTo("164.07.57"))

	then.AssertThat(t, spec.Device.DeviceType, is.EqualTo("urn:dslforum-org:device:InternetGatewayDevice:1"))
	then.AssertThat(t, spec.Device.FriendlyName, is.EqualTo("foobar"))
	then.AssertThat(t, spec.Device.Manufacturer, is.EqualTo("AVM"))
	then.AssertThat(t, spec.Device.ManufacturerURL, is.EqualTo("www.avm.de"))
	then.AssertThat(t, spec.Device.ModelDescription, is.EqualTo("FRITZ!Box 7530"))
	then.AssertThat(t, spec.Device.ModelName, is.EqualTo("FRITZ!Box 7530"))
	then.AssertThat(t, spec.Device.ModelNumber, is.EqualTo("7530 - avm"))
	then.AssertThat(t, spec.Device.ModelURL, is.EqualTo("www.avm.de"))
	then.AssertThat(t, spec.Device.UDN, is.EqualTo("uuid:11223344-1122-1122-3344-112233445566"))
	then.AssertThat(t, spec.Device.SerialNumber, is.EqualTo("AA:BB:CC:DD:EE:FF"))
	then.AssertThat(t, spec.Device.PresentationURL, is.EqualTo("http://192.168.178.1"))

	then.AssertThat(t, spec.Device.IconList, has.Length[Icon](1))
	then.AssertThat(t, spec.Device.IconList[0].Mimetype, is.EqualTo("image/gif"))
	then.AssertThat(t, spec.Device.IconList[0].Width, is.EqualTo("118"))
	then.AssertThat(t, spec.Device.IconList[0].Height, is.EqualTo("119"))
	then.AssertThat(t, spec.Device.IconList[0].Depth, is.EqualTo("8"))
	then.AssertThat(t, spec.Device.IconList[0].Url, is.EqualTo("/ligd.gif"))

	then.AssertThat(t, spec.Device.ServiceList, has.Length[Service](24))
	then.AssertThat(t, spec.Device.ServiceList[0].ServiceType, is.EqualTo("urn:dslforum-org:service:DeviceInfo:1"))
	then.AssertThat(t, spec.Device.ServiceList[0].ServiceId, is.EqualTo("urn:DeviceInfo-com:serviceId:DeviceInfo1"))
	then.AssertThat(t, spec.Device.ServiceList[0].ControlURL, is.EqualTo("/upnp/control/deviceinfo"))
	then.AssertThat(t, spec.Device.ServiceList[0].EventSubURL, is.EqualTo("/upnp/control/deviceinfo"))
	then.AssertThat(t, spec.Device.ServiceList[0].SCPDURL, is.EqualTo("/deviceinfoSCPD.xml"))

	then.AssertThat(t, spec.Device.DeviceList, has.Length[Device](2))
	then.AssertThat(t, spec.Device.DeviceList[0].DeviceType, is.EqualTo("urn:dslforum-org:device:LANDevice:1"))
	then.AssertThat(t, spec.Device.DeviceList[0].FriendlyName, is.EqualTo("LANDevice - FRITZ!Box 7530"))
	then.AssertThat(t, spec.Device.DeviceList[0].Manufacturer, is.EqualTo("AVM"))
	then.AssertThat(t, spec.Device.DeviceList[0].ManufacturerURL, is.EqualTo("www.avm.de"))
	then.AssertThat(t, spec.Device.DeviceList[0].ModelDescription, is.EqualTo("LANDevice - FRITZ!Box 7530"))
	then.AssertThat(t, spec.Device.DeviceList[0].ModelName, is.EqualTo("LANDevice - FRITZ!Box 7530"))
	then.AssertThat(t, spec.Device.DeviceList[0].ModelNumber, is.EqualTo("7530 - avm"))
	then.AssertThat(t, spec.Device.DeviceList[0].ModelURL, is.EqualTo("www.avm.de"))
	then.AssertThat(t, spec.Device.DeviceList[0].UDN, is.EqualTo("uuid:11223344-1122-1122-3344-112233445566"))
	then.AssertThat(t, spec.Device.DeviceList[0].UPC, is.EqualTo("AVM TR-064"))

	then.AssertThat(t, spec.Device.DeviceList[0].ServiceList[0].ServiceType, is.EqualTo("urn:dslforum-org:service:WLANConfiguration:1"))
	then.AssertThat(t, spec.Device.DeviceList[0].ServiceList[0].ServiceId, is.EqualTo("urn:WLANConfiguration-com:serviceId:WLANConfiguration1"))
	then.AssertThat(t, spec.Device.DeviceList[0].ServiceList[0].ControlURL, is.EqualTo("/upnp/control/wlanconfig1"))
	then.AssertThat(t, spec.Device.DeviceList[0].ServiceList[0].EventSubURL, is.EqualTo("/upnp/control/wlanconfig1"))
	then.AssertThat(t, spec.Device.DeviceList[0].ServiceList[0].SCPDURL, is.EqualTo("/wlanconfigSCPD.xml"))
}

func Test_parsing_hostsSCPD_xml_response(t *testing.T) {
	var spec ServiceControlledProtocolDescriptions
	err := xml.Unmarshal(([]byte)(hostsScpdXmlResponse), &spec)
	then.AssertThat(t, err, is.Nil())

	then.AssertThat(t, spec.ActionList, has.Length[Action](19))
	then.AssertThat(t, spec.ActionList[1].Name, is.EqualTo("GetSpecificHostEntry"))

	then.AssertThat(t, spec.ActionList[1].ArgumentList, has.Length[Argument](7))
	then.AssertThat(t, spec.ActionList[1].ArgumentList[0].Name, is.EqualTo("NewMACAddress"))
	then.AssertThat(t, spec.ActionList[1].ArgumentList[0].Direction, is.EqualTo("in"))
	then.AssertThat(t, spec.ActionList[1].ArgumentList[0].RelatedStateVariable, is.EqualTo("MACAddress"))

	then.AssertThat(t, spec.ServiceStateTable, has.Length[StateVariable](34))
	then.AssertThat(t, spec.ServiceStateTable[25].Name, is.EqualTo("X_AVM-DE_UpdateSuccessful"))
	then.AssertThat(t, spec.ServiceStateTable[25].DataType, is.EqualTo("string"))
	then.AssertThat(t, spec.ServiceStateTable[25].DefaultValue, is.EqualTo("unknown"))
	then.AssertThat(t, spec.ServiceStateTable[25].SendEvents, is.EqualTo("no"))
	then.AssertThat(t, spec.ServiceStateTable[25].AllowedValueList, has.Length[string](3))
	then.AssertThat(t, spec.ServiceStateTable[25].AllowedValueList[0], is.EqualTo("unknown"))
	then.AssertThat(t, spec.ServiceStateTable[25].AllowedValueList[1], is.EqualTo("failed"))
	then.AssertThat(t, spec.ServiceStateTable[25].AllowedValueList[2], is.EqualTo("succeeded"))
}

var tr64descXmlResponse = `
<?xml version="1.0"?>
<root xmlns="urn:dslforum-org:device-1-0">
    <specVersion>
        <major>1</major>
        <minor>0</minor>
    </specVersion>
    <systemVersion>
        <HW>236</HW>
        <Major>164</Major>
        <Minor>7</Minor>
        <Patch>57</Patch>
        <Buildnumber>107823</Buildnumber>
        <Display>164.07.57</Display>
    </systemVersion>
    <device>
        <deviceType>urn:dslforum-org:device:InternetGatewayDevice:1</deviceType>
        <friendlyName>foobar</friendlyName>
        <manufacturer>AVM</manufacturer>
        <manufacturerURL>www.avm.de</manufacturerURL>
        <modelDescription>FRITZ!Box 7530</modelDescription>
        <modelName>FRITZ!Box 7530</modelName>
        <modelNumber>7530 - avm</modelNumber>
        <modelURL>www.avm.de</modelURL>
        <UDN>uuid:11223344-1122-1122-3344-112233445566</UDN>
        <serialNumber>AA:BB:CC:DD:EE:FF</serialNumber>
        <originUDN></originUDN>
        <iconList>
            <icon>
                <mimetype>image/gif</mimetype>
                <width>118</width>
                <height>119</height>
                <depth>8</depth>
                <url>/ligd.gif</url>
            </icon>
        </iconList>
        <serviceList>
            <service>
                <serviceType>urn:dslforum-org:service:DeviceInfo:1</serviceType>
                <serviceId>urn:DeviceInfo-com:serviceId:DeviceInfo1</serviceId>
                <controlURL>/upnp/control/deviceinfo</controlURL>
                <eventSubURL>/upnp/control/deviceinfo</eventSubURL>
                <SCPDURL>/deviceinfoSCPD.xml</SCPDURL>
            </service>
            <service>
                <serviceType>urn:dslforum-org:service:DeviceConfig:1</serviceType>
                <serviceId>urn:DeviceConfig-com:serviceId:DeviceConfig1</serviceId>
                <controlURL>/upnp/control/deviceconfig</controlURL>
                <eventSubURL>/upnp/control/deviceconfig</eventSubURL>
                <SCPDURL>/deviceconfigSCPD.xml</SCPDURL>
            </service>
            <service>
                <serviceType>urn:dslforum-org:service:Layer3Forwarding:1</serviceType>
                <serviceId>urn:Layer3Forwarding-com:serviceId:Layer3Forwarding1</serviceId>
                <controlURL>/upnp/control/layer3forwarding</controlURL>
                <eventSubURL>/upnp/control/layer3forwarding</eventSubURL>
                <SCPDURL>/layer3forwardingSCPD.xml</SCPDURL>
            </service>
            <service>
                <serviceType>urn:dslforum-org:service:LANConfigSecurity:1</serviceType>
                <serviceId>urn:LANConfigSecurity-com:serviceId:LANConfigSecurity1</serviceId>
                <controlURL>/upnp/control/lanconfigsecurity</controlURL>
                <eventSubURL>/upnp/control/lanconfigsecurity</eventSubURL>
                <SCPDURL>/lanconfigsecuritySCPD.xml</SCPDURL>
            </service>
            <service>
                <serviceType>urn:dslforum-org:service:ManagementServer:1</serviceType>
                <serviceId>urn:ManagementServer-com:serviceId:ManagementServer1</serviceId>
                <controlURL>/upnp/control/mgmsrv</controlURL>
                <eventSubURL>/upnp/control/mgmsrv</eventSubURL>
                <SCPDURL>/mgmsrvSCPD.xml</SCPDURL>
            </service>
            <service>
                <serviceType>urn:dslforum-org:service:Time:1</serviceType>
                <serviceId>urn:Time-com:serviceId:Time1</serviceId>
                <controlURL>/upnp/control/time</controlURL>
                <eventSubURL>/upnp/control/time</eventSubURL>
                <SCPDURL>/timeSCPD.xml</SCPDURL>
            </service>
            <service>
                <serviceType>urn:dslforum-org:service:UserInterface:1</serviceType>
                <serviceId>urn:UserInterface-com:serviceId:UserInterface1</serviceId>
                <controlURL>/upnp/control/userif</controlURL>
                <eventSubURL>/upnp/control/userif</eventSubURL>
                <SCPDURL>/userifSCPD.xml</SCPDURL>
            </service>
            <service>
                <serviceType>urn:dslforum-org:service:X_AVM-DE_Storage:1</serviceType>
                <serviceId>urn:X_AVM-DE_Storage-com:serviceId:X_AVM-DE_Storage1</serviceId>
                <controlURL>/upnp/control/x_storage</controlURL>
                <eventSubURL>/upnp/control/x_storage</eventSubURL>
                <SCPDURL>/x_storageSCPD.xml</SCPDURL>
            </service>
            <service>
                <serviceType>urn:dslforum-org:service:X_AVM-DE_WebDAVClient:1</serviceType>
                <serviceId>urn:X_AVM-DE_WebDAV-com:serviceId:X_AVM-DE_WebDAVClient1</serviceId>
                <controlURL>/upnp/control/x_webdav</controlURL>
                <eventSubURL>/upnp/control/x_webdav</eventSubURL>
                <SCPDURL>/x_webdavSCPD.xml</SCPDURL>
            </service>
            <service>
                <serviceType>urn:dslforum-org:service:X_AVM-DE_UPnP:1</serviceType>
                <serviceId>urn:X_AVM-DE_UPnP-com:serviceId:X_AVM-DE_UPnP1</serviceId>
                <controlURL>/upnp/control/x_upnp</controlURL>
                <eventSubURL>/upnp/control/x_upnp</eventSubURL>
                <SCPDURL>/x_upnpSCPD.xml</SCPDURL>
            </service>
            <service>
                <serviceType>urn:dslforum-org:service:X_AVM-DE_Speedtest:1</serviceType>
                <serviceId>urn:X_AVM-DE_Speedtest-com:serviceId:X_AVM-DE_Speedtest1</serviceId>
                <controlURL>/upnp/control/x_speedtest</controlURL>
                <eventSubURL>/upnp/control/x_speedtest</eventSubURL>
                <SCPDURL>/x_speedtestSCPD.xml</SCPDURL>
            </service>
            <service>
                <serviceType>urn:dslforum-org:service:X_AVM-DE_RemoteAccess:1</serviceType>
                <serviceId>urn:X_AVM-DE_RemoteAccess-com:serviceId:X_AVM-DE_RemoteAccess1</serviceId>
                <controlURL>/upnp/control/x_remote</controlURL>
                <eventSubURL>/upnp/control/x_remote</eventSubURL>
                <SCPDURL>/x_remoteSCPD.xml</SCPDURL>
            </service>
            <service>
                <serviceType>urn:dslforum-org:service:X_AVM-DE_MyFritz:1</serviceType>
                <serviceId>urn:X_AVM-DE_MyFritz-com:serviceId:X_AVM-DE_MyFritz1</serviceId>
                <controlURL>/upnp/control/x_myfritz</controlURL>
                <eventSubURL>/upnp/control/x_myfritz</eventSubURL>
                <SCPDURL>/x_myfritzSCPD.xml</SCPDURL>
            </service>
            <service>
                <serviceType>urn:dslforum-org:service:X_VoIP:1</serviceType>
                <serviceId>urn:X_VoIP-com:serviceId:X_VoIP1</serviceId>
                <controlURL>/upnp/control/x_voip</controlURL>
                <eventSubURL>/upnp/control/x_voip</eventSubURL>
                <SCPDURL>/x_voipSCPD.xml</SCPDURL>
            </service>
            <service>
                <serviceType>urn:dslforum-org:service:X_AVM-DE_OnTel:1</serviceType>
                <serviceId>urn:X_AVM-DE_OnTel-com:serviceId:X_AVM-DE_OnTel1</serviceId>
                <controlURL>/upnp/control/x_contact</controlURL>
                <eventSubURL>/upnp/control/x_contact</eventSubURL>
                <SCPDURL>/x_contactSCPD.xml</SCPDURL>
            </service>
            <service><serviceType>urn:dslforum-org:service:X_AVM-DE_Dect:1</serviceType><serviceId>urn:X_AVM-DE_Dect-com:serviceId:X_AVM-DE_Dect1</serviceId><controlURL>/upnp/control/x_dect</controlURL><eventSubURL>/upnp/control/x_dect</eventSubURL><SCPDURL>/x_dectSCPD.xml</SCPDURL></service>
            <service><serviceType>urn:dslforum-org:service:X_AVM-DE_TAM:1</serviceType><serviceId>urn:X_AVM-DE_TAM-com:serviceId:X_AVM-DE_TAM1</serviceId><controlURL>/upnp/control/x_tam</controlURL><eventSubURL>/upnp/control/x_tam</eventSubURL><SCPDURL>/x_tamSCPD.xml</SCPDURL></service>
            <service>
                <serviceType>urn:dslforum-org:service:X_AVM-DE_AppSetup:1</serviceType>
                <serviceId>urn:X_AVM-DE_AppSetup-com:serviceId:X_AVM-DE_AppSetup1</serviceId>
                <controlURL>/upnp/control/x_appsetup</controlURL>
                <eventSubURL>/upnp/control/x_appsetup</eventSubURL>
                <SCPDURL>/x_appsetupSCPD.xml</SCPDURL>
            </service>
            <service>
                <serviceType>urn:dslforum-org:service:X_AVM-DE_Homeauto:1</serviceType>
                <serviceId>urn:X_AVM-DE_Homeauto-com:serviceId:X_AVM-DE_Homeauto1</serviceId>
                <controlURL>/upnp/control/x_homeauto</controlURL>
                <eventSubURL>/upnp/control/x_homeauto</eventSubURL>
                <SCPDURL>/x_homeautoSCPD.xml</SCPDURL>
            </service>
            <service>
                <serviceType>urn:dslforum-org:service:X_AVM-DE_Homeplug:1</serviceType>
                <serviceId>urn:X_AVM-DE_Homeplug-com:serviceId:X_AVM-DE_Homeplug1</serviceId>
                <controlURL>/upnp/control/x_homeplug</controlURL>
                <eventSubURL>/upnp/control/x_homeplug</eventSubURL>
                <SCPDURL>/x_homeplugSCPD.xml</SCPDURL>
            </service>
            <service>
                <serviceType>urn:dslforum-org:service:X_AVM-DE_Filelinks:1</serviceType>
                <serviceId>urn:X_AVM-DE_Filelinks-com:serviceId:X_AVM-DE_Filelinks1</serviceId>
                <controlURL>/upnp/control/x_filelinks</controlURL>
                <eventSubURL>/upnp/control/x_filelinks</eventSubURL>
                <SCPDURL>/x_filelinksSCPD.xml</SCPDURL>
            </service>
            <service>
                <serviceType>urn:dslforum-org:service:X_AVM-DE_Auth:1</serviceType>
                <serviceId>urn:X_AVM-DE_Auth-com:serviceId:X_AVM-DE_Auth1</serviceId>
                <controlURL>/upnp/control/x_auth</controlURL>
                <eventSubURL>/upnp/control/x_auth</eventSubURL>
                <SCPDURL>/x_authSCPD.xml</SCPDURL>
            </service>
            <service>
                <serviceType>urn:dslforum-org:service:X_AVM-DE_HostFilter:1</serviceType>
                <serviceId>urn:X_AVM-DE_HostFilter-com:serviceId:X_AVM-DE_HostFilter1</serviceId>
                <controlURL>/upnp/control/x_hostfilter</controlURL>
                <eventSubURL>/upnp/control/x_hostfilter</eventSubURL>
                <SCPDURL>/x_hostfilterSCPD.xml</SCPDURL>
            </service>
            <service>
                <serviceType>urn:dslforum-org:service:X_AVM-DE_USPController:1</serviceType>
                <serviceId>urn:X_AVM-DE_USPController-com:serviceId:X_AVM-DE_USPController1</serviceId>
                <controlURL>/upnp/control/x_uspcontroller</controlURL>
                <eventSubURL>/upnp/control/x_uspcontroller</eventSubURL>
                <SCPDURL>/x_uspcontrollerSCPD.xml</SCPDURL>
            </service>
        </serviceList>
        <deviceList>
            <device>
                <deviceType>urn:dslforum-org:device:LANDevice:1</deviceType>
                <friendlyName>LANDevice - FRITZ!Box 7530</friendlyName>
                <manufacturer>AVM</manufacturer>
                <manufacturerURL>www.avm.de</manufacturerURL>
                <modelDescription>LANDevice - FRITZ!Box 7530</modelDescription>
                <modelName>LANDevice - FRITZ!Box 7530</modelName>
                <modelNumber>7530 - avm</modelNumber>
                <modelURL>www.avm.de</modelURL>
                <UDN>uuid:11223344-1122-1122-3344-112233445566</UDN>
                <UPC>AVM TR-064</UPC>
                <serviceList>
                    <service>
                        <serviceType>urn:dslforum-org:service:WLANConfiguration:1</serviceType>
                        <serviceId>urn:WLANConfiguration-com:serviceId:WLANConfiguration1</serviceId>
                        <controlURL>/upnp/control/wlanconfig1</controlURL>
                        <eventSubURL>/upnp/control/wlanconfig1</eventSubURL>
                        <SCPDURL>/wlanconfigSCPD.xml</SCPDURL>
                    </service>
                    <service><serviceType>urn:dslforum-org:service:WLANConfiguration:2</serviceType><serviceId>urn:WLANConfiguration-com:serviceId:WLANConfiguration2</serviceId><controlURL>/upnp/control/wlanconfig2</controlURL><eventSubURL>/upnp/control/wlanconfig2</eventSubURL><SCPDURL>/wlanconfigSCPD.xml</SCPDURL></service>
					<!-- 2 -->
					<service><serviceType>urn:dslforum-org:service:WLANConfiguration:3</serviceType><serviceId>urn:WLANConfiguration-com:serviceId:WLANConfiguration3</serviceId><controlURL>/upnp/control/wlanconfig3</controlURL><eventSubURL>/upnp/control/wlanconfig3</eventSubURL><SCPDURL>/wlanconfigSCPD.xml</SCPDURL></service>
                    <service>
                        <serviceType>urn:dslforum-org:service:Hosts:1</serviceType>
                        <serviceId>urn:LanDeviceHosts-com:serviceId:Hosts1</serviceId>
                        <controlURL>/upnp/control/hosts</controlURL>
                        <eventSubURL>/upnp/control/hosts</eventSubURL>
                        <SCPDURL>/hostsSCPD.xml</SCPDURL>
                    </service>
                    <service>
                        <serviceType>urn:dslforum-org:service:LANEthernetInterfaceConfig:1</serviceType>
                        <serviceId>urn:LANEthernetIfCfg-com:serviceId:LANEthernetInterfaceConfig1</serviceId>
                        <controlURL>/upnp/control/lanethernetifcfg</controlURL>
                        <eventSubURL>/upnp/control/lanethernetifcfg</eventSubURL>
                        <SCPDURL>/ethifconfigSCPD.xml</SCPDURL>
                    </service>
                    <service>
                        <serviceType>urn:dslforum-org:service:LANHostConfigManagement:1</serviceType>
                        <serviceId>urn:LANHCfgMgm-com:serviceId:LANHostConfigManagement1</serviceId>
                        <controlURL>/upnp/control/lanhostconfigmgm</controlURL>
                        <eventSubURL>/upnp/control/lanhostconfigmgm</eventSubURL>
                        <SCPDURL>/lanhostconfigmgmSCPD.xml</SCPDURL>
                    </service>
                </serviceList>
            </device>
            <device>
                <deviceType>urn:dslforum-org:device:WANDevice:1</deviceType>
                <friendlyName>WANDevice - FRITZ!Box 7530</friendlyName>
                <manufacturer>AVM</manufacturer>
                <manufacturerURL>www.avm.de</manufacturerURL>
                <modelDescription>WANDevice - FRITZ!Box 7530</modelDescription>
                <modelName>WANDevice - FRITZ!Box 7530</modelName>
                <modelNumber>7530 - avm</modelNumber>
                <modelURL>www.avm.de</modelURL>
                <UDN>uuid:11223344-1122-1122-3344-112233445566</UDN>
                <UPC>AVM TR-064</UPC>
                <serviceList>
                    <service>
                        <serviceType>urn:dslforum-org:service:WANCommonInterfaceConfig:1</serviceType>
                        <serviceId>urn:WANCIfConfig-com:serviceId:WANCommonInterfaceConfig1</serviceId>
                        <controlURL>/upnp/control/wancommonifconfig1</controlURL>
                        <eventSubURL>/upnp/control/wancommonifconfig1</eventSubURL>
                        <SCPDURL>/wancommonifconfigSCPD.xml</SCPDURL>
                    </service>
                    <service>
                        <serviceType>urn:dslforum-org:service:WANDSLInterfaceConfig:1</serviceType>
                        <serviceId>urn:WANDSLIfConfig-com:serviceId:WANDSLInterfaceConfig1</serviceId>
                        <controlURL>/upnp/control/wandslifconfig1</controlURL>
                        <eventSubURL>/upnp/control/wandslifconfig1</eventSubURL>
                        <SCPDURL>/wandslifconfigSCPD.xml</SCPDURL>
                    </service>
                </serviceList>
                <deviceList>
                    <device>
                        <deviceType>urn:dslforum-org:device:WANConnectionDevice:1</deviceType>
                        <friendlyName>WANConnectionDevice - FRITZ!Box 7530</friendlyName>
                        <manufacturer>AVM</manufacturer>
                        <manufacturerURL>www.avm.de</manufacturerURL>
                        <modelDescription>WANConnectionDevice - FRITZ!Box 7530</modelDescription>
                        <modelName>WANConnectionDevice - FRITZ!Box 7530</modelName>
                        <modelNumber>7530 - avm</modelNumber>
                        <modelURL>www.avm.de</modelURL>
                        <UDN>uuid:11223344-1122-1122-3344-112233445566</UDN>
                        <UPC>AVM TR-064</UPC>
                        <serviceList>
                            <service>
                                <serviceType>urn:dslforum-org:service:X_AVM-DE_WANMobileConnection:1</serviceType>
                                <serviceId>urn:X_AVM-DE_WANMobileConnection-com:serviceId:X_AVM-DE_WANMobileConnection1</serviceId>
                                <controlURL>/upnp/control/x_wanmobileconn</controlURL>
                                <eventSubURL>/upnp/control/x_wanmobileconn</eventSubURL>
                                <SCPDURL>/x_wanmobileconnSCPD.xml</SCPDURL>
                            </service>
                            <service>
                                <serviceType>urn:dslforum-org:service:WANDSLLinkConfig:1</serviceType>
                                <serviceId>urn:WANDSLLinkConfig-com:serviceId:WANDSLLinkConfig1</serviceId>
                                <controlURL>/upnp/control/wandsllinkconfig1</controlURL>
                                <eventSubURL>/upnp/control/wandsllinkconfig1</eventSubURL>
                                <SCPDURL>/wandsllinkconfigSCPD.xml</SCPDURL>
                            </service>
                            <service>
                                <serviceType>urn:dslforum-org:service:WANEthernetLinkConfig:1</serviceType>
                                <serviceId>urn:WANEthernetLinkConfig-com:serviceId:WANEthernetLinkConfig1</serviceId>
                                <controlURL>/upnp/control/wanethlinkconfig1</controlURL>
                                <eventSubURL>/upnp/control/wanethlinkconfig1</eventSubURL>
                                <SCPDURL>/wanethlinkconfigSCPD.xml</SCPDURL>
                            </service>
                            <service>
                                <serviceType>urn:dslforum-org:service:WANPPPConnection:1</serviceType>
                                <serviceId>urn:WANPPPConnection-com:serviceId:WANPPPConnection1</serviceId>
                                <controlURL>/upnp/control/wanpppconn1</controlURL>
                                <eventSubURL>/upnp/control/wanpppconn1</eventSubURL>
                                <SCPDURL>/wanpppconnSCPD.xml</SCPDURL>
                            </service>
                            <service>
                                <serviceType>urn:dslforum-org:service:WANIPConnection:1</serviceType>
                                <serviceId>urn:WANIPConnection-com:serviceId:WANIPConnection1</serviceId>
                                <controlURL>/upnp/control/wanipconnection1</controlURL>
                                <eventSubURL>/upnp/control/wanipconnection1</eventSubURL>
                                <SCPDURL>/wanipconnSCPD.xml</SCPDURL>
                            </service>
                        </serviceList>
                    </device>
                </deviceList>
            </device>
        </deviceList>
        <presentationURL>http://192.168.178.1</presentationURL>
    </device>
</root>
`
var hostsScpdXmlResponse = `
<?xml version="1.0"?>
<scpd xmlns="urn:dslforum-org:service-1-0">
    <specVersion>
        <major>1</major>
        <minor>0</minor>
    </specVersion>
    <actionList>
        <action>
            <name>GetHostNumberOfEntries</name>
            <argumentList>
                <argument>
                    <name>NewHostNumberOfEntries</name>
                    <direction>out</direction>
                    <relatedStateVariable>HostNumberOfEntries</relatedStateVariable>
                </argument>
            </argumentList>
        </action>
        <action>
            <name>GetSpecificHostEntry</name>
            <argumentList>
                <argument>
                    <name>NewMACAddress</name>
                    <direction>in</direction>
                    <relatedStateVariable>MACAddress</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewIPAddress</name>
                    <direction>out</direction>
                    <relatedStateVariable>IPAddress</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewAddressSource</name>
                    <direction>out</direction>
                    <relatedStateVariable>AddressSource</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewLeaseTimeRemaining</name>
                    <direction>out</direction>
                    <relatedStateVariable>LeaseTimeRemaining</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewInterfaceType</name>
                    <direction>out</direction>
                    <relatedStateVariable>InterfaceType</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewActive</name>
                    <direction>out</direction>
                    <relatedStateVariable>Active</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewHostName</name>
                    <direction>out</direction>
                    <relatedStateVariable>HostName</relatedStateVariable>
                </argument>
            </argumentList>
        </action>
        <action>
            <name>GetGenericHostEntry</name>
            <argumentList>
                <argument>
                    <name>NewIndex</name>
                    <direction>in</direction>
                    <relatedStateVariable>HostNumberOfEntries</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewIPAddress</name>
                    <direction>out</direction>
                    <relatedStateVariable>IPAddress</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewAddressSource</name>
                    <direction>out</direction>
                    <relatedStateVariable>AddressSource</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewLeaseTimeRemaining</name>
                    <direction>out</direction>
                    <relatedStateVariable>LeaseTimeRemaining</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewMACAddress</name>
                    <direction>out</direction>
                    <relatedStateVariable>MACAddress</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewInterfaceType</name>
                    <direction>out</direction>
                    <relatedStateVariable>InterfaceType</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewActive</name>
                    <direction>out</direction>
                    <relatedStateVariable>Active</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewHostName</name>
                    <direction>out</direction>
                    <relatedStateVariable>HostName</relatedStateVariable>
                </argument>
            </argumentList>
        </action>
        <action>
            <name>X_AVM-DE_GetInfo</name>
            <argumentList>
                <argument>
                    <name>NewX_AVM-DE_FriendlynameMinChars</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_FriendlynameMinChars</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_FriendlynameMaxChars</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_FriendlynameMaxChars</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_HostnameMinChars</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_HostnameMinChars</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_HostnameMaxChars</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_HostnameMaxChars</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_HostnameAllowedChars</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_HostnameAllowedChars</relatedStateVariable>
                </argument>
            </argumentList>
        </action>
        <action>
            <name>X_AVM-DE_GetChangeCounter</name>
            <argumentList>
                <argument>
                    <name>NewX_AVM-DE_ChangeCounter</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_ChangeCounter</relatedStateVariable>
                </argument>
            </argumentList>
        </action>
        <action>
            <name>X_AVM-DE_SetHostNameByMACAddress</name>
            <argumentList>
                <argument>
                    <name>NewMACAddress</name>
                    <direction>in</direction>
                    <relatedStateVariable>MACAddress</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewHostName</name>
                    <direction>in</direction>
                    <relatedStateVariable>HostName</relatedStateVariable>
                </argument>
            </argumentList>
        </action>
        <action>
            <name>X_AVM-DE_GetAutoWakeOnLANByMACAddress</name>
            <argumentList>
                <argument>
                    <name>NewMACAddress</name>
                    <direction>in</direction>
                    <relatedStateVariable>MACAddress</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewAutoWOLEnabled</name>
                    <direction>out</direction>
                    <relatedStateVariable>AutoWOLEnabled</relatedStateVariable>
                </argument>
            </argumentList>
        </action>
        <action>
            <name>X_AVM-DE_SetAutoWakeOnLANByMACAddress</name>
            <argumentList>
                <argument>
                    <name>NewMACAddress</name>
                    <direction>in</direction>
                    <relatedStateVariable>MACAddress</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewAutoWOLEnabled</name>
                    <direction>in</direction>
                    <relatedStateVariable>AutoWOLEnabled</relatedStateVariable>
                </argument>
            </argumentList>
        </action>
        <action>
            <name>X_AVM-DE_WakeOnLANByMACAddress</name>
            <argumentList>
                <argument>
                    <name>NewMACAddress</name>
                    <direction>in</direction>
                    <relatedStateVariable>MACAddress</relatedStateVariable>
                </argument>
            </argumentList>
        </action>
        <action>
            <name>X_AVM-DE_GetSpecificHostEntryByIP</name>
            <argumentList>
                <argument>
                    <name>NewIPAddress</name>
                    <direction>in</direction>
                    <relatedStateVariable>IPAddress</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewMACAddress</name>
                    <direction>out</direction>
                    <relatedStateVariable>MACAddress</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewActive</name>
                    <direction>out</direction>
                    <relatedStateVariable>Active</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewHostName</name>
                    <direction>out</direction>
                    <relatedStateVariable>HostName</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewInterfaceType</name>
                    <direction>out</direction>
                    <relatedStateVariable>InterfaceType</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_Port</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_Port</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_Speed</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_Speed</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_UpdateAvailable</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_UpdateAvailable</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_UpdateSuccessful</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_UpdateSuccessful</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_InfoURL</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_InfoURL</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_MACAddressList</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_MACAddressList</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_Model</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_Model</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_URL</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_URL</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_Guest</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_Guest</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_RequestClient</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_RequestClient</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_VPN</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_VPN</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_WANAccess</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_WANAccess</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_Disallow</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_Disallow</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_IsMeshable</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_IsMeshable</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_Priority</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_Priority</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_FriendlyName</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_FriendlyName</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_FriendlyNameIsWriteable</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_FriendlyNameIsWriteable</relatedStateVariable>
                </argument>
            </argumentList>
        </action>
        <action>
            <name>X_AVM-DE_HostsCheckUpdate</name>
        </action>
        <action>
            <name>X_AVM-DE_HostDoUpdate</name>
            <argumentList>
                <argument>
                    <name>NewMACAddress</name>
                    <direction>in</direction>
                    <relatedStateVariable>MACAddress</relatedStateVariable>
                </argument>
            </argumentList>
        </action>
        <action>
            <name>X_AVM-DE_SetPrioritizationByIP</name>
            <argumentList>
                <argument>
                    <name>NewIPAddress</name>
                    <direction>in</direction>
                    <relatedStateVariable>IPAddress</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_Priority</name>
                    <direction>in</direction>
                    <relatedStateVariable>X_AVM-DE_Priority</relatedStateVariable>
                </argument>
            </argumentList>
        </action>
        <action>
            <name>X_AVM-DE_GetHostListPath</name>
            <argumentList>
                <argument>
                    <name>NewX_AVM-DE_HostListPath</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_HostListPath</relatedStateVariable>
                </argument>
            </argumentList>
        </action>
        <action>
            <name>X_AVM-DE_GetMeshListPath</name>
            <argumentList>
                <argument>
                    <name>NewX_AVM-DE_MeshListPath</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_MeshListPath</relatedStateVariable>
                </argument>
            </argumentList>
        </action>
        <action>
            <name>X_AVM-DE_GetFriendlyName</name>
            <argumentList>
                <argument>
                    <name>NewX_AVM-DE_FriendlyName</name>
                    <direction>out</direction>
                    <relatedStateVariable>X_AVM-DE_FriendlyName</relatedStateVariable>
                </argument>
            </argumentList>
        </action>
        <action>
            <name>X_AVM-DE_SetFriendlyName</name>
            <argumentList>
                <argument>
                    <name>NewX_AVM-DE_FriendlyName</name>
                    <direction>in</direction>
                    <relatedStateVariable>X_AVM-DE_FriendlyName</relatedStateVariable>
                </argument>
            </argumentList>
        </action>
        <action>
            <name>X_AVM-DE_SetFriendlyNameByIP</name>
            <argumentList>
                <argument>
                    <name>NewIPAddress</name>
                    <direction>in</direction>
                    <relatedStateVariable>IPAddress</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_FriendlyName</name>
                    <direction>in</direction>
                    <relatedStateVariable>X_AVM-DE_FriendlyName</relatedStateVariable>
                </argument>
            </argumentList>
        </action>
        <action>
            <name>X_AVM-DE_SetFriendlyNameByMAC</name>
            <argumentList>
                <argument>
                    <name>NewMACAddress</name>
                    <direction>in</direction>
                    <relatedStateVariable>MACAddress</relatedStateVariable>
                </argument>
                <argument>
                    <name>NewX_AVM-DE_FriendlyName</name>
                    <direction>in</direction>
                    <relatedStateVariable>X_AVM-DE_FriendlyName</relatedStateVariable>
                </argument>
            </argumentList>
        </action>
    </actionList>
    <serviceStateTable>
        <stateVariable sendEvents="yes">
            <name>HostNumberOfEntries</name>
            <dataType>ui2</dataType>
            <defaultValue>0</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>MACAddress</name>
            <dataType>string</dataType>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>IPAddress</name>
            <dataType>string</dataType>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>AddressSource</name>
            <dataType>string</dataType>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>LeaseTimeRemaining</name>
            <dataType>i4</dataType>
            <defaultValue>0</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>InterfaceType</name>
            <dataType>string</dataType>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>Active</name>
            <dataType>boolean</dataType>
            <defaultValue>0</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>HostName</name>
            <dataType>string</dataType>
        </stateVariable>
        <stateVariable sendEvents="yes">
            <name>X_AVM-DE_ChangeCounter</name>
            <dataType>ui4</dataType>
            <defaultValue>0</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>AutoWOLEnabled</name>
            <dataType>boolean</dataType>
            <defaultValue>0</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_Disallow</name>
            <dataType>boolean</dataType>
            <defaultValue>0</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_FriendlyName</name>
            <dataType>string</dataType>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_FriendlyNameIsWriteable</name>
            <dataType>boolean</dataType>
            <defaultValue>0</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_FriendlynameMaxChars</name>
            <dataType>ui2</dataType>
            <defaultValue>64</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_FriendlynameMinChars</name>
            <dataType>ui2</dataType>
            <defaultValue>1</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_Guest</name>
            <dataType>boolean</dataType>
            <defaultValue>0</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_HostnameAllowedChars</name>
            <dataType>string</dataType>
            <defaultValue>0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_HostnameMaxChars</name>
            <dataType>ui2</dataType>
            <defaultValue>63</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_HostnameMinChars</name>
            <dataType>ui2</dataType>
            <defaultValue>0</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_MACAddressList</name>
            <dataType>string</dataType>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_Port</name>
            <dataType>ui4</dataType>
            <defaultValue>0</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_RequestClient</name>
            <dataType>boolean</dataType>
            <defaultValue>0</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_Speed</name>
            <dataType>ui4</dataType>
            <defaultValue>0</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_URL</name>
            <dataType>string</dataType>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_UpdateAvailable</name>
            <dataType>boolean</dataType>
            <defaultValue>0</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_UpdateSuccessful</name>
            <dataType>string</dataType>
            <allowedValueList>
                <allowedValue>unknown</allowedValue>
                <allowedValue>failed</allowedValue>
                <allowedValue>succeeded</allowedValue>
            </allowedValueList>
            <defaultValue>unknown</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_VPN</name>
            <dataType>boolean</dataType>
            <defaultValue>0</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_InfoURL</name>
            <dataType>string</dataType>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_IsMeshable</name>
            <dataType>boolean</dataType>
            <defaultValue>0</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_Priority</name>
            <dataType>boolean</dataType>
            <defaultValue>0</defaultValue>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_Model</name>
            <dataType>string</dataType>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_HostListPath</name>
            <dataType>string</dataType>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_MeshListPath</name>
            <dataType>string</dataType>
        </stateVariable>
        <stateVariable sendEvents="no">
            <name>X_AVM-DE_WANAccess</name>
            <dataType>string</dataType>
            <allowedValueList>
                <allowedValue>unknown</allowedValue>
                <allowedValue>error</allowedValue>
                <allowedValue>granted</allowedValue>
                <allowedValue>denied</allowedValue>
            </allowedValueList>
            <defaultValue>unknown</defaultValue>
        </stateVariable>
    </serviceStateTable>
</scpd>
`
