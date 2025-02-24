package tcpadapter

import (
	"encoding/json"
	logging "github.com/vet-clinic-back/metrics-service/pkg/logger"
	"net"
)

func (a *TCPAdapter) handleConn(conn net.Conn) {
	log := logging.GetLogger().WithField("op", "TCPAdapter.handleConn")
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.WithError(err).Error("Error closing TCP connection")
		}
	}(conn)

	decoder := json.NewDecoder(conn)
	for {
		if err := a.handler.CommonTCPHandler(decoder); err != nil {
			break
		}
	}
}
