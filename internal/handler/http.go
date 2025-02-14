package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/vet-clinic-back/metrics-service/internal/usecase"
)

type MetricsResponse struct {
	Temperature    float64 `json:"temperature"`
	MuscleActivity float64 `json:"muscle_activity"`
	ChestExpansion float64 `json:"chest_expansion"`
	Pulse          float64 `json:"pulse"`
}

type HTTPHandler struct {
	useCase *usecase.SensorDataUseCase
	log     *logrus.Logger
}

func NewHTTPHandler(uc *usecase.SensorDataUseCase, log *logrus.Logger) *HTTPHandler {
	return &HTTPHandler{useCase: uc, log: log}
}

// @Summary Get latest metrics
// @Description Get latest health metrics from sensors
// @Tags metrics
// @Accept  json
// @Produce  json
// @Success 200 {object} MetricsResponse
// @Failure 500 {object} map[string]string
// @Router /metrics [get]
func (h *HTTPHandler) GetMetrics(c *gin.Context) {
	data, err := h.useCase.GetLatest()
	if err != nil {
		h.log.Error("Get metrics error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	response := MetricsResponse{
		Temperature:    data.Temperature,
		MuscleActivity: data.ValueMuscleActivity,
		ChestExpansion: data.Value1LoadCell,
		Pulse:          data.ValuePulse,
	}

	c.JSON(http.StatusOK, response)
}
