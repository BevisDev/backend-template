package rest

import (
	"bytes"
	"context"
	"errors"
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/helper/json"
	"github.com/BevisDev/backend-template/src/main/helper/logger"
	"github.com/BevisDev/backend-template/src/main/helper/utils"
	"io"
	"net/http"
	"sync"
	"time"
)

var once sync.Once
var instance IRestClient

type ClientExec struct {
	httpClient *http.Client
	timeout    time.Duration
}

type Request struct {
	State  string
	URL    string
	Params map[string]any
	Header map[string]string
	Body   any
	Result any
}

type Response struct {
	StatusCode int
	Header     http.Header
	Body       any
	HasError   bool
	IsTimeout  bool
	Error      error
}

func NewRestClient() IRestClient {
	once.Do(func() {
		instance = &ClientExec{
			httpClient: &http.Client{},
			timeout:    time.Duration(config.AppConfig.ServerConfig.ClientTimeout) * time.Second,
		}
	})
	return instance
}

func (c *ClientExec) POST(req *Request) *Response {
	if utils.IsNilOrEmpty(req.State) {
		req.State = utils.GenUUID()
	}
	var body []byte
	// serialize body
	if !utils.IsNilOrEmpty(req.Body) {
		body = json.ToJSON(req.Body)
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	// created request
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, req.URL, bytes.NewBuffer(body))
	if err != nil {
		logger.Error(req.State, "Error created request {}", err)
		return &Response{HasError: true, Error: err}
	}
	// build header
	addHeaders(request, req.Header)

	// send request
	resp, err := c.httpClient.Do(request)
	if utils.IsNilOrEmpty(resp) {
		logger.Error(req.State, "Error response is nil")
		return &Response{HasError: true, Error: errors.New("error response is nil")}
	}
	if err != nil {
		logger.Error(req.State, "Error while sending request {}", err)
		// error timeout
		if utils.IsTimedOut(err) {
			return &Response{
				HasError:  true,
				IsTimeout: true,
				Error:     err,
			}
		}
		return &Response{
			HasError: true,
			Error:    err,
		}
	}
	defer resp.Body.Close()

	// read body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error(req.State, "Error while doing request {}", err)
		return &Response{HasError: true, Error: err}
	}

	var response Response
	if resp.StatusCode >= 400 {
		response.HasError = true
	}
	response.StatusCode = resp.StatusCode
	response.Header = resp.Header

	// mapping result
	if !utils.IsNilOrEmpty(req.Result) {
		response.Body = json.ToStruct(respBody, req.Result)
	} else {
		response.Body = json.FromJSONBytes(respBody)
	}

	return &response
}

func addHeaders(r *http.Request, headers map[string]string) {
	if utils.IsNilOrEmpty(headers) || headers["Content-Type"] == "" {
		r.Header.Set("Content-Type", "application/json")
		return
	}

	for key, value := range headers {
		r.Header.Add(key, value)
	}
}
