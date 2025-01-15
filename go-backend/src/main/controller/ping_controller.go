package controller

import (
	"github.com/BevisDev/backend-template/src/main/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingController struct {
	pingService service.IPingService
}

func NewPingController(
	pingService service.IPingService,
) *PingController {
	return &PingController{
		pingService: pingService,
	}
}

// Ping godoc
// @Summary Ping System API
// @Description Check health system
// @Tags Check Health
// @Produce  json
// @Success 200 {object} map[string]interface{} "Successful"
// @Router /ping [get]
func (p *PingController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// PingDB godoc
// @Summary Ping DB API
// @Description Check health DB System
// @Tags Check Health
// @Produce  json
// @Success 200 {object} map[string]interface{} "Successful"
// @Router /db [get]
func (p *PingController) PingDB(c *gin.Context) {
	c.JSON(http.StatusOK, p.pingService.PingDB(c))
}

// PingRedis godoc
// @Summary Ping Redis API
// @Description Check health Redis System
// @Tags Check Health
// @Produce  json
// @Success 200 {object} map[string]interface{} "Successful"
// @Router /redis [get]
func (p *PingController) PingRedis(c *gin.Context) {
	c.JSON(http.StatusOK, p.pingService.PingRedis(c))
}
