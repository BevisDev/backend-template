package logger

func Info(state, msg string, args ...interface{}) {
	NewAppLogger().Info(state, msg, args...)
}

func Error(state, msg string, args ...interface{}) {
	NewAppLogger().Error(state, msg, args...)
}

func Warn(state, msg string, args ...interface{}) {
	NewAppLogger().Warn(state, msg, args...)
}

func Panic(state string, msg string, args ...interface{}) {
	NewAppLogger().Panic(state, msg, args...)
}

func RequestLogger(req *Request) {
	NewRrLogger().RequestLogger(req)
}

func ResponseLogger(resp *Response) {
	NewRrLogger().ResponseLogger(resp)
}

func Sync(state string) {
	NewAppLogger().Sync(state)
	NewRrLogger().Sync(state)
}
