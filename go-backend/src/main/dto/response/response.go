package response

import (
	"github.com/BevisDev/backend-template/src/main/consts"
	"github.com/BevisDev/backend-template/src/main/helper"
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
	return helper.TimeToString(time.Now(), consts.DATETIME_NO_TZ)
}

func OK(c *gin.Context, data interface{}, code int) {
	c.JSON(http.StatusOK, &Data{
		State:      helper.GetState(c),
		IsSuccess:  true,
		Data:       data,
		Code:       code,
		Message:    consts.Message[code],
		ResponseAt: getResponseAt(),
	})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, &Data{
		State:      helper.GetState(c),
		IsSuccess:  true,
		Data:       data,
		Code:       consts.Created,
		Message:    consts.Message[consts.Created],
		ResponseAt: getResponseAt(),
	})
}

func Unauthorized(c *gin.Context, code int) {
	c.JSON(http.StatusUnauthorized, &DataError{
		State:      helper.GetState(c),
		IsSuccess:  false,
		ResponseAt: getResponseAt(),
		Error: &Error{
			ErrorCode: code,
			Message:   consts.Message[code],
		},
	})
}

func BadRequest(c *gin.Context, code int) {
	c.JSON(http.StatusBadRequest, &DataError{
		State:      helper.GetState(c),
		IsSuccess:  false,
		ResponseAt: getResponseAt(),
		Error: &Error{
			ErrorCode: code,
			Message:   consts.Message[code],
		},
	})
}

func ServerError(c *gin.Context, code int) {
	c.JSON(http.StatusInternalServerError, &DataError{
		State:      helper.GetState(c),
		IsSuccess:  false,
		ResponseAt: getResponseAt(),
		Error: &Error{
			ErrorCode: code,
			Message:   consts.Message[code],
		},
	})
}

func SetError(c *gin.Context, httpCode, code int) {
	c.JSON(httpCode, &DataError{
		State:      helper.GetState(c),
		IsSuccess:  false,
		ResponseAt: getResponseAt(),
		Error: &Error{
			ErrorCode: code,
			Message:   consts.Message[code],
		},
	})
}

func SetErrMsg(c *gin.Context, httpCode, code int, message string) {
	c.JSON(httpCode, &DataError{
		State:      helper.GetState(c),
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
		State:      helper.GetState(c),
		IsSuccess:  false,
		ResponseAt: getResponseAt(),
		Error: &Error{
			ErrorCode: consts.ServerTimeout,
			Message:   consts.Message[consts.ServerTimeout],
		},
	})
}
