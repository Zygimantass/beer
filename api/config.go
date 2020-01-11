package api

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port int
}

func InitConfig() (*Config, error) {
	config := &Config{
		Port: viper.GetInt("port"),
	}

	if config.Port == 0 {
		config.Port = 7777
	}

	if config.Port < 1 || config.Port > 65535 {
		return nil, fmt.Errorf("API config port has to be inbetween 1-65535")
	}

	return config, nil
}
