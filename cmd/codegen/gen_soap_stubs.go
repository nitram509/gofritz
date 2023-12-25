package main

import (
	_ "embed"
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
	Parameters     []actionInputParam
}

type actionInputParam struct {
	SoapName         string
	VarName          string
	TypeName         string
	AddParamFuncName string
}

func deriveParamVarName(name string) string {
	name, _ = strings.CutPrefix(name, "New")
	name = deriveSnakeCase(name)
	name = string2CamelCase(name)
	name = strings.ToLower(string(name[0])) + name[1:]
	// workaround special keywords
	if name == "type" {
		name = "aType"
	}
	if name == "interface" {
		name = "aInterface"
	}
	return name
}

func generateSoapServiceStubs(deviceType string, serviceId string, rootSpec scpd.ServiceControlledProtocolDescriptions, serviceSpec scpd.ServiceControlledProtocolDescriptions, snc *structNameCollector) {
	for _, action := range serviceSpec.ActionList {
		tmpl, err := template.New("soap_stub.go.gohtml").Parse(templateSoapStubsGo)
		if err != nil {
			panic(err)
		}

		structName := deriveStructNameCamelCase(serviceId, action.Name)
		funcName := determineFuncName(serviceId, action.Name)
		packageName := derivePackageName(deviceType)
		scpdUrl := findService(rootSpec, serviceId).SCPDURL
		var soapParams []actionInputParam
		funcParams := strings.Builder{}
		for _, argument := range filterArguments(action.ArgumentList, "in") {
			typeName := determineTypeName(serviceSpec, argument.RelatedStateVariable)
			varName := deriveParamVarName(argument.Name)
			soapParams = append(soapParams, actionInputParam{
				SoapName:         argument.Name,
				VarName:          varName,
				TypeName:         typeName,
				AddParamFuncName: string2CamelCase("Add_" + typeName + "_param"),
			})
			funcParams.WriteString(", ")
			funcParams.WriteString(varName + " " + typeName)
		}

		sd := soapStubTemplateDate{
			PackageName:    packageName,
			FuncName:       funcName,
			FuncParameters: funcParams.String(),
			ReturnTypeName: structName,
			ReqPath:        findService(rootSpec, serviceId).ControlURL,
			Uri:            findService(rootSpec, serviceId).ServiceType,
			SCDPShortName:  deriveScdpShortName(scpdUrl),
			SCDPUrlPath:    scpdUrl,
			SystemVersion:  rootSpec.SystemVersion.Display,
			SoapActionName: action.Name,
			Parameters:     soapParams,
		}

		snc.add(soapMetaData{
			deviceType:     deviceType,
			serviceId:      serviceId,
			actionName:     action.Name,
			packageName:    packageName,
			funcName:       funcName,
			respStructName: structName,
		})

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
	}
}
