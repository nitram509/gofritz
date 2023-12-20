package main

import (
	"fmt"
	"github.com/nitram509/gofitz/pkg/tr064"
	"github.com/nitram509/gofitz/pkg/tr064/lan"
	"github.com/nitram509/gofitz/pkg/tr064model"
	"log"
)

func main() {
	session := tr064.Login()
	log.Printf("SessionID: %s", session)

	//resp := tr064.XAvmGetSpecificHostEntryByIp(*session, "192.168.178.40")
	//print(resp.HostName)

	var resp []tr064model.XAvmGetHostListResponse
	resp = lan.XAvmGetHostList(*session)
	print(fmt.Sprintf("%v", resp))
}
