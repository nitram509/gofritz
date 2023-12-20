package main

import (
	"fmt"
	"github.com/nitram509/gofitz/pkg/tr064"
	"log"
)

func fetchTr64SDescription(hostname string) tr064.SpecVersion {
	// http://192.168.178.1:49000/tr64desc.xml
	return tr064.SpecVersion{}
}

func main() {
	session := tr064.Login()
	log.Printf("SessionID: %s", session)

	//resp := tr064.XAvmGetSpecificHostEntryByIp(*session, "192.168.178.40")
	//print(resp.HostName)

	resp := tr064.XAvmGetHostList(*session)
	print(fmt.Sprintf("%v", resp))
}
