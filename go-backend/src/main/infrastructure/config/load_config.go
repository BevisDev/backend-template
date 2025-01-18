package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
)

var (
	onceConfig sync.Once
	AppConfig  *Config
)

func LoadConfig() *Config {
	onceConfig.Do(func() {
		profile := os.Getenv("GO_PROFILE")
		if profile == "" {
			profile = "dev" // set default
		}

		v := viper.New()
		v.AddConfigPath("./src/resources")
		v.SetConfigName(profile)
		v.SetConfigType("yaml")

		// read config
		if err := v.ReadInConfig(); err != nil {
			log.Fatalf("Error can not read configuration %v", err)
		}

		if err := v.Unmarshal(&AppConfig); err != nil {
			log.Fatalf("Error decode config into struct, %v", err)
		}

		serverConfig := AppConfig.ServerConfig
		log.Println("================================")
		log.Printf("Load configuration profile %s successful", serverConfig.Profile)
		log.Printf("Welcome to %s version %s ", serverConfig.Name, serverConfig.Version)
		log.Println("================================")
	})
	return AppConfig
}
