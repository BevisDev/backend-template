package logger

import (
	"os"
	"sync"

	"github.com/BevisDev/backend-template/src/main/global"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once   sync.Once
	logger *LoggerWrapper
)

type LoggerWrapper struct {
	logger *zap.Logger
}

func NewLogger() *LoggerWrapper {
	once.Do(func() {
		encoder := getEncoderLog()
		writeSync := writeSync()
		core := zapcore.NewCore(encoder, writeSync, zapcore.InfoLevel)
		logger = &LoggerWrapper{logger: zap.New(core, zap.AddCaller())}
	})
	return logger
}

func getEncoderLog() zapcore.Encoder {
	var encodeConfig zapcore.EncoderConfig
	appConfig := global.AppConfig

	if appConfig.ServerConfig.Profile == "prod" {
		encodeConfig = zap.NewProductionEncoderConfig()
		// 1716714967.877995 -> 2024-05-26T16:16:07.877+0700
		encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		// ts -> time
		encodeConfig.TimeKey = "time"
		// msg -> message
		encodeConfig.MessageKey = "message"
		// info -> INFO
		encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		//"caller": main.log.go:24
		encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	} else {
		encodeConfig = zap.NewDevelopmentEncoderConfig()
		encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encodeConfig.LevelKey = "level"
		encodeConfig.CallerKey = "caller"
		encodeConfig.MessageKey = "message"
	}

	return zapcore.NewJSONEncoder(encodeConfig)
}

func writeSync() zapcore.WriteSyncer {
	loggerConfig := global.AppConfig.LoggerConfig
	logger := lumberjack.Logger{
		Filename:   loggerConfig.LogDir,
		MaxSize:    loggerConfig.MaxSize,
		MaxBackups: loggerConfig.MaxBackups,
		MaxAge:     loggerConfig.MaxAge,
		Compress:   loggerConfig.Compress,
	}

	return zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		zapcore.AddSync(&logger),
	)
}

func getLogLevel(logger *zap.Logger, level zapcore.Level, msg string, args ...interface{}) {
	switch level {
	case zapcore.InfoLevel:
		logger.Info(msg, zap.Any("args", args))
	case zapcore.WarnLevel:
		logger.Warn(msg, zap.Any("args", args))
	case zapcore.ErrorLevel:
		logger.Error(msg, zap.Any("args", args))
	default:
		logger.Info(msg, zap.Any("args", args))
	}
}
