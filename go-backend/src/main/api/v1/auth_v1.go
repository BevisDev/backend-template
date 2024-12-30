package v1

import (
	"github.com/BevisDev/backend-template/src/main/controller"
	"github.com/gin-gonic/gin"
)

func AuthApiGroup(group *gin.RouterGroup) {
	group.GET("/signin", controller.SignIn)
	group.POST("/signup", controller.SignUp)
}
