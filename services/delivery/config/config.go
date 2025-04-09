package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port             string
	DataDir          string
	OrderServiceAddr string
}

func LoadConfig() *Config {
	viper.AutomaticEnv()
	viper.SetDefault("PORT", "50054")
	viper.SetDefault("DATA_DIR", "./data")
	viper.SetDefault("ORDER_SERVICE_ADDR", "order-service:50053")

	return &Config{
		Port:             viper.GetString("PORT"),
		DataDir:          viper.GetString("DATA_DIR"),
		OrderServiceAddr: viper.GetString("ORDER_SERVICE_ADDR"),
	}
}
