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
	Body        interface{}
	ContentType string
}

func NewRestClient(url string) *RestClient {
	return &RestClient{
		URL:         url,
		Header:      make(map[string]string),
		ContentType: "application/json",
	}
}

func POST(restClient RestClient) {
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

}

func addHeaders(r *http.Request, headers map[string]string) {
	r.Header.Set("Content-Type", "application/json")
	if !helper.IsNilOrEmpty(headers) {
		for key, value := range headers {
			r.Header.Add(key, value)
		}
	}
}
