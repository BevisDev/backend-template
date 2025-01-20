package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SyncAll() {
	if appLogger != nil {
		appLogger.Sync()
	}
	if rrLogger != nil {
		rrLogger.Sync()
	}
	if extLogger != nil {
		extLogger.Sync()
	}
}

func Info(state, msg string, args ...interface{}) {
	logApp(zapcore.InfoLevel, state, msg, args...)
}

func Error(state, msg string, args ...interface{}) {
	logApp(zapcore.ErrorLevel, state, msg, args...)
}

func Warn(state, msg string, args ...interface{}) {
	logApp(zapcore.WarnLevel, state, msg, args...)
}

func Fatal(state, msg string, args ...interface{}) {
	logApp(zapcore.FatalLevel, state, msg, args...)
}

func LogRequest(req *RequestLogger) {
	rrLogger.WithOptions(
		zap.AddCallerSkip(1)).Info(
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

func LogResponse(resp *ResponseLogger) {
	rrLogger.WithOptions(
		zap.AddCallerSkip(1)).Info(
		"[===== RESPONSE INFO =====]",
		zap.String("state", resp.State),
		zap.Int("status", resp.Status),
		zap.Float64("durationSec", resp.DurationSec.Seconds()),
		zap.Any("header", resp.Header),
		zap.Any("body", resp.Body),
	)
}
