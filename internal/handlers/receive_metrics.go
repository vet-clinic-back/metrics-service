package handlers

import (
	"context"
	"encoding/json"
	"github.com/vet-clinic-back/metrics-service/internal/domains"
	logging "github.com/vet-clinic-back/metrics-service/pkg/logger"
	"io"
	"time"
)

const requestTimeout = 9 * time.Second

func (h *Handler) CommonTCPHandler(d *json.Decoder) error {
	return h.ReceiveMetrics(d)
}

func (h *Handler) ReceiveMetrics(d *json.Decoder) error {
	log := logging.GetLogger().WithField("op", "Handler.ReceiveMetrics")

	log.Info("receiving metrics")

	var metrics domains.Metrics
	if err := d.Decode(&metrics); err != nil {
		if err == io.EOF {
			log.WithError(err).Error("Client disconnected")
		} else {
			log.WithError(err).Error("Error decoding JSON")
		}
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	if err := h.service.Metrics.InsertMetrics(ctx, metrics); err != nil {
		log.WithError(err).Error("Error inserting metrics")
		return err
	}
	log.Debug("Inserted metrics successfully")

	return nil
}
