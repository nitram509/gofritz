package main

import (
	"fmt"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064/lan"
	"os"
)

func main() {
	username := os.Getenv("FB_USERNAME")
	password := os.Getenv("FB_PASSWORD")
	session := soap.NewSession("fritz.box", username, password)

	hostList, err := lan.XAvmGetHostList(session)
	if err != nil {
		panic(err)
	}
	for _, host := range hostList {
		println(fmt.Sprintf("Host: %s, Active:%v", host.XAvmPort, host.Active))
	}
}
