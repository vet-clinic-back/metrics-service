package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	TCP  TCPConfig  `mapstructure:"tcp"`
	HTTP HTTPConfig `mapstructure:"http"`
	DB   DBConfig   `mapstructure:"db"`
}

type TCPConfig struct {
	Addr string `mapstructure:"addr"`
}

type HTTPConfig struct {
	Addr string `mapstructure:"addr"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// Для отладки: список всех считанных ключей
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic occurred:", err)
		}
		fmt.Println("All settings:", viper.AllSettings())
	}()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal error: %v", err)
	}

	return &cfg, nil
}
