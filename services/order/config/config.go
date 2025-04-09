package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port               string
	DataDir            string
	ProductServiceAddr string
}

func LoadConfig() *Config {
	viper.AutomaticEnv()
	viper.SetDefault("PORT", "50053")
	viper.SetDefault("DATA_DIR", "./data")
	viper.SetDefault("PRODUCT_SERVICE_ADDR", "product-service:50052")

	return &Config{
		Port:               viper.GetString("PORT"),
		DataDir:            viper.GetString("DATA_DIR"),
		ProductServiceAddr: viper.GetString("PRODUCT_SERVICE_ADDR"),
	}
}
