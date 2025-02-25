package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vet-clinic-back/metrics-service/internal/domains"
	"github.com/vet-clinic-back/metrics-service/internal/services/metricservice"
	logging "github.com/vet-clinic-back/metrics-service/pkg/logger"
	"net/http"
	"strconv"
)

// GetMetrics return metrics by using filters
// @Summary Получить метрики
// @Description Получение метрик по фильтрам (временной интервал, ID устройства и т. д.).
// @Description ОБЯЗАТЕЛЬНО юзать device_id и interval
// @Tags metrics
// @Accept  json
// @Produce  json
// @Param interval query string true "Интервал ('minute hour day week')" example("minute")
// @Param device_id query int true "ID устройства" example(100500)
// @Param from_date query int false "Дата начала (timestamp в milliseconds)" example(1708700000000)
// @Param to_date query int false "Дата окончания (timestamp в milliseconds)" example(1708790000000)
// @Success 200 {object} domains.SuccessGet "Успешный ответ"
// @Failure 400 {object} domains.ErrorBody "Ошибка валидации запроса"
// @Failure 500 {object} domains.ErrorBody "Внутренняя ошибка сервера"
// @Router /metrics [get].
func (h *Handler) GetMetrics(c *gin.Context) {
	log := logging.GetLogger().WithField("op", "Handler.GetMetrics")
	log.Info("Request received. GetMetrics")
	var filters domains.MetricsFilters

	filters.Interval = c.Query("interval")

	if deviceID := c.Query("device_id"); deviceID != "" {
		uint64Value, err := strconv.ParseUint(deviceID, 10, 64)
		if err != nil {
			msg := "Could not parse device id from query string"
			log.WithError(err).Error(msg)
			c.JSON(http.StatusBadRequest, domains.ErrorBody{Message: msg})
			return
		}
		filters.DeviceID = &uint64Value
	}

	if fromDate := c.Query("from_date"); fromDate != "" {
		uint64Value, err := strconv.ParseUint(fromDate, 10, 64)
		if err != nil {
			msg := "Could not parse from_date from query string"
			log.WithError(err).Error(msg)
			c.JSON(http.StatusBadRequest, domains.ErrorBody{Message: msg})
			return
		}
		uintValue := uint(uint64Value)
		filters.FromDate = uintValue
	}

	if toDate := c.Query("to_date"); toDate != "" {
		uint64Value, err := strconv.ParseUint(toDate, 10, 64)
		if err != nil {
			msg := "Could not parse to_date from query string"
			log.WithError(err).Error(msg)
			c.JSON(http.StatusBadRequest, domains.ErrorBody{Message: msg})
			return
		}
		uintValue := uint(uint64Value)
		filters.ToDate = uintValue
	}

	if err := validator.New().Struct(&filters); err != nil {
		log.WithError(err).Error("Error validating query")
		c.JSON(http.StatusBadRequest, domains.ErrorBody{Message: "query validation failed"})
		return
	}

	searchRes, err := h.service.Metrics.GetMetrics(c, filters)
	if err != nil {
		log.WithError(err).Error("Error getting metrics")
		if errors.Is(err, metricservice.ErrNoDeviceID) {
			c.JSON(http.StatusBadRequest, domains.ErrorBody{Message: "no device ID"})
			return
		}
		if errors.Is(err, metricservice.ErrNoDeviceID) {
			c.JSON(http.StatusBadRequest, domains.ErrorBody{Message: "no interval"})
			return
		}
		c.JSON(http.StatusInternalServerError, domains.ErrorBody{Message: "internal server error"})
		return
	}

	c.JSON(http.StatusOK, domains.SuccessGet{Result: searchRes})
}
