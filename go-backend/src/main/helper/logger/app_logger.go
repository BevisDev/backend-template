package logger

import (
	"fmt"
	"github.com/BevisDev/backend-template/src/main/helper/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
	"sync"
)

var (
	appOnce     sync.Once
	appInstance IAppLogger
)

type appLogger struct {
	appLogger *zap.Logger
}

func NewAppLogger() IAppLogger {
	appOnce.Do(func() {
		encoder := getEncoderLog()
		// new appLogger
		appWrite := writeSync(appFilename)
		appCore := zapcore.NewCore(encoder, appWrite, zapcore.InfoLevel)
		logger := zap.New(appCore, zap.AddCaller())
		appInstance = &appLogger{
			appLogger: logger,
		}
	})
	return appInstance
}

func (c *appLogger) Sync(state string) {
	if c.appLogger != nil {
		if err := c.appLogger.Sync(); err != nil {
			c.Error(state, "Error syncing app logger: {}", err)
		}
	}
}

func (c *appLogger) Info(state, msg string, args ...interface{}) {
	c.log(zapcore.InfoLevel, state, msg, args...)
}

func (c *appLogger) Error(state, msg string, args ...interface{}) {
	c.log(zapcore.ErrorLevel, state, msg, args...)
}

func (c *appLogger) Warn(state, msg string, args ...interface{}) {
	c.log(zapcore.WarnLevel, state, msg, args...)
}

func (c *appLogger) Fatal(state, msg string, args ...interface{}) {
	c.log(zapcore.FatalLevel, state, msg, args...)
}

func (c *appLogger) log(level zapcore.Level, state string, msg string, args ...interface{}) {
	// new instance
	if c.appLogger == nil {
		NewAppLogger()
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
	logging := c.appLogger.WithOptions(zap.AddCallerSkip(3))
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
