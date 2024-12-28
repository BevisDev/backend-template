package middlewares

import (
	"bytes"
	"github.com/BevisDev/backend-template/src/main/consts"
	"github.com/BevisDev/backend-template/src/main/dto/request"
	"github.com/BevisDev/backend-template/src/main/dto/response"
	"github.com/BevisDev/backend-template/src/main/helper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		accessToken := c.GetHeader("Authorization")
		if helper.IsNilOrEmpty(accessToken) {
			response.ErrorResponse(c, http.StatusUnauthorized, consts.InvalidAccessToken, "")
			c.Abort()
			return
		}

		signature := c.GetHeader("X-Signature")
		if helper.IsNilOrEmpty(signature) {
			response.ErrorResponse(c, http.StatusUnauthorized, consts.InvalidSignature, "")
			c.Abort()
			return
		}

		requestID := c.GetHeader("X-Request-ID")
		if helper.IsNilOrEmpty(requestID) {
			requestID = uuid.NewString()
			c.Writer.Header().Set("X-Request-ID", requestID)
		}

		var req request.Data
		if err := c.ShouldBind(&req); err != nil {
			response.ErrorResponse(c, http.StatusBadRequest, consts.InvalidRequest, "")
			return
		}

		reqBody, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))

		c.Next()

		duration := time.Since(startTime)

	}
}
