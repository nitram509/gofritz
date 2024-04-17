package main

import (
	"encoding/xml"
	"github.com/nitram509/gofritz/pkg/http"
	"github.com/nitram509/gofritz/pkg/scpd"
	"os"
	"path/filepath"
	"strings"
)

type fetcher interface {
	fetchAndParseResponse(name string) (scpd.ServiceControlledProtocolDescriptions, error)
}

type httpFetcher struct {
	baseUrl string
}

type localFileFetcher struct {
	rootPath string
}

func (h httpFetcher) fetchAndParseResponse(urlPath string) (scpd.ServiceControlledProtocolDescriptions, error) {
	bytes, err := http.DoHttpRequest(h.baseUrl + urlPath)
	if err != nil {
		panic(err)
	}
	if safeXmlData2Disc {
		fName, _ := strings.CutPrefix(urlPath, "/")
		safe2Disc(fName, bytes)
	}
	result := scpd.ServiceControlledProtocolDescriptions{}
	err = xml.Unmarshal(bytes, &result)
	return result, err
}

func (l localFileFetcher) fetchAndParseResponse(name string) (scpd.ServiceControlledProtocolDescriptions, error) {
	bytes, err := os.ReadFile(filepath.Join(l.rootPath, name))
	if err != nil {
		panic(err)
	}
	result := scpd.ServiceControlledProtocolDescriptions{}
	err = xml.Unmarshal(bytes, &result)
	return result, err
}
