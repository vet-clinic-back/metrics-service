package usecase

import (
	"github.com/sirupsen/logrus"
	"github.com/vet-clinic-back/metrics-service/internal/domain"
	"github.com/vet-clinic-back/metrics-service/internal/repository"
)

type MetricUsecase struct {
	repo   repository.MetricRepository
	logger *logrus.Logger
}

func New(repo repository.MetricRepository, logger *logrus.Logger) *MetricUsecase {
	return &MetricUsecase{repo: repo, logger: logger}
}

func (uc *MetricUsecase) Save(metric *domain.Metric) error {
	if err := uc.repo.Save(metric); err != nil {
		uc.logger.WithError(err).Error("Failed to save metric")
		return err
	}
	return nil
}

func (uc *MetricUsecase) GetLatest() ([]domain.Metric, error) {
	return uc.repo.GetLatest()
}
