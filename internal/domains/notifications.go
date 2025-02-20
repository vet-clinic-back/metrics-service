package domains

// NotifStatus is an object contains in postgres db and references to Notification.
type Metrics struct {
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
}
