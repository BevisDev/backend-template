package logger

import (
	"fmt"
	"os"
	"strings"
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
	sugarLogger *zap.SugaredLogger
}

func initLogger() {
	once.Do(func() {
		encoder := getEncoderLog()
		writeSync := writeSync()
		core := zapcore.NewCore(encoder, writeSync, zapcore.InfoLevel)
		logger = &LoggerWrapper{
			sugarLogger: zap.New(core, zap.AddCaller()).Sugar(),
		}
	})
}

func getEncoderLog() zapcore.Encoder {
	var encodeConfig zapcore.EncoderConfig
	appConfig := global.AppConfig
	isDev := appConfig.ServerConfig.Profile == "dev"

	// handle profile prod
	if appConfig.ServerConfig.Profile == "prod" {
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

	if isDev {
		return zapcore.NewConsoleEncoder(encodeConfig)
	}

	return zapcore.NewJSONEncoder(encodeConfig)
}

func writeSync() zapcore.WriteSyncer {
	loggerConfig := global.AppConfig.LoggerConfig
	logger := lumberjack.Logger{
		Filename:   loggerConfig.LogDir + "/log1.log",
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

func log(level zapcore.Level, msg string, args ...interface{}) {
	var message string

	// formater message
	if len(args) != 0 {
		message = formatMessage(msg, args...)
	} else {
		message = msg
	}

	switch level {
	case zapcore.InfoLevel:
		logger.sugarLogger.Info(message)
	case zapcore.WarnLevel:
		logger.sugarLogger.Error(message)
	case zapcore.ErrorLevel:
		logger.sugarLogger.Warn(message)
	default:
		logger.sugarLogger.Info(message)
	}
}

func formatMessage(msg string, args ...interface{}) string {
	var message string
	if !strings.Contains(msg, "%") {
		message = strings.ReplaceAll(msg, "{}", "%v")
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
