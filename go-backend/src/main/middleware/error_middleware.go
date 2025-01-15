package middleware

import (
	"github.com/BevisDev/backend-template/src/main/consts"
	"github.com/BevisDev/backend-template/src/main/dto/response"
	"github.com/BevisDev/backend-template/src/main/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			state := helper.GetState(c)
			if err := recover(); err != nil {
				helper.LogError(state, "LogError occurred {}", err)
			}
			response.ServerError(c, consts.ServerError)
			c.Abort()
		}()

		c.Next()

		if len(c.Errors) != 0 {
			err := c.Errors.Last().Err
			response.SetErrMsg(c, http.StatusInternalServerError, consts.ServerError, err.Error())
		}
	}
}
