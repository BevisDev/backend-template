package consts

const (
	// ServerError server error
	ServerError        = 5000
	ServerTimeout      = 5001
	ServerDown         = 5002
	ServiceUnavailable = 5003

	// OK code success
	OK      = 1000
	Created = 1001

	// InvalidRequest client error
	InvalidRequest      = 4000
	InvalidCredentials  = 4001
	NotAuthorizedAccess = 4002
	InvalidAccessToken  = 4003
	InvalidSignature    = 4004
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
}
