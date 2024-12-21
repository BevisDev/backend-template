package response

import "time"

type ApiResponse[T any] struct {
	RequestId  string    `json:"request_id"`
	Data       T         `json:"data"`
	RequestAt  time.Time `json:"request_at"`
	ResponseAt time.Time `json:"response_at`
}
