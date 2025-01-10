package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
	"time"
)

var (
	rrOnce     sync.Once
	rrInstance IRrLogger
)

type rrLogger struct {
	rrLogger *zap.Logger
}

type Request struct {
	State  string
	URL    string
	Time   time.Time
	Query  string
	Method string
	Header any
	Body   any
}

type Response struct {
	State       string
	DurationSec time.Duration
	Status      int
	Header      any
	Body        any
}

func NewRrLogger() IRrLogger {
	rrOnce.Do(func() {
		encoder := getEncoderLog()
		rrWrite := writeSync(rrFilename)
		rrCore := zapcore.NewCore(encoder, rrWrite, zapcore.InfoLevel)
		logger := zap.New(rrCore, zap.AddCaller())
		rrInstance = &rrLogger{
			rrLogger: logger,
		}
	})
	return rrInstance
}

func (r *rrLogger) Sync(state string) {
	if r.rrLogger != nil {
		if err := r.rrLogger.Sync(); err != nil {
			NewAppLogger().Error(state, "Error syncing rr logger: {}", err)
		}
	}
}

func (r *rrLogger) RequestLogger(req *Request) {
	r.rrLogger.WithOptions(
		zap.AddCallerSkip(2)).Info(
		"[===== REQUEST INFO =====]",
		zap.String("state", req.State),
		zap.String("url", req.URL),
		zap.Time("time", req.Time),
		zap.String("method", req.Method),
		zap.String("query", req.Query),
		zap.Any("header", req.Header),
		zap.Any("body", req.Body),
	)
}

func (r *rrLogger) ResponseLogger(resp *Response) {
	r.rrLogger.WithOptions(
		zap.AddCallerSkip(2)).Info(
		"[===== RESPONSE INFO =====]",
		zap.String("state", resp.State),
		zap.Int("status", resp.Status),
		zap.Float64("durationSec", resp.DurationSec.Seconds()),
		zap.Any("header", resp.Header),
		zap.Any("body", resp.Body),
	)
}
