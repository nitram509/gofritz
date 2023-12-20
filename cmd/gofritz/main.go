package main

import (
	"fmt"
	"github.com/nitram509/gofitz/pkg/tr064"
	"github.com/nitram509/gofitz/pkg/tr064/lan"
	"log"
)

func main() {
	session := tr064.Login()
	log.Printf("SessionID: %s", session)

	//resp, _ := lan.XAvmGetSpecificHostEntryByIp(*session, "192.168.178.40")
	//print(resp.HostName)

	//resp, _ := lan.XAvmGetHostList(*session)
	//print(fmt.Sprintf("%v", resp))

	resp, _ := lan.GetLanHCfgMgmInfo(*session)
	print(fmt.Sprintf("%v", resp))
}
