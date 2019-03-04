package config

import "github.com/spf13/viper"

type Config struct {
	appPort int
}

var appConfig *Config

func Load() {
	viper.AddConfigPath("./")
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	appConfig = &Config{
		appPort: viper.GetInt("APP_PORT"),
	}
}

func AppPort() int {
	return appConfig.appPort
}
