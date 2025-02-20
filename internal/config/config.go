package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	logging "github.com/vet-clinic-back/metrics-service/pkg/logger"
)

type Config struct {
	Postgres Postgres `yaml:"postgres"`
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST" validate:"required"`
	Port     string `env:"POSTGRES_PORT" validate:"required"`
	User     string `env:"POSTGRES_USER" validate:"required"`
	Password string `env:"POSTGRES_PASSWORD" validate:"required"`
	Database string `env:"POSTGRES_DB" validate:"required"`
}

// MustConfigure is a configurator for a config.
// @params: configPath is a path to config. is used by tests from any package with configuration.
func MustConfigure(configPath string) Config {
	config := &Config{}
	log := logging.GetLogger().WithField("op", "config.MustConfigure")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	
	err := cleanenv.ReadEnv(config)
	if err != nil {
		description, _ := cleanenv.GetDescription(config, nil)
		log.Fatalf("failed to read config %s %s", description, err)
	}

	validate := validator.New()

	if err = validate.Struct(*config); err != nil {
		log.Fatal(err)
	}

	return *config
}
