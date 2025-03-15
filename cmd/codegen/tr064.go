package main

import (
	"errors"
	"fmt"
	"github.com/nitram509/gofritz/pkg/scpd"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var sdcpFolderName = time.Now().UTC().Format("20060102T150405Z")

type tr64Desc struct {
	root     scpd.ServiceControlledProtocolDescriptions
	services []struct {
		deviceType string
		serviceId  string
		spec       scpd.ServiceControlledProtocolDescriptions
	}
}

func fetchAllTr64SDescription(f fetcher) tr64Desc {
	result := tr64Desc{}
	root, err := f.fetchAndParseResponse("/tr64desc.xml")
	if err != nil {
		panic(err)
	}
	result.root = root

	for _, service := range root.Device.ServiceList {
		serviceSpec := retrieveServiceSepc(f, service)
		serviceSpec.deviceType = root.Device.DeviceType
		result.services = append(result.services, serviceSpec)
	}
	for _, device := range root.Device.DeviceList {
		for _, service := range device.ServiceList {
			serviceSpec := retrieveServiceSepc(f, service)
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

func retrieveServiceSepc(f fetcher, service scpd.Service) struct {
	deviceType string
	serviceId  string
	spec       scpd.ServiceControlledProtocolDescriptions
} {
	spec, err := f.fetchAndParseResponse(service.SCPDURL)
	if err != nil {
		panic(err)
	}
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

func safe2Disc(fName string, bytes []byte) {
	const basePath = "__sdcp__"
	targetFolder := filepath.Join(basePath, sdcpFolderName)
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

func updateSdcpFolderName(descriptions scpd.ServiceControlledProtocolDescriptions) {
	s := fmt.Sprintf("%s_%s", descriptions.Device.ModelDescription, descriptions.SystemVersion.Display)
	s = strings.Replace(s, "!", "", -1)
	s = strings.Replace(s, " ", "_", -1)
	if len(s) > 5 {
		// only available, when reading the tr64desc.xml
		sdcpFolderName = s
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
