package rest

import (
	"bytes"
	"context"
	"github.com/BevisDev/backend-template/src/main/common/consts"
	"github.com/BevisDev/backend-template/src/main/common/utils"
	"github.com/BevisDev/backend-template/src/main/infrastructure/logger"
	"io"
	"net/http"
	"time"
)

func addHeaders(r *http.Request, headers map[string]string) {
	if utils.IsNilOrEmpty(headers) || headers[consts.ContentType] == "" {
		r.Header.Set(consts.ContentType, "application/json")
		return
	}

	for key, value := range headers {
		r.Header.Add(key, value)
	}
}

func Post(c context.Context, restReq *RestRequest) *RestResponse {
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
	logger.LogExtRequest(&logger.RequestLogger{
		URL:    restReq.URL,
		Method: http.MethodPost,
		Body:   utils.ToJSONStr(restReq.Body),
		Time:   startTime,
	})

	ctx, cancel := utils.CreateCtxTimeout(c, clientTimeout)
	defer cancel()

	// created request
	request, err = http.NewRequestWithContext(ctx, http.MethodPost,
		restReq.URL, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		logger.Error(state, "Error created request {}", err)
		return &RestResponse{HasError: true, Error: err}
	}

	// build header
	addHeaders(request, restReq.Header)

	// execute request
	return execute(state, request, restReq, startTime)
}

func PostForm(c context.Context, restReq *RestRequest) *RestResponse {
	var (
		state   = utils.GetState(c)
		err     error
		request *http.Request
		reqBody = restReq.BodyForm.Encode()
	)
	startTime := time.Now()
	// log
	logger.LogExtRequest(&logger.RequestLogger{
		State:  state,
		URL:    restReq.URL,
		Method: http.MethodPost,
		Body:   reqBody,
		Time:   startTime,
	})
	ctx, cancel := utils.CreateCtxTimeout(c, clientTimeout)
	defer cancel()

	// created request
	request, err = http.NewRequestWithContext(ctx, http.MethodPost, restReq.URL,
		bytes.NewBufferString(reqBody))
	if err != nil {
		logger.Error(state, "Error created request {}", err)
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
	return execute(state, request, restReq, startTime)
}

func execute(state string, request *http.Request, restReq *RestRequest, startTime time.Time) *RestResponse {
	var (
		response      *http.Response
		err           error
		respBodyBytes []byte
	)
	response, err = httpClient.Do(request)
	if err != nil {
		logger.Error(state, "Error while sending request {}", err)
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
		logger.Error(state, "Error while doing request {}", err)
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
	logger.LogExtResponse(&logger.ResponseLogger{
		State:       state,
		Status:      response.StatusCode,
		DurationSec: duration,
		Body:        respStr,
	})
	return &result
}
