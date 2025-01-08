package logger

import (
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/consts"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

const appFilename = "app.log"
const rrFilename = "req_res.log"
const extFilename = "ext.log"

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
