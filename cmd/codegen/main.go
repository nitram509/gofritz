package main

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const HOST = "fritz.box"
const URL = "http://" + HOST + ":49000"

const debug = false

type StructNameCollector interface {
	addStruct(structName string, actionName string)
}

type structNameCollector struct {
	structs []string
	actions []string
}

func contains(strs []string, item string) int {
	for i, s := range strs {
		if s == item {
			return i
		}
	}
	return -1
}

func (snc *structNameCollector) addStruct(structName string, actionName string) {
	if i := contains(snc.structs, structName); i >= 0 {
		snc.actions[i] = snc.actions[i] + "," + actionName
	} else {
		snc.structs = append(snc.structs, structName)
		snc.actions = append(snc.actions, actionName)
	}
}

func generateModels(description tr64Desc) {
	snc := structNameCollector{}
	for _, service := range description.services {
		generateResponseStructs(service.deviceType, service.serviceId, description.root, service.spec, &snc)
		generateSoapServiceStubs(service.deviceType, service.serviceId, description.root, service.spec)
	}
	//generateSoapResponseAttributes(snc)
}

func generateSoapResponseAttributes(snc structNameCollector) {
	const startMarker = "// --- CODE-GENERATION-START ---"
	const endMarker = "// --- CODE-GENERATION-END ---"

	fileSoapResponseGo := filepath.Join("pkg", "tr064model", "soap_response.go")
	data, err := os.ReadFile(fileSoapResponseGo)
	if err != nil {
		panic(err)
	}

	text := string(data)
	idxStart := strings.Index(text, startMarker)
	idxEnd := strings.Index(text, endMarker)
	if idxStart < 0 || idxEnd < 0 {
		panic(errors.New(fmt.Sprintf("markers not found, %s=%d, %s=%d", startMarker, idxStart, endMarker, idxEnd)))
	}

	sb := strings.Builder{}
	for i := 0; i < len(snc.structs); i++ {
		structName := snc.structs[i]
		actionName := snc.actions[i] + "Response"
		decl := fmt.Sprintf("%s %s `xml:\"%s,omitempty\"`\n", structName, structName, actionName)
		sb.WriteString(decl)
	}

	newText := text[:idxStart] + startMarker + "\n" + sb.String() + text[idxEnd:]
	err = os.WriteFile(fileSoapResponseGo, []byte(newText), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	tr64SDescription := fetchAllTr64SDescription()
	println("URL..............: " + tr64SDescription.root.Device.PresentationURL)
	println("Model............: " + tr64SDescription.root.Device.ModelDescription)
	println("System-Version...: " + tr64SDescription.root.SystemVersion.Display)
	println("No# services.....: " + fmt.Sprintf("%d", len(tr64SDescription.services)))
	for _, service := range tr64SDescription.services {
		println(fmt.Sprintf("- (%s) - %s", service.deviceType, service.serviceId))
	}

	generateModels(tr64SDescription)
}
