package client

import (
	"bytes"
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/helper"
	"github.com/BevisDev/backend-template/src/main/helper/json"
	"io"
	"net/http"
	"time"
)

var (
	client = &http.Client{
		Timeout: time.Duration(config.AppConfig.ServerConfig.ClientTimeout) * time.Second,
	}
)

type RestClient struct {
	URL         string
	Params      map[string]any
	Header      map[string]string
	Body        any
	ContentType string
	Result      any
}

type Response struct {
	StatusCode int
	Header     http.Header
	Body       any
}

func POST(restClient RestClient) (*Response, error) {
	var body []byte
	var err error

	// serialize body
	if restClient.Body != nil {
		body = json.ToJSON(restClient.Body)
	}

	// created request
	request, err := http.NewRequest("POST", restClient.URL, bytes.NewBuffer(body))
	if err != nil {
		// log error
	}

	// build header
	addHeaders(request, restClient.Header)

	// send request
	resp, err := client.Do(request)
	if err != nil {
		// logger
	}
	defer resp.Body.Close()

	// read body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		// logger
	}

	var response Response
	response.StatusCode = resp.StatusCode
	response.Header = resp.Header

	// mapping result
	if restClient.Result != nil {
		response.Body = json.ToStruct(respBody, restClient.Result)
	} else {
		response.Body = json.FromJSONBytes(respBody)
	}

	return &response, nil
}

func addHeaders(r *http.Request, headers map[string]string) {
	r.Header.Set("Content-Type", "application/json")
	if !helper.IsNilOrEmpty(headers) {
		for key, value := range headers {
			r.Header.Add(key, value)
		}
	}
}
