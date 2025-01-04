package logger

import (
	"go.uber.org/zap"
	"time"
)

type RequestLogger struct {
	State  string
	URL    string
	Query  string
	Method string
	Header any
	Body   any
}

type ResponseLogger struct {
	State       string
	DurationSec time.Duration
	Status      int
	Header      any
	Body        any
}

func RequestInfo(req *RequestLogger) {
	newLogger(rrFilename).WithOptions(
		zap.AddCallerSkip(1)).Info("[===== REQUEST INFO =====]",
		zap.String("state", req.State),
		zap.String("url", req.URL),
		zap.String("method", req.Method),
		zap.String("query", req.Query),
		zap.Any("header", req.Header),
		zap.Any("body", req.Body),
	)
}

func ResponseInfo(resp *ResponseLogger) {
	newLogger(rrFilename).WithOptions(
		zap.AddCallerSkip(1)).Info("[===== RESPONSE INFO =====]",
		zap.String("state", resp.State),
		zap.Int("status", resp.Status),
		zap.Float64("durationSec", resp.DurationSec.Seconds()),
		zap.Any("header", resp.Header),
		zap.Any("body", resp.Body),
	)
}
