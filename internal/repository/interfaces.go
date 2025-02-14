package repository

import "github.com/vet-clinic-back/metrics-service/internal/domain"

type MetricRepository interface {
	Save(*domain.Metric) error
	GetLatest() ([]domain.Metric, error)
}
