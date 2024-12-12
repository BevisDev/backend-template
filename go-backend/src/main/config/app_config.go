package config

type AppConfig struct {
	ServerConfig ServerConfig `mapstructure:"server"`
}

type ServerConfig struct {
	Profile string `mapstructure:"profile"`
	Port    int    `mapstructure:"port"`
	Version string `mapstructure:"version"`
}
