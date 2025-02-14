package clickhouse

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"
	"github.com/vet-clinic-back/metrics-service/internal/domain"
)

type ClickHouseRepo struct {
	conn    clickhouse.Conn
	builder squirrel.StatementBuilderType
	logger  *logrus.Logger
}

func New(conn clickhouse.Conn, logger *logrus.Logger) *ClickHouseRepo {
	return &ClickHouseRepo{
		conn:    conn,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question),
		logger:  logger,
	}
}

func (r *ClickHouseRepo) Save(m *domain.Metric) error {
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

	return r.conn.AsyncInsert(context.Background(), query, false, args...)
}

func (r *ClickHouseRepo) GetLatest() ([]domain.Metric, error) {
	query, args, err := r.builder.
		Select(
			"temperature",
			"muscle_activity",
			"chest_expansion1",
			"chest_expansion2",
			"pulse",
		).
		From("metrics").
		OrderBy("timestamp DESC").
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
		if err := rows.Scan(
			&m.Temperature,
			&m.MuscleActivity,
			&m.ChestExpansion1,
			&m.ChestExpansion2,
			&m.Pulse,
		); err != nil {
			return nil, err
		}
		metrics = append(metrics, m)
	}

	return metrics, nil
}
