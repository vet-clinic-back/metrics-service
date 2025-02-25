package handlers

import (
	"context"
	"fmt"
	"github.com/vet-clinic-back/metrics-service/internal/adapters"
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

	a.HTTPAdapter.SetHandlers(h)
	a.TCPAdapter.SetHandler(h)

	go a.HTTPAdapter.MustRun()
	go a.TCPAdapter.Listen()

	// graceful shutdown
	g, gCtx := errgroup.WithContext(ctx)
	<-gCtx.Done()
	tCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	g.Go(func() error {
		log.Info("shutting down http...")

		if err := a.HTTPAdapter.Shutdown(tCtx); err != nil {
			return fmt.Errorf("failed to shutdown HTTP adapter: %w", err)
		}

		log.Info("http adapter shutdown completed")
		return nil
	})

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

// var deviceID uint64 = 100500
// m, err := h.service.Metrics.GetMetrics(gCtx, domains.MetricsFilters{
// Interval: "minute",
// DeviceID: &deviceID,
// })
// if err != nil {
// log.WithError(err).Error("Error getting metrics")
// }
//
// for _, m := range m {
// log.WithField("metric", m).Info("got metrics")
// }

//  insertMetrics := domains.Metrics{
//  	DeviceID:    100500,
//  	Pulse:       50,
//  	Temperature: 36.6,
//  	LoadCell: domains.LoadCell{
//  		Output1: 22.3,
//  		Output2: 22.3,
//  	},
//  	MuscleActivity: domains.MuscleActivity{
//  		Output1: 52.2,
//  		Output2: 48,
//  	},
//  }
//  err = h.service.Metrics.InsertMetrics(ctx, insertMetrics)
//  if err != nil {
//  	log.WithError(err).Error("Error metrics insert")
//  }
//  var deviceID uint64 = 100500
// m, err := h.service.Metrics.GetMetrics(gCtx, domains.MetricsFilters{
// Interval: "minute",
// DeviceID: &deviceID,
// })
// if err != nil {
// log.WithError(err).Error("Error getting metrics")
// }
//
// for _, m := range m {
// log.WithField("metric", m).Info("got metrics")
// }

//  insertMetrics := domains.Metrics{
//  	DeviceID:    100500,
//  	Pulse:       50,
//  	Temperature: 36.6,
//  	LoadCell: domains.LoadCell{
//  		Output1: 22.3,
//  		Output2: 22.3,
//  	},
//  	MuscleActivity: domains.MuscleActivity{
//  		Output1: 52.2,
//  		Output2: 48,
//  	},
//  }
//  err = h.service.Metrics.InsertMetrics(ctx, insertMetrics)
//  if err != nil {
//  	log.WithError(err).Error("Error metrics insert")
//  }
//  log.Debug("metrics inserted")
