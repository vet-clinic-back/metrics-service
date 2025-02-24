package tcpadapter

import (
	"encoding/json"
	"errors"
	"github.com/vet-clinic-back/metrics-service/internal/config"
	logging "github.com/vet-clinic-back/metrics-service/pkg/logger"
	"net"
)

type TCPAdapter struct {
	port    string
	handler TCPHandler
}

type TCPHandler interface {
	CommonTCPHandler(d *json.Decoder) error
}

func NewTCPAdapter(cfg config.TCPConfig) *TCPAdapter {
	return &TCPAdapter{port: cfg.Port}
}

func (a *TCPAdapter) SetHandler(h TCPHandler) *TCPAdapter {
	a.handler = h
	return a
}

func (a *TCPAdapter) Listen() {
	log := logging.GetLogger().WithField("op", "tcpadapter.Listen")

	if a.port == "" {
		log.Fatal("TCP adapter port is empty")
	}

	if a.handler == nil {
		log.Fatal("TCP adapter handler is nil")
	}

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.WithError(err).Fatal("TCP adapter failed to listen")
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.WithError(err).Error("TCP adapter failed to close listener")
		}
	}(listener)

	log.Info("TCP server started on port 9000")

	//connLabel:
	for {
		conn, err := listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				log.Warn("Listener closed, restarting...")
			} else {
				log.WithError(err).Fatal("Error accepting on TCP adapter")
			}
			continue
		}

		go a.handleConn(conn)
	}
}
