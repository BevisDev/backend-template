package startup

import "github.com/BevisDev/backend-template/src/main/helper/logger"

func startLogger(state string) {
	logger.NewAppLogger()
	logger.NewRrLogger()
	logger.Info(state, "LOGGER is started successfully {}", true)
}
