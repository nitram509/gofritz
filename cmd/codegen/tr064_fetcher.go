package main

import (
	"encoding/xml"
	"errors"
	"github.com/nitram509/gofitz/pkg/http"
	"github.com/nitram509/gofitz/pkg/scpd"
)

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
		if debug {
			break
		}
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
			if debug {
				break
			}
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
	err = xml.Unmarshal(bytes, &result)
	if err != nil {
		panic(err)
	}
	return result
}
