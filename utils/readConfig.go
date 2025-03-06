package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type DbConfig struct {
	ServerPort   string
	Environtment string
	Host         string
	Port         string
	Username     string
	Password     string
	DbName       string
	DbType       string
	SslMode      string
}

func ReadConfig() *DbConfig {
	viper.SetConfigFile("yaml")
	viper.AddConfigPath("./")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	dbConfig := &DbConfig{}
	dbConfig.ServerPort = viper.GetString("serverPort")
	dbConfig.Environtment = viper.GetString("environment")
	dbConfig.Host = viper.GetString("database." + dbConfig.Environtment + ".host")
	dbConfig.Port = viper.GetString("database." + dbConfig.Environtment + ".port")
	dbConfig.Username = viper.GetString("database." + dbConfig.Environtment + ".username")
	dbConfig.Password = viper.GetString("database." + dbConfig.Environtment + ".password")
	dbConfig.DbName = viper.GetString("database." + dbConfig.Environtment + ".dbname")
	dbConfig.DbType = viper.GetString("database." + dbConfig.Environtment + ".database")
	dbConfig.SslMode = viper.GetString("database." + dbConfig.Environtment + ".sslmode")

	return dbConfig
}
