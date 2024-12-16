package startup

import (
	"log"

	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("../../")
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")

	// read config
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error can not read configuration %v", err)
	}

	var appConfig config.AppConfig

	err = viper.Unmarshal(&appConfig)
	if err != nil {
		log.Fatalf("Error decode config into struct, %v", err)
	}

	log.Printf("Load configuration profile %v successful", appConfig.ServerConfig.Profile)
	log.Printf("Welcome %v version %v ", appConfig.ServerConfig.Name, appConfig.ServerConfig.Version)
}
