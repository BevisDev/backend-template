package response

import (
	consts2 "github.com/BevisDev/backend-template/src/main/common/consts"
	"github.com/BevisDev/backend-template/src/main/common/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Data struct {
	State      string      `json:"state,omitempty" example:"8137ce10-305b-42f5-8f14-9c48dd6f23f0"`
	IsSuccess  bool        `json:"is_success" example:"true"`
	Data       interface{} `json:"data,omitempty"`
	Code       int         `json:"code,omitempty" example:"2000"`
	Message    string      `json:"message,omitempty" example:"Success"`
	ResponseAt string      `json:"response_at,omitempty" example:"2025-01-14 16:44:47.510"`
}

type DataError struct {
	State      string `json:"state,omitempty" example:"8137ce10-305b-42f5-8f14-9c48dd6f23f0"`
	IsSuccess  bool   `json:"is_success" example:"false"`
	ResponseAt string `json:"response_at,omitempty" example:"2025-01-14 16:44:47.510"`
	Error      *Error `json:"error,omitempty"`
}

type Error struct {
	ErrorCode int    `json:"error_code" example:"3000"`
	Message   string `json:"message" example:"Invalid RequestLogger"`
}

func getResponseAt() string {
	return utils.TimeToString(time.Now(), consts2.DATETIME_NO_TZ)
}

func OK(c *gin.Context, data interface{}, code int) {
	c.JSON(http.StatusOK, &Data{
		State:      utils.GetState(c),
		IsSuccess:  true,
		Data:       data,
		Code:       code,
		Message:    consts2.Message[code],
		ResponseAt: getResponseAt(),
	})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, &Data{
		State:      utils.GetState(c),
		IsSuccess:  true,
		Data:       data,
		Code:       consts2.Created,
		Message:    consts2.Message[consts2.Created],
		ResponseAt: getResponseAt(),
	})
}

func Unauthorized(c *gin.Context, code int) {
	c.JSON(http.StatusUnauthorized, &DataError{
		State:      utils.GetState(c),
		IsSuccess:  false,
		ResponseAt: getResponseAt(),
		Error: &Error{
			ErrorCode: code,
			Message:   consts2.Message[code],
		},
	})
}

func BadRequest(c *gin.Context, code int) {
	c.JSON(http.StatusBadRequest, &DataError{
		State:      utils.GetState(c),
		IsSuccess:  false,
		ResponseAt: getResponseAt(),
		Error: &Error{
			ErrorCode: code,
			Message:   consts2.Message[code],
		},
	})
}

func ServerError(c *gin.Context, code int) {
	c.JSON(http.StatusInternalServerError, &DataError{
		State:      utils.GetState(c),
		IsSuccess:  false,
		ResponseAt: getResponseAt(),
		Error: &Error{
			ErrorCode: code,
			Message:   consts2.Message[code],
		},
	})
}

func SetError(c *gin.Context, httpCode, code int) {
	c.JSON(httpCode, &DataError{
		State:      utils.GetState(c),
		IsSuccess:  false,
		ResponseAt: getResponseAt(),
		Error: &Error{
			ErrorCode: code,
			Message:   consts2.Message[code],
		},
	})
}

func SetErrMsg(c *gin.Context, httpCode, code int, message string) {
	c.JSON(httpCode, &DataError{
		State:      utils.GetState(c),
		IsSuccess:  false,
		ResponseAt: getResponseAt(),
		Error: &Error{
			ErrorCode: code,
			Message:   message,
		},
	})
}

func Timeout(c *gin.Context) {
	c.JSON(http.StatusGatewayTimeout, &DataError{
		State:      utils.GetState(c),
		IsSuccess:  false,
		ResponseAt: getResponseAt(),
		Error: &Error{
			ErrorCode: consts2.ServerTimeout,
			Message:   consts2.Message[consts2.ServerTimeout],
		},
	})
}
