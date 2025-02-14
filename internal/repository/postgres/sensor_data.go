package postgres

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/vet-clinic-back/metrics-service/internal/entity"
)

type SensorDataRepo struct {
	db      *sqlx.DB
	builder squirrel.StatementBuilderType
}

func NewSensorDataRepo(db *sqlx.DB) *SensorDataRepo {
	return &SensorDataRepo{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *SensorDataRepo) Save(data *entity.SensorData) error {
	query, args, err := r.builder.Insert("sensor_data").
		Columns(
			"value1_load_cell",
			"value2_load_cell",
			"voltage1_load_cell",
			"voltage2_load_cell_real",
			"voltage2_load_cell_imag",
			"temperature",
			"value_pulse",
			"voltage_pulse",
			"value_muscle_activity",
			"voltage_muscle_activity",
		).
		Values(
			data.Value1LoadCell,
			data.Value2LoadCell,
			data.Voltage1LoadCell,
			data.Voltage2LoadCellReal,
			data.Voltage2LoadCellImag,
			data.Temperature,
			data.ValuePulse,
			data.VoltagePulse,
			data.ValueMuscleActivity,
			data.VoltageMuscleActivity,
		).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *SensorDataRepo) GetLatest() (*entity.SensorData, error) {
	query, args, err := r.builder.Select("*").
		From("sensor_data").
		OrderBy("created_at DESC").
		Limit(1).
		ToSql()

	if err != nil {
		return nil, err
	}

	var data entity.SensorData
	err = r.db.Get(&data, query, args...)
	return &data, err
}
