package rest

import (
	"github.com/BevisDev/backend-template/src/main/common/utils"
	"github.com/BevisDev/backend-template/src/main/infrastructure/config"
	"github.com/BevisDev/backend-template/src/main/infrastructure/logger"
	"net/http"
	"net/url"
	"sync"
)

var (
	restOnce      sync.Once
	httpClient    *http.Client
	clientTimeout int
)

type RestRequest struct {
	State    string
	URL      string
	Params   map[string]any
	Header   map[string]string
	Body     any
	BodyForm url.Values
	Result   any
}

type RestResponse struct {
	StatusCode int
	Header     http.Header
	Body       string
	HasError   bool
	IsTimeout  bool
	Error      error
}

func InitRestClient(state string) *http.Client {
	cf := config.AppConfig
	if utils.IsNilOrEmpty(cf) ||
		utils.IsNilOrEmpty(cf.ServerConfig) {
		logger.Fatal(state, "Error appConfig is not initialized")
		return nil
	}
	restOnce.Do(func() {
		httpClient = &http.Client{}
		clientTimeout = config.AppConfig.ServerConfig.ClientTimeout
	})
	return httpClient
}
