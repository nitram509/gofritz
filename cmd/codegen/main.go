package main

import (
	"encoding/xml"
	"fmt"
	"github.com/nitram509/gofitz/pkg/http"
	"github.com/nitram509/gofitz/pkg/scpd"
)

const HOST = "fritz.box"
const URL = "http://" + HOST + ":49000"

type tr64Desc struct {
	root     scpd.SpecVersion
	services []struct {
		serviceId string
		spec      scpd.SpecVersion
	}
}

func fetchAllTr64SDescription() tr64Desc {
	result := tr64Desc{}
	root := fetchAndParseResponse("/tr64desc.xml")
	result.root = root

	for _, service := range root.Device.ServiceList {
		spec := fetchAndParseResponse(service.SCPDURL)
		serviceSpec := struct {
			serviceId string
			spec      scpd.SpecVersion
		}{
			serviceId: service.ServiceId,
			spec:      spec,
		}
		result.services = append(result.services, serviceSpec)
	}
	return result
}

func fetchAndParseResponse(urlPath string) scpd.SpecVersion {
	result := scpd.SpecVersion{}
	bytes, err := http.DoHttpRequest(URL + urlPath)
	if err != nil {
		panic(err)
	}
	err = xml.Unmarshal(bytes, &result)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	tr64SDescription := fetchAllTr64SDescription()
	println("URL..............: " + tr64SDescription.root.Device.PresentationURL)
	println("Model............: " + tr64SDescription.root.Device.ModelDescription)
	println("System-Version...: " + tr64SDescription.root.SystemVersion.Display)
	println("No# services.....: " + fmt.Sprintf("%d", len(tr64SDescription.services)))
	for _, service := range tr64SDescription.services {
		println("- " + service.serviceId)
	}
}
