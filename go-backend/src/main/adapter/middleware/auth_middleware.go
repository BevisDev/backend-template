package middleware

import (
	"github.com/BevisDev/backend-template/src/main/adapter/dto/response"
	"github.com/BevisDev/backend-template/src/main/common/consts"
	"github.com/BevisDev/backend-template/src/main/common/utils"
	"github.com/gin-gonic/gin"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		if utils.IsNilOrEmpty(accessToken) {
			response.Unauthorized(c, consts.InvalidAccessToken)
			c.Abort()
			return
		}

		signature := c.GetHeader("signature")
		if utils.IsNilOrEmpty(signature) {
			response.Unauthorized(c, consts.InvalidSignature)
			c.Abort()
			return
		}

		c.Next()
	}
}
