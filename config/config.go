package config

import "github.com/spf13/viper"

type Config struct {
	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}

	Server struct {
		Port string
	}
}

func LoadConfiguration() (*Config, error) {
	var config *Config

	viper.SetConfigFile("config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
