package middleware

import (
	"context"
	"github.com/BevisDev/backend-template/src/main/consts"
	"github.com/BevisDev/backend-template/src/main/dto/response"
	"github.com/BevisDev/backend-template/src/main/helper/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		if utils.IsNilOrEmpty(accessToken) {
			response.ErrorResponse(c, http.StatusUnauthorized, consts.InvalidAccessToken, "")
			c.Abort()
			return
		}

		signature := c.GetHeader("signature")
		if utils.IsNilOrEmpty(signature) {
			response.ErrorResponse(c, http.StatusUnauthorized, consts.InvalidSignature, "")
			c.Abort()
			return
		}

		state := c.GetHeader("state")
		if utils.IsNilOrEmpty(state) {
			state = utils.GenUUID()
		}
		// write state in header response
		c.Writer.Header().Set("state", state)

		// store state in context
		ctx := context.WithValue(c.Request.Context(), "state", state)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
