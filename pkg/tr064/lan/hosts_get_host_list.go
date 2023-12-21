package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/http"
	"github.com/nitram509/gofitz/pkg/soap"
	"github.com/nitram509/gofitz/pkg/tr064model"
)

func XAvmGetHostList(session *soap.SoapSession) ([]tr064model.XAvmGetHostListResponse, error) {
	hostListPathResp, err := XAvmGetHostListPath(session)

	var resp struct {
		XMLName xml.Name                             `xml:"List"`
		Items   []tr064model.XAvmGetHostListResponse `xml:"Item,omitempty"`
	}

	bytes, err := http.DoHttpRequest(session.BaseUrl() + hostListPathResp.HostListPath)
	if err != nil {
		panic(err)
	}
	err = xml.Unmarshal(bytes, &resp)
	if err != nil {
		panic(err)
	}
	return resp.Items, err
}
