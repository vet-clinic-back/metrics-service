package entity

type SensorData struct {
	ID                    int     `json:"id" db:"id"`
	Value1LoadCell        float64 `json:"value1_load_cell" db:"value1_load_cell"`
	Value2LoadCell        float64 `json:"value2_load_cell" db:"value2_load_cell"`
	Voltage1LoadCell      float64 `json:"voltage1_load_cell" db:"voltage1_load_cell"`
	Voltage2LoadCellReal  float64 `json:"voltage2_load_cell_real" db:"voltage2_load_cell_real"`
	Voltage2LoadCellImag  float64 `json:"voltage2_load_cell_imag" db:"voltage2_load_cell_imag"`
	Temperature           float64 `json:"temperature" db:"temperature"`
	ValuePulse            float64 `json:"value_pulse" db:"value_pulse"`
	VoltagePulse          float64 `json:"voltage_pulse" db:"voltage_pulse"`
	ValueMuscleActivity   float64 `json:"value_muscle_activity" db:"value_muscle_activity"`
	VoltageMuscleActivity float64 `json:"voltage_muscle_activity" db:"voltage_muscle_activity"`
}
