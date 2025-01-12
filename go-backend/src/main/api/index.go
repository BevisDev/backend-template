package api

import (
	"github.com/BevisDev/backend-template/src/main/api/auth"
	"github.com/gin-gonic/gin"
)

func APIs(r *gin.RouterGroup) {
	auth.AuthAPI(r)
}
