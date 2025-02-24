package domains

type Metrics struct {
	ID             uint64         `json:"id"`
	DeviceID       uint64         `json:"device_id" validate:"required"`
	Pulse          float64        `json:"pulse"`
	Temperature    float64        `json:"temperature"`
	LoadCell       LoadCell       `json:"load_cell"` // Тензодатчик
	MuscleActivity MuscleActivity `json:"muscle_activity"`
	Timestamp      string         `json:"timestamp"`
}

type LoadCell struct {
	Output1 float64 `json:"output1"`
	Output2 float64 `json:"output2"`
}

type MuscleActivity struct {
	Output1 float64 `json:"output1"`
	Output2 float64 `json:"output2"`
}

type MetricsFilters struct {
	FromDate uint    `json:"from_date"` // Unix Milliseconds
	ToDate   uint    `json:"to_date"`   // Unix Milliseconds
	Interval string  `json:"interval" validate:"required,oneof=minute hour day week"`
	DeviceID *uint64 `json:"device_id" validate:"required"`
}
