package main

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	ApiConfig
	DatabaseConfig
}

type ApiConfig struct {
	Host string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

func NewApiConfig() (*AppConfig, error) {
	viper.SetConfigFile("./config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	apiConfig := AppConfig{}
	apiConfig.init()

	return &apiConfig, nil
}

func (ac *AppConfig) init() {

	ac.ApiConfig.Host = viper.GetString("api.host")

}
