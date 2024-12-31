package middleware

import (
	"bytes"
	"github.com/BevisDev/backend-template/src/main/consts"
	"github.com/BevisDev/backend-template/src/main/dto/request"
	"github.com/BevisDev/backend-template/src/main/dto/response"
	"github.com/BevisDev/backend-template/src/main/helper"
	"github.com/BevisDev/backend-template/src/main/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"strings"
	"time"
)

type ResponseWrapper struct {
	gin.ResponseWriter
	body     *bytes.Buffer
	status   int
	duration float64
}

func (w *ResponseWrapper) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

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

		requestId := c.GetHeader("X-Request-Id")
		if helper.IsNilOrEmpty(requestId) {
			requestId = uuid.NewString()
			c.Writer.Header().Set("X-Request-Id", requestId)
		}

		var req request.Data
		if err := c.ShouldBind(&req); err != nil {
			response.ErrorResponse(c, http.StatusBadRequest, consts.InvalidRequest, "")
			return
		}

		// ignore log some content-type
		ignoreBody := isIgnoreBody(c.Request.Header)

		// log request
		var reqBody string
		if !ignoreBody {
			reqBytes, _ := io.ReadAll(c.Request.Body)
			reqBody = string(reqBytes)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBytes))
		}

		logger.RequestInfo(&logger.RequestLogger{
			RequestId: requestId,
			URL:       c.Request.URL.String(),
			Query:     c.Request.URL.RawQuery,
			Method:    c.Request.Method,
			Header:    c.Request.Header,
			Body:      reqBody,
		})

		// wrap the responseWriter to capture the response body
		respBuffer := &bytes.Buffer{}
		writer := &ResponseWrapper{
			ResponseWriter: c.Writer,
			body:           respBuffer,
		}
		c.Writer = writer

		// process next
		c.Next()

		// log response
		respHeaders := c.Writer.Header()
		ignoreBody = isIgnoreBody(respHeaders)
		duration := time.Since(startTime)

		var respBody string
		if !ignoreBody {
			respBody = writer.body.String()
		}

		logger.ResponseInfo(&logger.ResponseLogger{
			RequestId: requestId,
			Status:    c.Writer.Status(),
			Duration:  duration,
			Header:    respHeaders,
			Body:      respBody,
		})
	}
}

func isIgnoreBody(headers http.Header) bool {
	contentType := headers.Get("Content-Type")
	return strings.HasPrefix(contentType, "image") ||
		strings.HasPrefix(contentType, "video") ||
		strings.HasPrefix(contentType, "audio")
}
