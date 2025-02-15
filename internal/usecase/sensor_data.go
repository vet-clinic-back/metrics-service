package usecase

import (
	"context"

	"github.com/vet-clinic-back/metrics-service/internal/entity"
)

type SensorDataRepository interface {
	Save(ctx context.Context, data *entity.SensorData) error
	GetLatest(ctx context.Context) (*entity.SensorData, error)
}

type SensorDataUseCase struct {
	repo SensorDataRepository
}

func NewSensorDataUseCase(repo SensorDataRepository) *SensorDataUseCase {
	return &SensorDataUseCase{repo: repo}
}

func (uc *SensorDataUseCase) ProcessData(ctx context.Context, data *entity.SensorData) error {
	return uc.repo.Save(ctx, data)
}

func (uc *SensorDataUseCase) GetMetrics(ctx context.Context) (*entity.SensorData, error) {
	return uc.repo.GetLatest(ctx)
}
