package services

import (
	"github.com/vet-clinic-back/metrics-service/internal/storages"
)

type Service struct {
}

func MustNew(s *storages.Storage) *Service {
	return &Service{}
}
