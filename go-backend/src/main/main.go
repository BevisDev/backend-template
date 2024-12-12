package main

import (
	"net/http"

	"github.com/BevisDev/backend-template/src/main/initialize"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	initialize.LoadConfig()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.SetTrustedProxies([]string{"127.0.0.1"})
	server.Run(":8081") // change port is here, default 8080
}
