package response

import (
	"github.com/BevisDev/backend-template/src/main/consts"
	"github.com/BevisDev/backend-template/src/main/helper/datetime"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Data struct {
	State      string      `json:"State,omitempty"`
	IsSuccess  bool        `json:"is_success"`
	Data       interface{} `json:"data,omitempty"`
	Code       int         `json:"code,omitempty"`
	Message    string      `json:"message,omitempty"`
	ResponseAt string      `json:"response_at,omitempty"`
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
		ResponseAt: datetime.ToString(time.Now(), consts.DATETIME_NO_TZ),
	})
}

func SetError(c *gin.Context, httpCode, code int) {
	message := consts.Message[code]

	c.JSON(httpCode, Data{
		IsSuccess:  false,
		ResponseAt: datetime.ToString(time.Now(), consts.DATETIME_NO_TZ),
		Error: &Error{
			ErrorCode: code,
			Message:   message,
		},
	})
}

func SetErrorMsg(c *gin.Context, httpCode, code int, message string) {
	c.JSON(httpCode, Data{
		IsSuccess:  false,
		ResponseAt: datetime.ToString(time.Now(), consts.DATETIME_NO_TZ),
		Error: &Error{
			ErrorCode: code,
			Message:   message,
		},
	})
}
