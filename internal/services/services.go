package services

import (
	"context"
	"github.com/vet-clinic-back/metrics-service/internal/domains"
	"github.com/vet-clinic-back/metrics-service/internal/services/metricservice"
	"github.com/vet-clinic-back/metrics-service/internal/storages"
)

type Metrics interface {
	InsertMetrics(context.Context, domains.Metrics) error
	GetMetrics(context.Context, domains.MetricsFilters) ([]domains.Metrics, error)
}

type Service struct {
	Metrics Metrics
}

func MustNew(s *storages.Storage) *Service {
	return &Service{
		Metrics: metricservice.NewMetricService(s),
	}
}
