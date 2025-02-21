package postgres

import (
	"context"
	"github.com/vet-clinic-back/metrics-service/internal/domains"
)

func (p *Postgres) InsertMetrics(ctx context.Context, metrics domains.Metrics) error {

	q := "INSERT INTO metrics" +
		"(pulse, temperature, loadcell_output1, loadcell_output2, muscleactivity_output1, muscleactivity_output2," +
		"device_id)" +
		"VALUES ($1, $2, $3, $4, $5, $6, $7)"

	_, err := p.db.Exec(
		q,
		metrics.Pulse, metrics.Temperature, metrics.LoadCell.Output1, metrics.LoadCell.Output2,
		metrics.MuscleActivity.Output1, metrics.MuscleActivity.Output2, metrics.DeviceID,
	)
	if err != nil {
		return err
	}

	return nil
}
