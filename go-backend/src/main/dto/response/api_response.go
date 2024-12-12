package response

import "time"

type ApiResponse[T any] struct {
	requestId  string    `json:"request_id"`
	data       T         `json:"data"`
	requestAt  time.Time `json:"request_at"`
	responseAt time.Time `json:"response_at`
}
