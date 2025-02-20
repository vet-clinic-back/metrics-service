package postgres

import (
	"github.com/vet-clinic-back/metrics-service/internal/domains"
	logging "github.com/vet-clinic-back/metrics-service/pkg/logger"
)

func (p *Postgres) GetNotifications() ([]domains.NotifStatus, error) {
	log := logging.GetLogger().WithField("op", "Postgres.GetNotifStates")
	log.Debug("Getting all notifications MOCK")

	return []domains.NotifStatus{}, nil
}
