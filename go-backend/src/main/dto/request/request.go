package request

import "time"

type Data struct {
	State     string      `json:"state,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	RequestAt *time.Time  `json:"request_at,omitempty"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
