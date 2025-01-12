package logger

func Init() {
	getAppLogger()
	getRrLogger()
}

func getAppLogger() IAppLogger {
	return NewAppLogger()
}

func getRrLogger() IRrLogger {
	return NewRrLogger()
}

func Sync(state string) {
	getAppLogger().Sync(state)
	getAppLogger().Sync(state)
}

func Info(state, msg string, args ...interface{}) {
	getAppLogger().Info(state, msg, args...)
}

func Error(state, msg string, args ...interface{}) {
	getAppLogger().Error(state, msg, args...)
}

func Warn(state, msg string, args ...interface{}) {
	getAppLogger().Warn(state, msg, args...)
}

func Fatal(state string, msg string, args ...interface{}) {
	getAppLogger().Fatal(state, msg, args...)
}

func RequestLogger(req *Request) {
	getRrLogger().RequestLogger(req)
}

func ResponseLogger(resp *Response) {
	getRrLogger().ResponseLogger(resp)
}
