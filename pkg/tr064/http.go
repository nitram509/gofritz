package tr064

import (
	"errors"
	"github.com/nitram509/gofitz/pkg"
	"io"
	"net/http"
)

func doHttpRequest(reqPath string) []byte {

	url := "http://fritz.box:49000" + reqPath
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "text/xml")
	req.Header.Set("User-Agent", "gofritz/"+pkg.VERSION)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	resp, err := io.ReadAll(response.Body)

	if response.StatusCode != 200 {
		panic(errors.New("Error calling '" + url + "', Status:" + response.Status))
	}

	if err != nil {
		panic(err)
	}

	return resp
}
