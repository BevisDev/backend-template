package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignIn godoc
// @Summary Sign In API
// @Description Returns a greeting message
// @Tags greeting
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{} "Successful"
// @Failure 400 {object} map[string]interface{} "Client Error"
// @Failure 500 {object} map[string]interface{} "Server Error"
// @Router /signin [get]
// @Security AccessTokenAuth
func SignIn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func SignUp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Register successful"})
}
