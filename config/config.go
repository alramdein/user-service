package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetConf() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w ", err))
	}
}

func HTTPHost() string {
	return viper.GetString("http.host")
}

func HTTPPort() string {
	return viper.GetString("http.port")
}

func GRPCPort() string {
	return viper.GetString("grpc.port")
}

func DBHost() string {
	return viper.GetString("database.host")
}

func DBPort() string {
	return viper.GetString("database.port")
}

func DBUsername() string {
	return viper.GetString("database.username")
}

func DBPassword() string {
	return viper.GetString("database.password")
}

func DBName() string {
	return viper.GetString("database.dbname")
}
