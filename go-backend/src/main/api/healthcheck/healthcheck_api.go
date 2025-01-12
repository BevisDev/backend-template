package healthcheck

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
