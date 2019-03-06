package config

import "github.com/spf13/viper"

type AppConfig struct {
	appPort int
}

type DatabaseConfig struct {
	dbUser     string
	dbPassword string
	dbName     string
	dbHost     string
}

var appConfig *AppConfig
var dbConfig *DatabaseConfig

func Load() {
	viper.AddConfigPath("./")
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	appConfig = &AppConfig{
		appPort: viper.GetInt("APP_PORT"),
	}

	dbConfig = &DatabaseConfig{
		dbHost: viper.GetString("DB_HOST"),
		dbUser: viper.GetString("DB_USER"),
		dbPassword: viper.GetString("DB_PASSWORD"),
		dbName: viper.GetString("DB_NAME"),
	}

}

func AppPort() int {
	return appConfig.appPort
}

func DatabaseUser() string {
	return dbConfig.dbUser
}

func DatabasePassword() string {
	return dbConfig.dbPassword
}

func DatabaseName() string {
	return dbConfig.dbName
}

func DatabaseHost() string {
	return dbConfig.dbHost
}
