package logger

import (
	"os"

	"github.com/BevisDev/backend-template/src/main/global"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Zap struct {
	*zap.Logger
}

func NewLogger() *Zap {
	loggerConfig := global.AppConfig.LoggerConfig

	// debug-> info-> warn ->error->fatal->panic
	var level zapcore.Level
	switch loggerConfig.LogLevel {
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
		Filename:   loggerConfig.Filename,
		MaxSize:    loggerConfig.MaxSize,
		MaxBackups: loggerConfig.MaxBackups,
		MaxAge:     loggerConfig.MaxAge,
		Compress:   loggerConfig.Compress,
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level)
	return &Zap{zap.New(core, zap.AddCaller())}
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// 1716714967.877995 -> 2024-05-26T16:16:07.877+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> time
	encodeConfig.TimeKey = "time"
	// info -> INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//"caller":"cli/main.log.go:24
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}
