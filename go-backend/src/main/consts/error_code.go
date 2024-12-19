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
	Unauthorized:   "You are not Authorized",
	NotFound:       "Not Found",
	ServerError:    "Server has error",
}
