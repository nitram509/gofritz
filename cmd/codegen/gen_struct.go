package main

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/nitram509/gofitz/pkg/scpd"
	"strings"
	"text/template"
)

type structData struct {
	Name           string
	SCDPShortName  string
	SCDPUrlPath    string
	SystemVersion  string
	SoapActionName string
	Variables      []string
}

//go:embed struct.go.gohtml
var templateStructGo string

func deriveScdpShortName(scpdUrl string) string {
	result := strings.TrimLeft(scpdUrl, "/")
	if strings.HasSuffix(result, ".xml") {
		result = result[:len(result)-4]
	}
	return result
}

func filterArguments(arguments []scpd.Argument, direction string) (result []scpd.Argument) {
	for _, arg := range arguments {
		if strings.EqualFold(arg.Direction, direction) {
			result = append(result, arg)
		}
	}
	return result
}

func deriveStructVariableName(name string) string {
	name, _ = strings.CutPrefix(name, "New")
	name, _ = strings.CutPrefix(name, "X_AVM-DE_")
	return strings.ToUpper(string(name[0])) + name[1:]
}

func determineVariableComment(spec scpd.ServiceControlledProtocolDescriptions, relatedStateVariable string) string {
	result := "//"
	for _, serviceState := range spec.ServiceStateTable {
		if relatedStateVariable == serviceState.Name {
			if len(serviceState.DefaultValue) > 0 {
				result = result + " default=" + strings.TrimSpace(serviceState.DefaultValue)

			}
			if len(serviceState.AllowedValueList) > 0 {
				if strings.Contains(result, "default") {
					result = result + ";"
				}
				for i, allowedValue := range serviceState.AllowedValueList {
					if i == 0 {
						result = result + " oneOf=[" + allowedValue
					} else {
						result = result + "," + allowedValue
					}
				}
				result = result + "]"
			}
			if strings.ToLower(serviceState.DataType) == "datetime" {
				result = result + "; type=Datetime"
			}
			if strings.ToLower(serviceState.DataType) == "uuid" {
				result = result + "; type=UUID"
			}
		}
	}
	return result
}

func derivePackageName(deviceType string) string {
	switch deviceType {
	case "urn:dslforum-org:device:InternetGatewayDevice:1":
		return "gateway"
	case "urn:dslforum-org:device:LANDevice:1":
		return "lan"
	case "urn:dslforum-org:device:WANDevice:1":
		return "wan"
	}
	panic(errors.New("missing package name mapping for deviceType '" + deviceType + "'"))
}

func serviceId2SnakeCase(serviceId string) string {
	serviceGroup := strings.Split(serviceId, ":")[1]
	if !strings.HasSuffix(serviceGroup, "-com") {
		panic(errors.New("unsupported serviceId mapping for '" + serviceId + "'"))
	}
	serviceGroup = serviceGroup[:len(serviceGroup)-4]
	serviceGroup = deriveSnakeCase(serviceGroup)
	return serviceGroup
}

func deriveStructNameCamelCase(serviceId string, actionName string) string {
	if strings.HasPrefix(actionName, "X_AVM-DE_") {
		actionName = "Xavm" + actionName[9:]
	}
	if actionName == "GetInfo" {
		actionName = "Get" + string2CamelCase(serviceId2SnakeCase(serviceId))
		// prevent double 'InfoInfo'
		if !strings.HasSuffix(actionName, "Info") {
			actionName = actionName + "Info"
		}
	} else if actionName == "XavmGetInfo" {
		actionName = "XavmGet" + string2CamelCase(serviceId2SnakeCase(serviceId))
		// prevent double 'InfoInfo'
		if !strings.HasSuffix(actionName, "Info") {
			actionName = actionName + "Info"
		}
	} else if actionName == "SetConfig" {
		actionName = "Set" + string2CamelCase(serviceId2SnakeCase(serviceId)) + "Config"
	} else if actionName == "SetEnable" {
		actionName = "Set" + string2CamelCase(serviceId2SnakeCase(serviceId)) + "Enable"
	} else if actionName == "GetStatistics" {
		actionName = "Get" + string2CamelCase(serviceId2SnakeCase(serviceId)) + "Statistics"
	}
	return actionName + responseSuffix
}

func generateResponseStructs(deviceType string, serviceId string, rootSpec scpd.ServiceControlledProtocolDescriptions, serviceSpec scpd.ServiceControlledProtocolDescriptions) {
	for _, action := range serviceSpec.ActionList {
		tmpl, err := template.New("struct.go.tmpl").Parse(templateStructGo)
		if err != nil {
			panic(err)
		}
		var variables []string
		varNameMaxLen := 0
		outArguments := filterArguments(action.ArgumentList, "out")
		// create variable declaration
		for _, argument := range outArguments {
			varName := deriveStructVariableName(argument.Name)
			varDeclaration := fmt.Sprintf("%s\t%s `xml:\"%s\"`",
				varName,
				determineTypeName(serviceSpec, argument.RelatedStateVariable),
				argument.Name,
			)
			variables = append(variables, varDeclaration)
			if len(varName) > varNameMaxLen {
				varNameMaxLen = len(varName)
			}
		}
		// fix block spacing for alignment
		var varDeclMaxLen = 0
		for i := range variables {
			varDecl := variables[i]
			idx := strings.Index(varDecl, "\t")
			varDecl = strings.Replace(varDecl, "\t", spaces[:varNameMaxLen-idx+1], 1)
			if len(varDecl) > varDeclMaxLen {
				varDeclMaxLen = len(varDecl)
			}
			variables[i] = varDecl
		}
		// append commends with block ident
		for i := range variables {
			varDecl := variables[i]
			comment := determineVariableComment(serviceSpec, outArguments[i].RelatedStateVariable)
			if len(comment) > 0 {
				varDecl = varDecl + spaces[:varDeclMaxLen-len(varDecl)+1] + comment
			}
			variables[i] = varDecl
		}
		// render the template
		strucName := deriveStructNameCamelCase(serviceId, action.Name)
		scpdUrl := findService(rootSpec, serviceId).SCPDURL
		sd := structData{
			Name:           strucName,
			SCDPShortName:  deriveScdpShortName(scpdUrl),
			SCDPUrlPath:    scpdUrl,
			SystemVersion:  rootSpec.SystemVersion.Display,
			SoapActionName: action.Name,
			Variables:      variables,
		}
		sb := strings.Builder{}
		err = tmpl.Execute(&sb, sd)
		if err != nil {
			panic(err)
		}
		structFileName := determineStructFileName(deviceType, serviceId, action.Name)
		if writeFiles {
			writeFileAndFormat(structFileName, []byte(sb.String()))
		} else {
			println("===== " + structFileName + "=====")
			println(sb.String())
		}
	}

}
