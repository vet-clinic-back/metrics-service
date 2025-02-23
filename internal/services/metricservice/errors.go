package metricservice

import "errors"

var (
	ErrNoInterval = errors.New("filter interval is empty")
	ErrNoDeviceID = errors.New("filter interval is empty")
)
