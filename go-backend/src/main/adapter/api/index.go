package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterAPIs(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		AuthAPI(v1)
	}
}
