package main

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const markerStart = "// === START AUTO-GENERATED CODE ==="
const markerEnd = "// === END AUTO-GENERATED CODE ==="

type soapMetaData struct {
	deviceType     string
	serviceId      string
	actionName     string
	packageName    string
	funcName       string
	respStructName string
}

type structNameCollector struct {
	soapMetaDataList []soapMetaData
}

func (snc *structNameCollector) add(smd soapMetaData) {
	snc.soapMetaDataList = append(snc.soapMetaDataList, smd)
}

func generateFunctionsTestCode(snc *structNameCollector) {
	testFileName := filepath.Join("pkg", "tr064", "functions_test.go")

	data, err := os.ReadFile(testFileName)
	if err != nil {
		panic(err)
	}

	funcTestCode := string(data)
	idxStart := strings.Index(funcTestCode, markerStart)
	idxEnd := strings.Index(funcTestCode, markerEnd)
	if idxStart < 0 || idxEnd < 0 {
		panic(errors.New(fmt.Sprintf("file %s missing: markerStart='%s', markerEnd:'%s'", testFileName, markerStart, markerEnd)))
	}

	sb := strings.Builder{}
	for _, md := range snc.soapMetaDataList {
		sb.WriteString(fmt.Sprintf("\"%s--%s--%s\": %s.%s,\n", md.deviceType, md.serviceId, md.actionName, md.packageName, md.funcName))
	}

	funcTestCode = funcTestCode[:idxStart] + markerStart + "\n" + sb.String() + funcTestCode[idxEnd:]

	writeFileAndFormat(testFileName, []byte(funcTestCode))
}
