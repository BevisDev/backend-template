package middleware

import (
	"github.com/BevisDev/backend-template/src/main/consts"
	"github.com/BevisDev/backend-template/src/main/dto/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) != 0 {
			err := c.Errors.Last().Err
			response.ErrorResponse2(c, http.StatusInternalServerError, consts.ServerError, err.Error())
		}
	}
}
