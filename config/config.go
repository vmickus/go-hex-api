package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	ApiConfig      `mapstructure:"api"`
	DatabaseConfig `mapstructure:"database"`
}

type ApiConfig struct {
	Host string `mapstructure:"host"`
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

// NewApiConfig function    read app configuration from file or environment variables and set in ApiConfig struct.
func NewApiConfig() (*AppConfig, error) {
	v := viper.New()
	apiConfig := new(AppConfig)

	v.SetConfigFile("./config.yaml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := v.Unmarshal(apiConfig); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config %w", err)
	}

	fmt.Printf("%+v\n", apiConfig)

	return apiConfig, nil
}
