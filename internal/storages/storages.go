package storages

import (
	"github.com/vet-clinic-back/metrics-service/internal/adapters"
)

type Storage struct {
}

func MustNew(a *adapters.Adapters) *Storage {
	return &Storage{}
}
