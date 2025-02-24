package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	logging "github.com/vet-clinic-back/metrics-service/pkg/logger"
)

type Config struct {
	HTTPConfig HTTPConfig
	Postgres   Postgres
	TCPConfig  TCPConfig
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST" validate:"required"`
	Port     string `env:"POSTGRES_PORT" validate:"required"`
	User     string `env:"POSTGRES_USER" validate:"required"`
	Password string `env:"POSTGRES_PASSWORD" validate:"required"`
	Database string `env:"POSTGRES_DB" validate:"required"`
}

type HTTPConfig struct {
	Port         string   `env:"METRICS_HTTP_PORT" env-default:"8080"`
	AllowOrigins []string `env:"METRICS_ALLOW_ORIGINS"`
}

type TCPConfig struct {
	Port string `env:"METRICS_TCP_LISTEN_PORT" env-default:"8080"`
}

// MustConfigure is a configurator for a config.
func MustConfigure() Config {
	config := &Config{}
	log := logging.GetLogger().WithField("op", "config.MustConfigure")

	//if err := godotenv.Load(); err != nil {
	//	log.Fatal("Error loading .env file")
	//}

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
