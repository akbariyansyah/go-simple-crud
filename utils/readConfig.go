package utils

import (
	"github.com/spf13/viper"
)

func ReadConfig(key string) string {
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		panic("Failed read config file")
	}
	configValue := viper.GetString(key)
	return configValue
}
