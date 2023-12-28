package main

import (
	_ "embed"
	"fmt"
)

const HOST = "fritz.box"
const URL = "http://" + HOST + ":49000"

const responseSuffix = "Response"
const safeXmlData2Disc = false
const writeFiles = true

func generateAllCode(description tr64Desc) {
	snc := &structNameCollector{}
	for _, service := range description.services {
		println(fmt.Sprintf("Generate: (%s) - %s", service.deviceType, service.serviceId))
		generateResponseStructs(service.deviceType, service.serviceId, description.root, service.spec)
		generateSoapServiceStubs(service.deviceType, service.serviceId, description.root, service.spec, snc)
	}
	generateFunctionsTestCode(snc)
	generateFunctionsRegistryCode(snc)
	updateReadme(snc, description.root)
}

func main() {
	tr64SDescription := fetchAllTr64SDescription()
	println("URL..............: " + tr64SDescription.root.Device.PresentationURL)
	println("Model............: " + tr64SDescription.root.Device.ModelDescription)
	println("System-Version...: " + tr64SDescription.root.SystemVersion.Display)
	println("No# services.....: " + fmt.Sprintf("%d", len(tr64SDescription.services)))
	generateAllCode(tr64SDescription)
}
