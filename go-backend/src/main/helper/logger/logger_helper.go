package logger

import (
	"fmt"
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/helper/datetime"
	"github.com/natefinch/lumberjack"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	once   sync.Once
	logger *zap.Logger
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

func initLogger() {
	once.Do(func() {
		encoder := getEncoderLog()
		writeSync := writeSync()
		core := zapcore.NewCore(encoder, writeSync, zapcore.InfoLevel)
		newLogger := zap.New(core, zap.AddCaller())
		logger = newLogger
	})
}

func getEncoderLog() zapcore.Encoder {
	var encodeConfig zapcore.EncoderConfig
	appConfig := config.AppConfig
	profile := appConfig.ServerConfig.Profile

	// handle profile prod
	if profile == "prod" {
		encodeConfig = zap.NewProductionEncoderConfig()
		// 1716714967.877995 -> 2024-12-19T20:04:31.255+0700
		encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		// ts -> time
		encodeConfig.TimeKey = "time"
		// msg -> message
		encodeConfig.MessageKey = "message"
		// info -> INFO
		encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		//"caller": logger/logger.go:91
		encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
		return zapcore.NewJSONEncoder(encodeConfig)
	}

	// handle other profile
	encodeConfig = zap.NewDevelopmentEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"
	encodeConfig.LevelKey = "level"
	encodeConfig.CallerKey = "caller"
	encodeConfig.MessageKey = "message"

	if profile == "dev" {
		return zapcore.NewConsoleEncoder(encodeConfig)
	}

	return zapcore.NewJSONEncoder(encodeConfig)
}

func writeSync() zapcore.WriteSyncer {
	appConfig := config.AppConfig

	// handle profile dev
	if appConfig.ServerConfig.Profile == "dev" {
		return zapcore.AddSync(os.Stdout)
	}

	loggerConfig := appConfig.LoggerConfig
	logger := lumberjack.Logger{
		Filename:   getFilename(loggerConfig.LogDir),
		MaxSize:    loggerConfig.MaxSize,
		MaxBackups: loggerConfig.MaxBackups,
		MaxAge:     loggerConfig.MaxAge,
		Compress:   loggerConfig.Compress,
	}

	// job runner to split log every day
	if loggerConfig.IsSplit {
		c := cron.New()
		c.AddFunc(loggerConfig.CronTime, func() {
			logger.Filename = getFilename(loggerConfig.LogDir)
			logger.Close()
		})
		c.Start()
	}

	return zapcore.AddSync(&logger)
}

func getFilename(folder string) string {
	now := time.Now().Format(datetime.YYYY_MM_DD)
	return filepath.Join(folder, now, "app.log")
}

func log(level zapcore.Level, state string, msg string, args ...interface{}) {
	var message string

	// formater message
	if len(args) != 0 {
		message = formatMessage(msg, args...)
	} else {
		message = msg
	}

	// skip caller before
	logging := logger.WithOptions(zap.AddCallerSkip(2))

	switch level {
	case zapcore.InfoLevel:
		logging.Info(message, zap.String("state", state))
	case zapcore.WarnLevel:
		logging.Warn(message, zap.String("state", state))
	case zapcore.ErrorLevel:
		logging.Error(message, zap.String("state", state))
	case zapcore.PanicLevel:
		logging.Panic(message, zap.String("state", state))
	default:
		logging.Info(message, zap.String("state", state))
	}
}

func formatMessage(msg string, args ...interface{}) string {
	var message string
	if !strings.Contains(msg, "%") {
		message = strings.ReplaceAll(msg, "{}", "%+v")
	} else {
		message = msg
	}
	return fmt.Sprintf(message, args...)
}

func Sync(state string) {
	if logger != nil {
		if err := logger.Sync(); err != nil {
			Error(state, "Error syncing logger: {}", err)
		}
	}
}

func RequestInfo(req *RequestLogger) {
	if logger == nil {
		initLogger()
	}
	logger.WithOptions(zap.AddCallerSkip(1)).Info("REQUEST INFO",
		zap.String("state", req.State),
		zap.String("url", req.URL),
		zap.String("method", req.Method),
		zap.String("query", req.Query),
		zap.Any("header", req.Header),
		zap.Any("body", req.Body),
	)
}

func ResponseInfo(resp *ResponseLogger) {
	if logger == nil {
		initLogger()
	}
	logger.WithOptions(zap.AddCallerSkip(1)).Info("RESPONSE INFO",
		zap.String("state", resp.State),
		zap.Int("status", resp.Status),
		zap.Float64("durationSec", resp.DurationSec.Seconds()),
		zap.Any("header", resp.Header),
		zap.Any("body", resp.Body),
	)
}

func Info(state string, msg string, args ...interface{}) {
	if logger == nil {
		initLogger()
	}
	log(zapcore.InfoLevel, state, msg, args...)
}

func Error(state string, msg string, args ...interface{}) {
	if logger == nil {
		initLogger()
	}
	log(zapcore.ErrorLevel, state, msg, args...)
}

func Warn(state string, msg string, args ...interface{}) {
	if logger == nil {
		initLogger()
	}
	log(zapcore.WarnLevel, state, msg, args...)
}

func Panic(state string, msg string, args ...interface{}) {
	if logger == nil {
		initLogger()
	}
	log(zapcore.PanicLevel, state, msg, args...)
}
