package main

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/nitram509/gofritz/pkg/scpd"
	"os"
	"path/filepath"
	"strings"
)

const markerMarkdownStart = "<!-- MARKDOWN-AUTO-DOCS:START (CODE:src=functions_registry.go) -->"
const markerMarkdownEnd = "<!-- MARKDOWN-AUTO-DOCS:END -->"

func updateReadme(snc *structNameCollector, rootDesc scpd.ServiceControlledProtocolDescriptions) {
	readmeFileName := filepath.Join("README.md")

	data, err := os.ReadFile(readmeFileName)
	if err != nil {
		panic(err)
	}

	funcTestCode := string(data)
	idxStart := strings.Index(funcTestCode, markerMarkdownStart)
	idxEnd := strings.Index(funcTestCode, markerMarkdownEnd)
	if idxStart < 0 || idxEnd < 0 {
		panic(errors.New(fmt.Sprintf("file %s missing: markerStart='%s', markerEnd:'%s'", readmeFileName, markerMarkdownStart, markerMarkdownEnd)))
	}

	sb := strings.Builder{}
	sb.WriteString("| (Go) function | Discovered via |\n")
	sb.WriteString("|---------------|----------------|\n")
	for _, md := range snc.soapMetaDataList {
		goFileName := determineSoapStubFileName(md.deviceType, md.serviceId, md.actionName)
		sb.WriteString(fmt.Sprintf("| [%s.%s](./%s) | %s v%s |\n",
			md.packageName, md.funcName, goFileName, rootDesc.Device.ModelName, rootDesc.SystemVersion.Display))

	}

	funcTestCode = funcTestCode[:idxStart] + markerMarkdownStart + "\n" + sb.String() + funcTestCode[idxEnd:]

	writeFileAndFormat(readmeFileName, []byte(funcTestCode))
}
