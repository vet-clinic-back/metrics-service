package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	TCP  TCPConfig
	HTTP HTTPConfig
	DB   DBConfig
}

type TCPConfig struct {
	Addr string
}

type HTTPConfig struct {
	Addr string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
