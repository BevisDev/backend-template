package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func SignUp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Register successful"})
}
