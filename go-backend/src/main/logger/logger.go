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
	logger *LoggerWrapper
)

type LoggerWrapper struct {
	sugarLogger *zap.SugaredLogger
}

func initLogger() {
	once.Do(func() {
		encoder := getEncoderLog()
		writeSync := writeSync()
		core := zapcore.NewCore(encoder, writeSync, zapcore.InfoLevel)
		sugarLogger := zap.New(core, zap.AddCaller()).Sugar()
		logger = &LoggerWrapper{
			sugarLogger: sugarLogger,
		}
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
		return zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
		)
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

	return zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		zapcore.AddSync(&logger),
	)
}

func getFilename(folder string) string {
	now := time.Now().Format(datetime.YYYY_MM_DD)
	return filepath.Join(folder, now, "app.log")
}

func log(level zapcore.Level, msg string, args ...interface{}) {
	var message string

	// formater message
	if len(args) != 0 {
		message = formatMessage(msg, args...)
	} else {
		message = msg
	}

	// skip caller before
	logging := logger.sugarLogger.WithOptions(zap.AddCallerSkip(2))

	switch level {
	case zapcore.InfoLevel:
		logging.Info(message)
	case zapcore.WarnLevel:
		logging.Warn(message)
	case zapcore.ErrorLevel:
		logging.Error(message)
	case zapcore.FatalLevel:
		logging.Fatal(message)
	case zapcore.PanicLevel:
		logging.Panic(message)
	default:
		logging.Info(message)
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

func Info(msg string, args ...interface{}) {
	if logger == nil {
		initLogger()
	}
	log(zapcore.InfoLevel, msg, args...)
}

func Error(msg string, args ...interface{}) {
	if logger == nil {
		initLogger()
	}
	log(zapcore.ErrorLevel, msg, args...)
}

func Warn(msg string, args ...interface{}) {
	if logger == nil {
		initLogger()
	}
	log(zapcore.WarnLevel, msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	if logger == nil {
		initLogger()
	}
	log(zapcore.FatalLevel, msg, args...)
}

func Panic(msg string, args ...interface{}) {
	if logger == nil {
		initLogger()
	}
	log(zapcore.PanicLevel, msg, args...)
}

func Sync() {
	if logger != nil {
		if err := logger.sugarLogger.Sync(); err != nil {
			Fatal("Error syncing logger: {}", zap.Error(err))
		}
	}
}
