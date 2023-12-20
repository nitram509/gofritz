package lan

import (
	"encoding/xml"
	"github.com/nitram509/gofitz/pkg/http"
	"github.com/nitram509/gofitz/pkg/models"
	"github.com/nitram509/gofitz/pkg/soap"
)

func XAvmGetHostList(session soap.SoapSession) []models.XAvmGetHostListResponse {
	hostListPathResp := XAvmGetHostListPath(session)

	var resp struct {
		XMLName xml.Name                         `xml:"List"`
		Items   []models.XAvmGetHostListResponse `xml:"Item,omitempty"`
	}
	// FIXME replace with host configuration
	bytes, err := http.DoHttpRequest("http://fritz.box:49000" + hostListPathResp.XAvmHostListPath)
	if err != nil {
		panic(err)
	}
	err = xml.Unmarshal(bytes, &resp)
	if err != nil {
		panic(err)
	}
	return resp.Items
}
