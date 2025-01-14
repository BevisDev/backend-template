package logger

import (
	"fmt"
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/consts"
	"github.com/BevisDev/backend-template/src/main/helper/utils"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	appFilename = "app.log"
	rrFilename  = "req_res.log"
	extFilename = "ext.log"
)

var (
	appLogger  *zap.Logger
	rrLogger   *zap.Logger
	extLogger  *zap.Logger
	onceLogger sync.Once
)

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

func InitLogger() *zap.Logger {
	if utils.IsNilOrEmpty(config.AppConfig.ServerConfig.Profile) ||
		utils.IsNilOrEmpty(config.AppConfig.LoggerConfig) {
		log.Fatal("Error config Logger is not initialized")
		return appLogger
	}

	onceLogger.Do(func() {
		appLogger = newLogger(appFilename)
		rrLogger = newLogger(rrFilename)
		extLogger = newLogger(extFilename)
	})
	return appLogger
}

func newLogger(name string) *zap.Logger {
	var logger *zap.Logger
	switch name {
	case appFilename:
		encoder := getEncoderLog()
		appWrite := writeSync(appFilename)
		appCore := zapcore.NewCore(encoder, appWrite, zapcore.InfoLevel)
		logger = zap.New(appCore, zap.AddCaller())
		break
	case rrFilename:
		encoder := getEncoderLog()
		rrWrite := writeSync(rrFilename)
		rrCore := zapcore.NewCore(encoder, rrWrite, zapcore.InfoLevel)
		logger = zap.New(rrCore, zap.AddCaller())
		break
	case extFilename:
		encoder := getEncoderLog()
		rrWrite := writeSync(extFilename)
		rrCore := zapcore.NewCore(encoder, rrWrite, zapcore.InfoLevel)
		logger = zap.New(rrCore, zap.AddCaller())
		break
	default:
		log.Fatalf("name %s logger is not supported", name)
	}
	return logger
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

func writeSync(kind string) zapcore.WriteSyncer {
	appConfig := config.AppConfig

	// handle profile dev
	if appConfig.ServerConfig.Profile == "dev" {
		return zapcore.AddSync(os.Stdout)
	}

	loggerConfig := appConfig.LoggerConfig
	lumberLogger := lumberjack.Logger{
		Filename:   getFilename(kind),
		MaxSize:    loggerConfig.MaxSize,
		MaxBackups: loggerConfig.MaxBackups,
		MaxAge:     loggerConfig.MaxAge,
		Compress:   loggerConfig.Compress,
	}

	// job runner to split log every day
	if loggerConfig.IsSplit {
		c := cron.New()
		c.AddFunc(loggerConfig.CronTime, func() {
			lumberLogger.Filename = getFilename(kind)
			lumberLogger.Close()
		})
		c.Start()
	}

	return zapcore.AddSync(&lumberLogger)
}

func getFilename(kind string) string {
	now := time.Now().Format(consts.YYYY_MM_DD)
	loggerConfig := config.AppConfig.LoggerConfig
	switch kind {
	case appFilename:
		return filepath.Join(loggerConfig.LogAppDir, now, appFilename)
	case rrFilename:
		return filepath.Join(loggerConfig.LogRRDir, now, rrFilename)
	case extFilename:
		return filepath.Join(loggerConfig.LogExtDir, now, extFilename)
	default:
		return filepath.Join(loggerConfig.LogAppDir, now, appFilename)
	}
}

func logApp(level zapcore.Level, state string, msg string, args ...interface{}) {
	// new instance
	if appLogger == nil {
		newLogger(appFilename)
	}
	// check state
	if utils.IsNilOrEmpty(state) {
		state = utils.GenUUID()
	}

	// formater message
	var message string
	if len(args) != 0 {
		message = formatMessage(msg, args...)
	} else {
		message = msg
	}

	// skip caller before
	logging := appLogger.WithOptions(zap.AddCallerSkip(2))
	switch level {
	case zapcore.InfoLevel:
		logging.Info(message, zap.String("state", state))
		break
	case zapcore.WarnLevel:
		logging.Warn(message, zap.String("state", state))
		break
	case zapcore.ErrorLevel:
		logging.Error(message, zap.String("state", state))
		break
	case zapcore.FatalLevel:
		logging.Fatal(message, zap.String("state", state))
		break
	default:
		logging.Info(message, zap.String("state", state))
	}
}

func formatMessage(msg string, args ...interface{}) string {
	var message string
	if !strings.Contains(msg, "%") && strings.Contains(msg, "{}") {
		message = strings.ReplaceAll(msg, "{}", "%+v")
	} else {
		message = msg
	}
	return fmt.Sprintf(message, args...)
}

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

func RequestLogger(req *Request) {
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

func ResponseLogger(resp *Response) {
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
