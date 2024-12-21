package response

import "time"

type Response struct {
	RequestID  string      `json:"request_id,omitempty"`
	RequestAt  time.Time   `json:"request_at,omitempty"`
	IsSuccess  bool        `json:"is_success"`
	Data       interface{} `json:"data,omitempty"`
	Code       string      `json:"code,omitempty"`
	Message    string      `json:"message,omitempty"`
	ResponseAt time.Time   `json:"response_at,omitempty"`
}

func OK(data interface{}) Response {
	return Response{
		IsSuccess:  true,
		Data:       data,
		ResponseAt: time.Now(),
	}
}

func Error() Response {
	return Response{
		IsSuccess:  true,
		ResponseAt: time.Now(),
	}
}
