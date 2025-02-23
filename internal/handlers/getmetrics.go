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

func (h *Handler) GetMetrics(c *gin.Context) {
	log := logging.GetLogger().WithField("op", "Handler.GetMetrics")
	log.Info("Request received. GetMetrics")
	var filters domains.MetricsFilters

	filters.Interval = c.Query("interval")

	if deviceID := c.Query("device_id"); deviceID != "" {
		uint64Value, err := strconv.ParseUint(deviceID, 10, 64)
		if err != nil {
			log.Error("Could not parse device id from query string")
			c.Status(http.StatusBadRequest) // fixme
			return
		}
		filters.DeviceID = &uint64Value
	}

	if fromDate := c.Query("from_date"); fromDate != "" {
		uint64Value, err := strconv.ParseUint(fromDate, 10, 64)
		if err != nil {
			log.Error("Could not parse from_date from query string")
			c.Status(http.StatusBadRequest) // fixme
			return
		}
		uintValue := uint(uint64Value)
		filters.FromDate = uintValue
	}

	if toDate := c.Query("to_date"); toDate != "" {
		uint64Value, err := strconv.ParseUint(toDate, 10, 64)
		if err != nil {
			log.Error("Could not parse to_date from query string")
			c.Status(http.StatusBadRequest) // fixme
			return
		}
		uintValue := uint(uint64Value)
		filters.ToDate = uintValue
	}

	if err := validator.New().Struct(&filters); err != nil {
		log.WithError(err).Error("Error validating query")
		c.JSON(http.StatusBadRequest, domains.ErrorBody{Message: "Error validating query"}) // fixme
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
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, domains.SuccessGet{Result: searchRes})
}
