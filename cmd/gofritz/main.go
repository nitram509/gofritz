package main

import (
	"fmt"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064/gateway"
	"os"
)

func main() {
	username := os.Getenv("FB_USERNAME")
	password := os.Getenv("FB_PASSWORD")
	session := soap.NewSession("fritz.box", username, password)

	//resp, _ := lan.XAvmGetSpecificHostEntryByIp(*session, "192.168.178.40")
	//print(resp.HostName)

	CreateUrlSIDResponse, _ := gateway.XavmCreateUrlSID(session)
	print(fmt.Sprintf("Session ID = %s", CreateUrlSIDResponse.UrlSID))

	//resp0, _ := gateway.GetAvmAuthInfo(session)
	//print(fmt.Sprintf("%v", resp0))

	//resp1, _ := lan.XAvmGetHostList(session)
	//print(fmt.Sprintf("%v", resp1))

	//resp2, _ := lan.GetLanHCfgMgmInfo(session)
	//print(fmt.Sprintf("%v", resp2))

}
