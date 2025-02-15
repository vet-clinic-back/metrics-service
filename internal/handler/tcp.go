package handler

import (
	"context"
	"encoding/json"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/vet-clinic-back/metrics-service/internal/entity"
	"github.com/vet-clinic-back/metrics-service/internal/usecase"
)

type TCPHandler struct {
	uc  *usecase.SensorDataUseCase
	log *logrus.Logger
}

func NewTCPHandler(uc *usecase.SensorDataUseCase, log *logrus.Logger) *TCPHandler {
	return &TCPHandler{uc: uc, log: log}
}

func (h *TCPHandler) Start(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	h.log.Info("TCP server started on ", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			h.log.Error("Accept error:", err)
			continue
		}
		go h.handleConnection(conn)
	}
}

func (h *TCPHandler) handleConnection(conn net.Conn) {
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	var data entity.SensorData

	if err := decoder.Decode(&data); err != nil {
		h.log.Error("Decode error:", err)
		return
	}

	if err := h.uc.ProcessData(context.Background(), &data); err != nil {
		h.log.Error("Process error:", err)
		return
	}

	h.log.Info("Data saved from TCP")
}
