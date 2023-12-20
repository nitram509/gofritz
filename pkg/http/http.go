package http

import (
	"errors"
	"github.com/nitram509/gofitz/pkg"
	"io"
	"net/http"
)

func DoHttpRequest(fullUrl string) ([]byte, error) {
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "text/xml")
	req.Header.Set("User-Agent", "gofritz/"+pkg.VERSION)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	resp, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Error calling '" + fullUrl + "', Status:" + response.Status)
	}

	return resp, nil
}
