package response

import (
	"github.com/BevisDev/backend-template/src/main/consts"
	"github.com/BevisDev/backend-template/src/main/helper/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Data struct {
	State      string      `json:"State,omitempty"`
	RequestAt  *time.Time  `json:"request_at,omitempty"`
	IsSuccess  bool        `json:"is_success"`
	Data       interface{} `json:"data,omitempty"`
	Code       int         `json:"code,omitempty"`
	Message    string      `json:"message,omitempty"`
	ResponseAt time.Time   `json:"response_at,omitempty"`
	Error      *Error      `json:"error,omitempty"`
}

type Error struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

func OK(c *gin.Context, data interface{}, code int) {
	c.JSON(http.StatusOK, Data{
		IsSuccess:  true,
		Data:       data,
		Code:       code,
		Message:    consts.Message[code],
		ResponseAt: time.Now(),
	})
}

func ErrorResponse(c *gin.Context, httpCode int, code int, message string) {
	if utils.IsNilOrEmpty(message) {
		message = consts.Message[code]
	}

	c.JSON(httpCode, Data{
		IsSuccess:  false,
		ResponseAt: time.Now(),
		Error: &Error{
			ErrorCode: code,
			Message:   message,
		},
	})
}
