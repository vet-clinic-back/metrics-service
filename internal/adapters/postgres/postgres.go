package postgres

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // postgres
	"github.com/vet-clinic-back/metrics-service/internal/config"
	logging "github.com/vet-clinic-back/metrics-service/pkg/logger"
	"log"
	"net"
)

type Postgres struct {
	db *sql.DB
}

func MustNew(cfg config.Postgres) *Postgres {
	db := mustInitConn(cfg)
	logging.GetLogger().WithField("op", "postgres.MustNew").Info("Connected to postgres successfully")

	return &Postgres{db}
}

func mustInitConn(cfg config.Postgres) *sql.DB {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.User, cfg.Password, net.JoinHostPort(cfg.Host, cfg.Port), cfg.Database,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
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
