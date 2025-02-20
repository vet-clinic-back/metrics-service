package logging

import "testing"

func TestLoggerOpts(t *testing.T) {
	GetLogger().WithField("test", t.Name()).Debug("test debug")
	GetLogger().WithField("test", t.Name()).Info("test info")

	InitDefaultLogger(WithDebug())

	GetLogger().WithField("test", t.Name()).Debug("test debug 2")

	InitDefaultLogger(WithDebug(), WithTextFormatter())
	GetLogger().WithField("test", t.Name()).Debug("test debug 3")

	InitDefaultLogger(WithDebug())
	GetLogger().WithField("test", t.Name()).Debug("test debug 3")
}
