package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/vet-clinic-back/metrics-service/internal/domain"
	"github.com/vet-clinic-back/metrics-service/internal/usecase"
)

type HttpHandler struct {
	uc     usecase.MetricUsecase
	logger *logrus.Logger
}

type SensorData struct {
	Temperature         int     `json:"temperature"`
	ValueMuscleActivity int     `json:"value_muscle_activity"`
	Value1LoadCell      int     `json:"value1_load_cell"`
	Value2LoadCell      int     `json:"value2_load_cell"`
	ValuePulse          float64 `json:"value_pulse"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewHttpHandler(router *gin.Engine, uc usecase.MetricUsecase, logger *logrus.Logger) *HttpHandler {
	handler := &HttpHandler{uc: uc, logger: logger}

	router.POST("/api/data", handler.HandleData)
	router.GET("/api/metrics", handler.HandleGetMetrics)
	router.GET("/sse", handler.HandleSSE)

	return handler
}

// HandleData godoc
// @Summary Save sensor data
// @Description Save medical metrics from sensors
// @Tags data
// @Accept  json
// @Produce  json
// @Param data body SensorData true "Sensor Data"
// @Success 201
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/data [post]
func (h *HttpHandler) HandleData(c *gin.Context) {
	var data SensorData
	if err := c.ShouldBindJSON(&data); err != nil {
		h.logger.WithError(err).Error("Bad request")
		c.JSON(400, ErrorResponse{Message: "Invalid request body"})
		return
	}

	metric := domain.Metric{
		Temperature:     data.Temperature,
		MuscleActivity:  data.ValueMuscleActivity,
		ChestExpansion1: data.Value1LoadCell,
		ChestExpansion2: data.Value2LoadCell,
		Pulse:           data.ValuePulse,
	}

	if err := h.uc.Save(&metric); err != nil {
		h.logger.WithError(err).Error("Failed to save metric")
		c.JSON(500, ErrorResponse{Message: "Internal server error"})
		return
	}

	c.Status(201)
}

// HandleGetMetrics godoc
// @Summary Get metrics
// @Description Get last 10 metrics from database
// @Tags data
// @Produce json
// @Success 200 {array} domain.Metric
// @Failure 500 {object} ErrorResponse
// @Router /api/metrics [get]
func (h *HttpHandler) HandleGetMetrics(c *gin.Context) {
	metrics, err := h.uc.GetLatest()
	if err != nil {
		h.logger.WithError(err).Error("Failed to get metrics")
		c.JSON(500, ErrorResponse{Message: "Failed to retrieve metrics"})
		return
	}

	c.JSON(200, metrics)
}
