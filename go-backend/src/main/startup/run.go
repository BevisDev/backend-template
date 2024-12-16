package startup

import (
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/logger"
)

func Run() {
	// load configuration
	LoadConfig()
	logger.NewLogger(config.AppConfig{})
	// server.SetTrustedProxies(config.AppConfig.ServerConfig.TrustedProxies[:])
	// server.Run(":8081")
}
