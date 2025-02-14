package usecase

import "github.com/vet-clinic-back/metrics-service/internal/entity"

type SensorDataRepository interface {
	Save(data *entity.SensorData) error
	GetLatest() (*entity.SensorData, error)
}

type SensorDataUseCase struct {
	repo SensorDataRepository
}

func NewSensorDataUseCase(repo SensorDataRepository) *SensorDataUseCase {
	return &SensorDataUseCase{repo: repo}
}

func (uc *SensorDataUseCase) ProcessData(data *entity.SensorData) error {
	return uc.repo.Save(data)
}

func (uc *SensorDataUseCase) GetLatest() (*entity.SensorData, error) {
	return uc.repo.GetLatest()
}
