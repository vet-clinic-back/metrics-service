package http

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *HttpHandler) HandleSSE(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			metrics, err := h.uc.GetLatest()
			if err != nil {
				h.logger.WithError(err).Error("SSE error")
				continue
			}

			data, err := json.Marshal(metrics)
			if err != nil {
				h.logger.WithError(err).Error("SSE marshal error")
				continue
			}

			c.SSEvent("message", string(data))
			c.Writer.Flush()

		case <-c.Request.Context().Done():
			return
		}
	}
}
