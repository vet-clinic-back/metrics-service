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

	_, err := h.service.Metrics.GetMetrics(ctx, domains.MetricsFilters{Interval: "minute"})
	if err != nil {
		log.WithError(err).Error("Error getting metrics")
	}

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
