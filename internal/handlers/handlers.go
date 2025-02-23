package handlers

import (
	"context"
	"fmt"
	"github.com/vet-clinic-back/metrics-service/internal/adapters"
	"github.com/vet-clinic-back/metrics-service/internal/domains"
	"github.com/vet-clinic-back/metrics-service/internal/services"
	logging "github.com/vet-clinic-back/metrics-service/pkg/logger"
	"golang.org/x/sync/errgroup"
	"time"
)

const timeout = time.Second * 10

type Handler struct {
	service *services.Service
}

func New(s *services.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) Run(ctx context.Context, a *adapters.Adapters) {
	log := logging.GetLogger().WithField("op", "app.Run")

	// graceful shutdown
	tCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	var deviceID uint64 = 100500
	m, err := h.service.Metrics.GetMetrics(ctx, domains.MetricsFilters{
		Interval: "minute",
		DeviceID: &deviceID,
	})
	if err != nil {
		log.WithError(err).Error("Error getting metrics")
	}

	for _, m := range m {
		log.WithField("metric", m).Info("got metrics")
	}

	//insertMetrics := domains.Metrics{
	//	DeviceID:    100500,
	//	Pulse:       50,
	//	Temperature: 36.6,
	//	LoadCell: domains.LoadCell{
	//		Output1: 22.3,
	//		Output2: 22.3,
	//	},
	//	MuscleActivity: domains.MuscleActivity{
	//		Output1: 52.2,
	//		Output2: 48,
	//	},
	//}
	//err = h.service.Metrics.InsertMetrics(ctx, insertMetrics)
	//if err != nil {
	//	log.WithError(err).Error("Error metrics insert")
	//}
	//log.Debug("metrics inserted")

	// g.Go(func() error {
	// 	<-gCtx.Done()
	// 	log.Info("shutting down postgres...")
	//
	// 	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// 	defer cancel()
	//
	// 	if err := a.Postgres.Shutdown(); err != nil {
	// 		return fmt.Errorf("failed to shutdown HTTP adapter: %w", err)
	// 	}
	//
	// 	log.Info("http adapter shutdown completed")
	// 	return nil
	// })

	if err := g.Wait(); err != nil {
		log.Errorf("shutdown handlers error: %v", err)
	}

	g.Go(func() error {
		if err := a.Postgres.Shutdown(tCtx); err != nil {
			return fmt.Errorf("failed to shutdown postgres adapter: %w", err)
		}
		log.Info("postgres adapter shutdown completed")
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Errorf("shutdown error: %v", err)
	}
}
