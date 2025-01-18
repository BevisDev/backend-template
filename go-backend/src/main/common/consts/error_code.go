package consts

const (
	// ServerError server error from 1000 to 1999
	ServerError        = 1000
	ServerTimeout      = 1001
	ServerDown         = 1002
	ServiceUnavailable = 1003

	// OK code success from 2000 to 2999
	OK      = 2000
	Created = 2001

	// InvalidRequest client error from 3000 to ...
	InvalidRequest      = 3000
	InvalidCredentials  = 3001
	NotAuthorizedAccess = 3002
	InvalidAccessToken  = 3003
	InvalidSignature    = 3004
	NotFound            = 3005
)

var Message = map[int]string{
	// message server error
	ServerError:        "Server has error",
	ServerTimeout:      "Server gateway is timed out",
	ServerDown:         "Server is down or under maintenance",
	ServiceUnavailable: "The service is temporarily unavailable",

	// message success
	OK:      "Success",
	Created: "Created",

	// message client error
	InvalidRequest:      "Invalid Request",
	InvalidCredentials:  "Security credentials is incorrect",
	NotAuthorizedAccess: "You are not authorized to access this resource",
	InvalidAccessToken:  "Access token is invalid",
	InvalidSignature:    "Signature is invalid",
	NotFound:            "Page Not Found",
}
