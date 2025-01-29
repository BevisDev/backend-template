package api

import (
	"github.com/BevisDev/backend-template/src/main/infrastructure/di"
	"github.com/gin-gonic/gin"
)

func AuthAPI(g *gin.RouterGroup) {
	c := di.NewAuthDI()
	g.GET("/signin", c.SignIn)
	g.GET("/signup", c.SignUp)
}
