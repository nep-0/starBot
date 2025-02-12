package config

import "github.com/spf13/viper"

type Config struct {
	Openai struct {
		ApiKey  string
		BaseURL string
	}
	Zilliz struct {
		Url         string
		BearerToken string
	}
	OneBot struct {
		QQ       string
		Nickname string
		Listen   string
		Endpoint string
	}
	Static struct {
		VvRoot string
	}
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
