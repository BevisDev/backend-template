package config

type AppConfig struct {
	ServerConfig   ServerConfig   `mapstructure:"server"`
	DatabaseConfig DatabaseConfig `mapstructure:"databases"`
	LoggerConfig   LoggerConfig   `mapstructure:"logger"`
}

type ServerConfig struct {
	Name           string   `mapstructure:"name"`
	Profile        string   `mapstructure:"profile"`
	Port           int      `mapstructure:"port"`
	Version        string   `mapstructure:"version"`
	TrustedProxies []string `mapstructure:"trustedProxies"`
}

type DatabaseConfig struct {
	MSSQLConfig      MSSQLConfig      `mapstructure:"mssql"`
	PostgreSQLConfig PostgreSQLConfig `mapstructure:"postgresql"`
}

type MSSQLConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Schema   string `mapstructure:"schema"`
}

type PostgreSQLConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Schema   string `mapstructure:"schema"`
}

type LoggerConfig struct {
	LogLevel   string `mapstructure:"logLevel"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge"`
	Compress   bool   `mapstructure:"compress"`
}
