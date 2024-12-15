package logger

import (
	"os"

	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config config.AppConfig) *LoggerZap {
	logLevel := config.LoggerConfig.LogLevel //"debug"
	// debug-> info-> warn ->error->fatal->panic

	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.LoggerConfig.Filename,
		MaxSize:    config.LoggerConfig.MaxSize,
		MaxBackups: config.LoggerConfig.MaxBackups,
		MaxAge:     config.LoggerConfig.MaxAge,
		Compress:   config.LoggerConfig.Compress,
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level)
	return &LoggerZap{zap.New(core, zap.AddCaller())}
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// 1716714967.877995 -> 2024-05-26T16:16:07.877+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> time
	encodeConfig.TimeKey = "time"
	// from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//"caller":"cli/main.log.go:24
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}
