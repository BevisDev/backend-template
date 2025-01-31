package rest

import (
	"bytes"
	"context"
	"github.com/BevisDev/backend-template/consts"
	"github.com/BevisDev/backend-template/logger"
	"github.com/BevisDev/backend-template/utils"
	"io"
	"net/http"
	"net/url"
	"time"
)

type RestRequest struct {
	State    string
	URL      string
	Params   map[string]any
	Header   map[string]string
	Body     any
	BodyForm url.Values
	Result   any
}

type RestResponse struct {
	StatusCode int
	Header     http.Header
	Body       string
	HasError   bool
	IsTimeout  bool
	Error      error
}

type RestClient struct {
	client     *http.Client
	timeoutSec int
}

func NewRestClient(timeoutSec int) *RestClient {
	return &RestClient{
		client:     &http.Client{},
		timeoutSec: timeoutSec,
	}
}

func addHeaders(rq *http.Request, headers map[string]string) {
	if utils.IsNilOrEmpty(headers) || headers[consts.ContentType] == "" {
		rq.Header.Set(consts.ContentType, "application/json")
		return
	}

	for key, value := range headers {
		rq.Header.Add(key, value)
	}
}

func (r *RestClient) execute(state string, request *http.Request, restReq *RestRequest, startTime time.Time) *RestResponse {
	var (
		response      *http.Response
		err           error
		respBodyBytes []byte
	)
	response, err = r.client.Do(request)
	if err != nil {
		// error timeout
		if utils.IsTimedOut(err) {
			return &RestResponse{
				HasError:  true,
				IsTimeout: true,
				Error:     err,
			}
		}
		return &RestResponse{
			HasError: true,
			Error:    err,
		}
	}
	defer response.Body.Close()

	// read body
	respBodyBytes, err = io.ReadAll(response.Body)
	if err != nil {
		return &RestResponse{HasError: true, Error: err}
	}

	var (
		result   RestResponse
		duration = time.Since(startTime)
		respStr  = utils.FromJSONBytes(respBodyBytes)
	)
	if response.StatusCode >= 400 {
		result.HasError = true
	}
	result.StatusCode = response.StatusCode
	result.Header = response.Header

	// mapping result
	if !utils.IsNilOrEmpty(restReq.Result) {
		err = utils.ToStruct(respBodyBytes, restReq.Result)
	} else {
		result.Body = respStr
	}

	// logger
	logger.NewLogger(nil).LogExtResponse(&logger.ResponseLogger{
		State:       state,
		Status:      response.StatusCode,
		DurationSec: duration,
		Body:        respStr,
	})
	return &result
}

func (r *RestClient) Post(c context.Context, restReq *RestRequest) *RestResponse {
	var (
		state        = utils.GetState(c)
		reqBodyBytes []byte
		err          error
		request      *http.Request
	)

	// serialize body
	if !utils.IsNilOrEmpty(restReq.Body) {
		reqBodyBytes = utils.ToJSON(restReq.Body)
	}

	startTime := time.Now()
	// log
	logger.NewLogger(nil).LogExtRequest(&logger.RequestLogger{
		URL:    restReq.URL,
		Method: http.MethodPost,
		Body:   utils.ToJSONStr(restReq.Body),
		Time:   startTime,
	})

	ctx, cancel := utils.CreateCtxTimeout(c, r.timeoutSec)
	defer cancel()

	// created request
	request, err = http.NewRequestWithContext(ctx, http.MethodPost,
		restReq.URL, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return &RestResponse{HasError: true, Error: err}
	}

	// build header
	addHeaders(request, restReq.Header)

	// execute request
	return r.execute(state, request, restReq, startTime)
}

func (r *RestClient) PostForm(c context.Context, restReq *RestRequest) *RestResponse {
	var (
		state   = utils.GetState(c)
		err     error
		request *http.Request
		reqBody = restReq.BodyForm.Encode()
	)
	startTime := time.Now()
	// log
	logger.NewLogger(nil).LogExtRequest(&logger.RequestLogger{
		State:  state,
		URL:    restReq.URL,
		Method: http.MethodPost,
		Body:   reqBody,
		Time:   startTime,
	})
	ctx, cancel := utils.CreateCtxTimeout(c, r.timeoutSec)
	defer cancel()

	// created request
	request, err = http.NewRequestWithContext(ctx, http.MethodPost, restReq.URL,
		bytes.NewBufferString(reqBody))
	if err != nil {
		return &RestResponse{HasError: true, Error: err}
	}

	// build header
	if utils.IsNilOrEmpty(restReq.Header) {
		restReq.Header = make(map[string]string)
		restReq.Header[consts.ContentType] = "application/x-www-form-urlencoded"
	} else if restReq.Header[consts.ContentType] == "" {
		restReq.Header[consts.ContentType] = "application/x-www-form-urlencoded"
	}
	addHeaders(request, restReq.Header)

	// execute request
	return r.execute(state, request, restReq, startTime)
}
