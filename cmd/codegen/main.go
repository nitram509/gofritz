package main

import (
	_ "embed"
	"fmt"
	"os"
)

const responseSuffix = "Response"
const safeXmlData2Disc = true
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

func fetchFromFritzboxAndGenerate(f fetcher) {
	tr64SDescription := fetchAllTr64SDescription(f)
	println("URL..............: " + tr64SDescription.root.Device.PresentationURL)
	println("Model............: " + tr64SDescription.root.Device.ModelDescription)
	println("System-Version...: " + tr64SDescription.root.SystemVersion.Display)
	println("No# services.....: " + fmt.Sprintf("%d", len(tr64SDescription.services)))
	generateAllCode(tr64SDescription)
}

func fetchLocalFilesAndGenerate(rootPath string) {
	f := localFileFetcher{
		rootPath: rootPath,
	}
	fetchFromFritzboxAndGenerate(f)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--host" {
		const HOST = "fritz.box"
		f := httpFetcher{baseUrl: "http://" + HOST + ":49000"}
		fetchFromFritzboxAndGenerate(f)
	} else if len(os.Args) > 1 && os.Args[1] == "--local" {
		//fetchLocalFilesAndGenerate("__sdcp__/FRITZBox_6490_Cable_141_07_57")
		//fetchLocalFilesAndGenerate("__sdcp__/FRITZBox_7590_154_07_57")
		fetchLocalFilesAndGenerate("__sdcp__/FRITZBox_7530_164.07.57")
	} else {
		println("usage: codegen <--host|--local>")
	}
}
