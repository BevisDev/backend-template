package middleware

import "github.com/gin-gonic/gin"

func RegisterMiddleware(r *gin.Engine) {
	r.Use(LoggerHandler())
	r.Use(AuthHandler())
	r.Use(ErrorHandler())
}
