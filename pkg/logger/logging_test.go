package logging_test

import (
	logging "github.com/vet-clinic-back/metrics-service/pkg/logger"
	"testing"
)

func TestLoggerOpts(t *testing.T) {
	logging.GetLogger().WithField("test", t.Name()).Debug("test debug")
	logging.GetLogger().WithField("test", t.Name()).Info("test info")

	logging.InitDefaultLogger(logging.WithDebug())

	logging.GetLogger().WithField("test", t.Name()).Debug("test debug 2")

	logging.InitDefaultLogger(logging.WithDebug(), logging.WithTextFormatter())
	logging.GetLogger().WithField("test", t.Name()).Debug("test debug 3")

	logging.InitDefaultLogger(logging.WithDebug())
	logging.GetLogger().WithField("test", t.Name()).Debug("test debug 3")
}
