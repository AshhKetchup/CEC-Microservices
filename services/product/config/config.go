package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port    string
	DataDir string
}

func LoadConfig() *Config {
	viper.AutomaticEnv()
	viper.SetDefault("PORT", "50052")
	viper.SetDefault("DATA_DIR", "./data")

	return &Config{
		Port:    viper.GetString("PORT"),
		DataDir: viper.GetString("DATA_DIR"),
	}
}
