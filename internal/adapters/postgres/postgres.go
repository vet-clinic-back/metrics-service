package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq" // postgres
	"github.com/vet-clinic-back/metrics-service/internal/config"
	logging "github.com/vet-clinic-back/metrics-service/pkg/logger"
	"net"
	"time"
)

const retryDelay = time.Second * 1

type Postgres struct {
	db      *sql.DB
	connStr string
}

func MustNew(cfg config.Postgres) *Postgres {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.User, cfg.Password, net.JoinHostPort(cfg.Host, cfg.Port), cfg.Database,
	)

	pg := &Postgres{connStr: connStr}

	if err := pg.EnsureConnection(); err != nil {
		log := logging.GetLogger().WithField("op", "postgres.MustNew")
		log.WithError(err).Fatal("failed to connect to postgres")
	}

	return pg
}

func (p *Postgres) EnsureConnection() error {
	log := logging.GetLogger().WithField("op", "Postgres.Reconnect")

	shouldConnect := true
	maxTries := 3

	if p.db != nil {
		if err := p.db.Ping(); err != nil {
			log.WithError(err).Error("Failed to ping postgres")
		} else {
			shouldConnect = false
		}
	}

	if !shouldConnect {
		return nil
	}
	for i := 0; i < maxTries; i++ {
		db, err := sql.Open("postgres", p.connStr)
		if err != nil {
			log.WithError(err).Error("Failed to connect to postgres")
			time.Sleep(retryDelay)
			continue
		}

		if err = db.Ping(); err != nil {
			log.WithError(err).Error("Failed to ping postgres")
			time.Sleep(retryDelay)
		} else {
			p.db = db
			log.Infof("Connected to postgres by %v try", i+1)
			return nil
		}
	}
	return errors.New("failed to connect to postgres")
}

func (p *Postgres) Shutdown(ctx context.Context) error {
	done := make(chan error, 1)
	go func() {
		done <- p.db.Close()
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-done:
		return err
	}
}
