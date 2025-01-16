package process

import "github.com/gin-gonic/gin"

func ProcessAPIs(g *gin.RouterGroup) {
	p := g.Group("/process")
	{
		p.POST(":name")
	}
}
