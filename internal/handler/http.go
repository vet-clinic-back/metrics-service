package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/vet-clinic-back/metrics-service/internal/entity"
	"github.com/vet-clinic-back/metrics-service/internal/usecase"
)

// HTTPHandler представляет обработчики HTTP запросов
type HTTPHandler struct {
	uc  *usecase.SensorDataUseCase
	log *logrus.Logger
}

// NewHTTPHandler создает новый экземпляр HTTPHandler
func NewHTTPHandler(uc *usecase.SensorDataUseCase, log *logrus.Logger) *HTTPHandler {
	return &HTTPHandler{uc: uc, log: log}
}

// GetMetrics обрабатывает запрос на получение метрик
// @Summary Получить последние метрики
// @Description Возвращает последние сохраненные показатели датчиков
// @Tags metrics
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.SensorData "Успешный ответ"
// @Failure 404 {object} map[string]string "Данные не найдены"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /metrics [get]
func (h *HTTPHandler) GetMetrics(c *gin.Context) {
	data, err := h.uc.GetMetrics(context.Background())
	if err != nil {
		h.log.Error("Get metrics error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	if data == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "no data"})
		return
	}

	response := map[string]interface{}{
		"temperature":     data.Temperature,
		"muscle_activity": data.ValueMuscleActivity,
		"chest_expansion": data.Value1LoadCell,
		"pulse":           data.ValuePulse,
	}

	c.JSON(http.StatusOK, response)
}

// SaveMetrics обрабатывает запрос на сохранение метрик
// @Summary Сохранить новые метрики
// @Description Принимает данные с датчиков и сохраняет их в базу данных
// @Tags metrics
// @Accept  json
// @Produce  json
// @Param   input body entity.SensorData true "Данные датчиков"
// @Success 201 "Данные успешно сохранены"
// @Failure 400 {object} map[string]string "Некорректные данные"
// @Failure 500 {object} map[string]string "Ошибка сохранения"
// @Router /metrics [post]
func (h *HTTPHandler) SaveMetrics(c *gin.Context) {
	var data entity.SensorData
	if err := c.ShouldBindJSON(&data); err != nil {
		h.log.Error("Bind error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	if err := h.uc.ProcessData(context.Background(), &data); err != nil {
		h.log.Error("Save error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "save failed"})
		return
	}

	c.Status(http.StatusCreated)
}
