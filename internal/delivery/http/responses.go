package http

import "github.com/vet-clinic-back/metrics-service/internal/domain"

// swagger:response ErrorResponse
type ErrorResponseWrapper struct {
	// in:body
	Body ErrorResponse
}

// swagger:response MetricsResponse
type MetricsResponseWrapper struct {
	// in:body
	Body []domain.Metric
}
