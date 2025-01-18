package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterAPIs(r *gin.RouterGroup) {
	AuthAPI(r)
}
