package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

func GetConfig(dest interface{}, module, kind string) error {
	var (
		err     error
		profile = os.Getenv("GO_PROFILE")
		path    = os.Getenv("GO_CONFIG_" + strings.ToUpper(module))
	)
	if profile == "" {
		profile = "dev" // set default
	}
	if path == "" {
		path = "D:\\Project\\backend\\backend-template\\go-backend\\config\\" + module
	}
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName(profile)
	v.SetConfigType(kind)

	// read config
	if err = v.ReadInConfig(); err != nil {
		log.Fatalf("Error can not read configuration %v", err)
	}
	if err = v.Unmarshal(&dest); err != nil {
		log.Fatalf("Error decode config into struct, %v", err)
	}
	return err
}
