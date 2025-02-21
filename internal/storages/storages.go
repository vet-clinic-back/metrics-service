package storages

import (
	"context"
	"github.com/vet-clinic-back/metrics-service/internal/adapters"
	"github.com/vet-clinic-back/metrics-service/internal/domains"
)

type MetricStorage interface {
	InsertMetrics(context.Context, domains.Metrics) error
	GetMetrics(context.Context, domains.MetricsFilters) ([]domains.Metrics, error)
}

type Storage struct {
	MetricStorage MetricStorage
}

func MustNew(a *adapters.Adapters) *Storage {
	return &Storage{
		a.Postgres,
	}
}
