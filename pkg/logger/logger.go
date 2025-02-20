package logging

import (
	"github.com/sirupsen/logrus"
)

var entry *logrus.Entry

var fieldMap = logrus.FieldMap{
	logrus.FieldKeyTime:  ".timestamp",
	logrus.FieldKeyLevel: "@level",
	logrus.FieldKeyMsg:   "@message",
	logrus.FieldKeyFunc:  "z_caller",
}

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{entry}
}

func InitDefaultLogger(opts ...func(*logrus.Logger)) {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.Formatter = &logrus.JSONFormatter{
		FieldMap: fieldMap,
	}

	for _, opt := range opts {
		opt(log)
	}

	entry = logrus.NewEntry(log)
}

type Flags struct {
	PrettyLog *bool
	Release   *bool
}

func UpdateByFlags(flags Flags) {
	if flags.PrettyLog != nil && *flags.PrettyLog {
		UpdateOpts(WithTextFormatter())
	}
	if flags.Release != nil && *flags.Release {
		UpdateOpts(WithDebug())
	}
}

func UpdateOpts(opts ...func(*logrus.Logger)) {
	log := GetLogger().Logger
	for _, opt := range opts {
		opt(log)
	}
	entry = logrus.NewEntry(log)
}
