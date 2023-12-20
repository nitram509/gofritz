package main

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/nitram509/gofitz/pkg/scpd"
	"strings"
	"text/template"
)

//go:embed soap_stub.go.gohtml
var templateSoapStubsGo string

type soapStubTemplateDate struct {
	PackageName    string
	FuncName       string
	FuncParameters string
	ReturnTypeName string
	ReqPath        string
	Uri            string
	SoapActionName string
	SCDPShortName  string
	SCDPUrlPath    string
	SystemVersion  string
	Parameters     []struct {
		Name  string
		Value string
	}
}

func findService(rootSpec scpd.ServiceControlledProtocolDescriptions, serviceId string) scpd.Service {
	for _, service := range rootSpec.Device.ServiceList {
		if serviceId == service.ServiceId {
			return service
		}
	}
	for _, device := range rootSpec.Device.DeviceList {
		for _, service := range device.ServiceList {
			if serviceId == service.ServiceId {
				return service
			}
		}
	}
	panic(errors.New(fmt.Sprintf("serviceId '%s' not found", serviceId)))
}

func generateSoapServiceStubs(deviceType string, serviceId string, rootSpec scpd.ServiceControlledProtocolDescriptions, serviceSpec scpd.ServiceControlledProtocolDescriptions) {
	for _, action := range serviceSpec.ActionList {
		tmpl, err := template.New("soap_stub.go.gohtml").Parse(templateSoapStubsGo)
		if err != nil {
			panic(err)
		}

		strucName := deriveStructNameCamelCase(serviceId, action.Name)
		funcName := strucName[:len(strucName)-len(RESPONSESUFFIX)]
		packageName := derivePackageName(deviceType)
		scpdUrl := findService(rootSpec, serviceId).SCPDURL
		sd := soapStubTemplateDate{
			PackageName:    packageName,
			FuncName:       funcName,
			FuncParameters: "",
			ReturnTypeName: strucName,
			ReqPath:        findService(rootSpec, serviceId).ControlURL,
			Uri:            findService(rootSpec, serviceId).ServiceType,
			SCDPShortName:  deriveScdpShortName(scpdUrl),
			SCDPUrlPath:    scpdUrl,
			SystemVersion:  rootSpec.SystemVersion.Display,
			SoapActionName: action.Name,
			Parameters:     nil,
		}

		sb := strings.Builder{}
		err = tmpl.Execute(&sb, sd)
		if err != nil {
			panic(err)
		}
		soapStubFileName := determineSoapStubFileName(deviceType, serviceId, action.Name)
		if writeFiles {
			writeFileAndFormat(soapStubFileName, []byte(sb.String()))
		} else {
			println("===== " + soapStubFileName + "=====")
			println(sb.String())
		}

		break
	}
}
