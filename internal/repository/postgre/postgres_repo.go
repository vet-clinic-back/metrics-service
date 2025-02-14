package postgres

import (
	"context"

	"github.com/vet-clinic-back/metrics-service/internal/domain"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type PostgresRepo struct {
	conn    *pgx.Conn
	builder squirrel.StatementBuilderType
	logger  *logrus.Logger
}

func New(conn *pgx.Conn, logger *logrus.Logger) *PostgresRepo {
	return &PostgresRepo{
		conn:    conn,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		logger:  logger,
	}
}

func (r *PostgresRepo) Save(m *domain.Metric) error {
	query, args, err := r.builder.
		Insert("metrics").
		Columns(
			"temperature",
			"muscle_activity",
			"chest_expansion1",
			"chest_expansion2",
			"pulse",
		).
		Values(
			m.Temperature,
			m.MuscleActivity,
			m.ChestExpansion1,
			m.ChestExpansion2,
			m.Pulse,
		).
		ToSql()

	if err != nil {
		r.logger.WithError(err).Error("Error building SQL query")
		return err
	}

	_, err = r.conn.Exec(context.Background(), query, args...)
	return err
}

func (r *PostgresRepo) GetLatest() ([]domain.Metric, error) {
	query, args, err := r.builder.
		Select(
			"temperature",
			"muscle_activity",
			"chest_expansion1",
			"chest_expansion2",
			"pulse",
		).
		From("metrics").
		OrderBy("created_at DESC").
		Limit(10).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.conn.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var metrics []domain.Metric
	for rows.Next() {
		var m domain.Metric
		err := rows.Scan(
			&m.Temperature,
			&m.MuscleActivity,
			&m.ChestExpansion1,
			&m.ChestExpansion2,
			&m.Pulse,
		)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, m)
	}

	return metrics, nil
}
