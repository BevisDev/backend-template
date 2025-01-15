package middleware

import (
	"github.com/BevisDev/backend-template/src/main/consts"
	"github.com/BevisDev/backend-template/src/main/dto/response"
	"github.com/BevisDev/backend-template/src/main/helper"
	"github.com/gin-gonic/gin"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		if helper.IsNilOrEmpty(accessToken) {
			response.Unauthorized(c, consts.InvalidAccessToken)
			c.Abort()
			return
		}

		signature := c.GetHeader("signature")
		if helper.IsNilOrEmpty(signature) {
			response.Unauthorized(c, consts.InvalidSignature)
			c.Abort()
			return
		}

		c.Next()
	}
}
