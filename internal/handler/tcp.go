package handler

import (
	"encoding/json"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/vet-clinic-back/metrics-service/internal/entity"
	"github.com/vet-clinic-back/metrics-service/internal/usecase"
)

type TCPHandler struct {
	useCase *usecase.SensorDataUseCase
	log     *logrus.Logger
}

func NewTCPHandler(uc *usecase.SensorDataUseCase, log *logrus.Logger) *TCPHandler {
	return &TCPHandler{useCase: uc, log: log}
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

	var data entity.SensorData
	if err := json.NewDecoder(conn).Decode(&data); err != nil {
		h.log.Error("JSON decode error:", err)
		return
	}

	if err := h.useCase.ProcessData(&data); err != nil {
		h.log.Error("Process data error:", err)
		return
	}

	h.log.Info("Data saved successfully")
}
