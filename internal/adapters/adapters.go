package adapters

import (
	"github.com/vet-clinic-back/metrics-service/internal/adapters/postgres"
	"github.com/vet-clinic-back/metrics-service/internal/config"
)

type Adapters struct {
	Postgres *postgres.Postgres
}

func NewAdapters(cfg config.Config) *Adapters {
	return &Adapters{
		Postgres: postgres.MustNew(cfg.Postgres),
	}
}
