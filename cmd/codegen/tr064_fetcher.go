package main

import (
	"encoding/xml"
	"errors"
	"github.com/nitram509/gofitz/pkg/http"
	"github.com/nitram509/gofitz/pkg/scpd"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const safeXmlData2Disc = false

type tr64Desc struct {
	root     scpd.ServiceControlledProtocolDescriptions
	services []struct {
		deviceType string
		serviceId  string
		spec       scpd.ServiceControlledProtocolDescriptions
	}
}

func fetchAllTr64SDescription() tr64Desc {
	result := tr64Desc{}
	root := fetchAndParseResponse("/tr64desc.xml")
	result.root = root

	for _, service := range root.Device.ServiceList {
		serviceSpec := retrieveServiceSepc(service)
		serviceSpec.deviceType = root.Device.DeviceType
		result.services = append(result.services, serviceSpec)
	}
	for _, device := range root.Device.DeviceList {
		for _, service := range device.ServiceList {
			serviceSpec := retrieveServiceSepc(service)
			serviceSpec.deviceType = device.DeviceType
			if len(serviceSpec.spec.Device.DeviceList) > 0 {
				panic(errors.New("more devices than expected"))
			}
			if len(serviceSpec.spec.Device.ServiceList) > 0 {
				panic(errors.New("more services than expected"))
			}
			result.services = append(result.services, serviceSpec)
		}
	}
	return result
}

func retrieveServiceSepc(service scpd.Service) struct {
	deviceType string
	serviceId  string
	spec       scpd.ServiceControlledProtocolDescriptions
} {
	spec := fetchAndParseResponse(service.SCPDURL)
	serviceSpec := struct {
		deviceType string
		serviceId  string
		spec       scpd.ServiceControlledProtocolDescriptions
	}{
		serviceId: service.ServiceId,
		spec:      spec,
	}
	return serviceSpec
}

func fetchAndParseResponse(urlPath string) scpd.ServiceControlledProtocolDescriptions {
	result := scpd.ServiceControlledProtocolDescriptions{}
	bytes, err := http.DoHttpRequest(URL + urlPath)
	if err != nil {
		panic(err)
	}
	if safeXmlData2Disc {
		fName, _ := strings.CutPrefix(urlPath, "/")
		safe2Disc(fName, bytes)
	}
	err = xml.Unmarshal(bytes, &result)
	if err != nil {
		panic(err)
	}
	return result
}

var timestamp = time.Now().UTC().Format("20060102T150405Z")

func safe2Disc(fName string, bytes []byte) {
	const basePath = "__sdcp__"
	targetFolder := filepath.Join(basePath, timestamp)
	if _, err := os.Stat(targetFolder); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(targetFolder, 0755)
		if err != nil {
			panic(err)
		}
	}
	bytes = anonymizeData(bytes)
	err := os.WriteFile(filepath.Join(targetFolder, fName), bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func anonymizeData(bytes []byte) []byte {
	data := string(bytes)
	// <serialNumber>11:22:33:44:55:66</serialNumber>
	serialNoRegEx, err := regexp.Compile("<serialNumber>([:0-9a-zA-Z]+)</serialNumber>")
	if err != nil {
		panic(err)
	}
	data = serialNoRegEx.ReplaceAllString(data, "<serialNumber>11:22:33:44:55:66</serialNumber>")
	// <UDN>uuid:11223344-aabb-ccdd-eeff-112233445566</UDN>
	udnRegEx, err := regexp.Compile("<UDN>([-:0-9a-zA-Z]+)</UDN>")
	if err != nil {
		panic(err)
	}
	data = udnRegEx.ReplaceAllString(data, "<UDN>uuid:11223344-aabb-ccdd-eeff-112233445566</UDN>")
	return []byte(data)
}
