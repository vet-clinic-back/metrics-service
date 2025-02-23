package postgres

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/vet-clinic-back/metrics-service/internal/domains"
	logging "github.com/vet-clinic-back/metrics-service/pkg/logger"
)

/*
CREATE TABLE metrics (
id SERIAL PRIMARY KEY,
device_id INT
pulse INT NOT NULL,
temperature DOUBLE PRECISION NOT NULL,
loadcell_output1 DOUBLE PRECISION NOT NULL,
loadcell_output2 DOUBLE PRECISION NOT NULL,
muscleactivity_output1 DOUBLE PRECISION NOT NULL,
muscleactivity_output2 DOUBLE PRECISION NOT NULL,
created_at BIGINT DEFAULT (EXTRACT(EPOCH FROM NOW()) * 1000)::BIGINT
);
*/

func (p *Postgres) GetMetrics(ctx context.Context, f domains.MetricsFilters) ([]domains.Metrics, error) {
	sq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	log := logging.GetLogger().WithField("op", "Postgres.GetMetrics")

	dateTrunc := fmt.Sprintf("DATE_TRUNC('%s', TO_TIMESTAMP(created_at / 1000))", f.Interval)

	query := sq.Select(
		dateTrunc+" AS time_group",
		"AVG(pulse) AS pulse_avg",
		"AVG(temperature) AS temp_avg",
		"AVG(loadcell_output1) AS lc_o1_avg",
		"AVG(loadcell_output2) AS lc_o2_avg",
		"AVG(muscleactivity_output1) AS m_act_o1_avg",
		"AVG(muscleactivity_output2) AS m_act_o2_avg",
	).
		From("metrics").
		Where(squirrel.Eq{"device_id": f.DeviceID}).
		GroupBy("time_group")

	if f.ToDate != 0 {
		query = query.Where(squirrel.LtOrEq{"created_at": f.ToDate})
	}
	if f.FromDate != 0 {
		query = query.Where(squirrel.GtOrEq{"created_at": f.FromDate})
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	log.WithField("query", sql).Debug("get metrics query")

	rows, err := p.db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var metrics []domains.Metrics
	for rows.Next() {
		var m domains.Metrics
		err := rows.Scan(
			&m.Timestamp,
			&m.Pulse,
			&m.Temperature,
			&m.LoadCell.Output1,
			&m.LoadCell.Output2,
			&m.MuscleActivity.Output1,
			&m.MuscleActivity.Output2,
		)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, m)
	}

	return metrics, nil
}
