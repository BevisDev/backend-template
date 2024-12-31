package request

import "time"

type Data struct {
	RequestID string      `json:"request_id,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	RequestAt *time.Time  `json:"request_at,omitempty"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
