package rest

import (
	"bytes"
	"context"
	"errors"
	"github.com/BevisDev/backend-template/src/main/common/utils"
	"github.com/BevisDev/backend-template/src/main/infrastructure/config"
	"github.com/BevisDev/backend-template/src/main/infrastructure/logger"
	"io"
	"net/http"
	"sync"
	"time"
)

var (
	onceRestClient sync.Once
	httpClient     *http.Client
	clientTimeout  time.Duration
)

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

func InitRestClient(state string) *http.Client {
	cf := config.AppConfig
	if utils.IsNilOrEmpty(cf) ||
		utils.IsNilOrEmpty(cf.ServerConfig) {
		logger.Fatal(state, "Error appConfig is not initialized")
		return nil
	}
	onceRestClient.Do(func() {
		httpClient = &http.Client{}
		clientTimeout = time.Duration(config.AppConfig.ServerConfig.ClientTimeout) * time.Second
	})
	return httpClient
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

func POST(ctx context.Context, req *Request) *Response {
	state := utils.GetState(ctx)
	var body []byte
	// serialize body
	if !utils.IsNilOrEmpty(req.Body) {
		body = utils.ToJSON(req.Body)
	}

	ctxClient, cancel := context.WithTimeout(ctx, clientTimeout)
	defer cancel()

	// created request
	request, err := http.NewRequestWithContext(ctxClient, http.MethodPost, req.URL, bytes.NewBuffer(body))
	if err != nil {
		logger.Error(state, "Error created request {}", err)
		return &Response{HasError: true, Error: err}
	}
	// build header
	addHeaders(request, req.Header)

	// send request
	resp, err := httpClient.Do(request)
	if utils.IsNilOrEmpty(resp) {
		logger.Error(state, "Error response is nil")
		return &Response{HasError: true, Error: errors.New("error response is nil")}
	}
	if err != nil {
		logger.Error(state, "Error while sending request {}", err)
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
		logger.Error(state, "Error while doing request {}", err)
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
		response.Body = utils.ToStruct(respBody, req.Result)
	} else {
		response.Body = utils.FromJSONBytes(respBody)
	}

	return &response
}
