package domains

type Metrics struct {
	ID          uint64  `json:"id"`
	DeviceID    string  `json:"device_id" validate:"required"`
	Pulse       uint    `json:"pulse"`
	Temperature float64 `json:"temperature"`
	LoadCell    struct {
		Output1 float64 `json:"output1"`
		Output2 float64 `json:"output2"`
	} `json:"LoadCell"` // Тензодатчик
	MuscleActivity struct {
		Output1 float64 `json:"output1"`
		Output2 float64 `json:"output2"`
	} `json:"MuscleActivity"`
	Timestamp string `json:"timestamp"`
}

type MetricsFilters struct {
	FromDate uint64 `json:"from_date"` // Unix Milliseconds
	ToDate   uint64 `json:"to_date"`   // Unix Milliseconds
	Interval string `json:"interval" validate:"required,oneof=minute hour day week"`
	DeviceID uint64 `json:"device_id"`
}
