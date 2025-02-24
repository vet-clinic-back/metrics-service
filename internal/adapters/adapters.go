package adapters

import (
	"github.com/vet-clinic-back/metrics-service/internal/adapters/httpadapter"
	"github.com/vet-clinic-back/metrics-service/internal/adapters/postgres"
	"github.com/vet-clinic-back/metrics-service/internal/adapters/tcpadapter"
	"github.com/vet-clinic-back/metrics-service/internal/config"
)

type Adapters struct {
	Postgres    *postgres.Postgres
	HTTPAdapter *httpadapter.HTTPAdapter
	TCPAdapter  *tcpadapter.TCPAdapter
}

func NewAdapters(cfg config.Config) *Adapters {
	return &Adapters{
		Postgres:    postgres.MustNew(cfg.Postgres),
		HTTPAdapter: httpadapter.New(cfg.HTTPConfig),
		TCPAdapter:  tcpadapter.NewTCPAdapter(cfg.TCPConfig),
	}
}
