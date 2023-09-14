package config

import "github.com/spf13/viper"

func init() {
	viper.AutomaticEnv()

	viper.SetDefault("ENVIRONMENT", "local")
}

type ServerConfig struct {
	Environment string
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Environment: viper.GetString("ENVIRONMENT"),
	}
}

func (sc *ServerConfig) IsLocal() bool {
	return sc.Environment == "local"
}
