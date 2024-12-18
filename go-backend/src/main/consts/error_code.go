package consts

const (
	OK             = 1000
	Created        = 1001
	InvalidRequest = 4000
	Unauthorized   = 4002
	NotFound       = 4004
	ServerError    = 5000
)

var message = map[int]string{
	OK:             "Success",
	Created:        "Created",
	InvalidRequest: "Invalid Request",
	ServerError:    "Server has error",
}
