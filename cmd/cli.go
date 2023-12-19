package main

import (
	"github.com/nitram509/gofitz/pkg/tr064"
)

func fetchTr64SDescription(hostname string) tr064.SpecVersion {
	// http://192.168.178.1:49000/tr64desc.xml
	return tr064.SpecVersion{}
}

func main() {
	session := tr064.Login()
	//log.Printf("SessionID: %s", session)

	resp := session.GetSpecificHostEntryByIp("192.168.178.22")
	print(resp)
}
