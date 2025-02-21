package metricservice

import (
	"context"
	"github.com/vet-clinic-back/metrics-service/internal/domains"
	"github.com/vet-clinic-back/metrics-service/internal/storages"
)

type MetricService struct {
	storage *storages.Storage
}

func NewMetricService(s *storages.Storage) *MetricService {
	return &MetricService{s}
}

func (m *MetricService) InsertMetrics(ctx context.Context, metrics domains.Metrics) error {
	return m.storage.MetricStorage.InsertMetrics(ctx, metrics)
}
func (m *MetricService) GetMetrics(ctx context.Context, filters domains.MetricsFilters) ([]domains.Metrics, error) {
	if filters.Interval == "" {
		return []domains.Metrics{}, ErrNoInterval
	}
	return m.storage.MetricStorage.GetMetrics(ctx, filters)
}
