package logger

type (
	IAppLogger interface {
		Sync(state string)
		Info(state, msg string, args ...interface{})
		Error(state, msg string, args ...interface{})
		Warn(state, msg string, args ...interface{})
		Fatal(state, msg string, args ...interface{})
	}

	IRrLogger interface {
		Sync(state string)
		RequestLogger(req *Request)
		ResponseLogger(resp *Response)
	}

	IExtLogger interface {
		Sync(state string)
		RequestLogger(req *Request)
		ResponseLogger(resp *Response)
	}
)
