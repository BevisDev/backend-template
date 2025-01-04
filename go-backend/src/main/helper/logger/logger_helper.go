package logger

import (
	"fmt"
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/consts"
	"github.com/BevisDev/backend-template/src/main/helper/utils"
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
	once      sync.Once
	appLogger *zap.Logger
	rrLogger  *zap.Logger
	extLogger *zap.Logger
)

const appFilename = "app.log"
const rrFilename = "req_res.log"
const extFilename = "ext.log"

func newLogger(kind string) *zap.Logger {
	once.Do(func() {
		encoder := getEncoderLog()
		// new appLogger
		appWrite := writeSync(appFilename)
		appCore := zapcore.NewCore(encoder, appWrite, zapcore.InfoLevel)
		appLogger = zap.New(appCore, zap.AddCaller())

		// new rrLogger
		rrWrite := writeSync(rrFilename)
		rrCore := zapcore.NewCore(encoder, rrWrite, zapcore.InfoLevel)
		rrLogger = zap.New(rrCore, zap.AddCaller())

		// new extLogger
		extWrite := writeSync(extFilename)
		extCore := zapcore.NewCore(encoder, extWrite, zapcore.InfoLevel)
		extLogger = zap.New(extCore, zap.AddCaller())
	})
	switch kind {
	case appFilename:
		return appLogger
	case rrFilename:
		return rrLogger
	case extFilename:
		return extLogger
	default:
		return appLogger
	}
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

func log(level zapcore.Level, state string, msg string, args ...interface{}) {
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
	case zapcore.PanicLevel:
		logging.Panic(message, zap.String("state", state))
		break
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
	if appLogger != nil {
		if err := appLogger.Sync(); err != nil {
			Error(state, "Error syncing logger: {}", err)
		}
	}
}

func Info(state string, msg string, args ...interface{}) {
	log(zapcore.InfoLevel, state, msg, args...)
}

func Error(state string, msg string, args ...interface{}) {
	log(zapcore.ErrorLevel, state, msg, args...)
}

func Warn(state string, msg string, args ...interface{}) {
	log(zapcore.WarnLevel, state, msg, args...)
}

func Panic(state string, msg string, args ...interface{}) {
	log(zapcore.PanicLevel, state, msg, args...)
}
