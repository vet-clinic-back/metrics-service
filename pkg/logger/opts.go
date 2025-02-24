package logging

import "github.com/sirupsen/logrus"

func WithDebug() func(*logrus.Logger) {
	return func(l *logrus.Logger) {
		l.SetLevel(logrus.DebugLevel)
	}
}

func WithInfo() func(*logrus.Logger) {
	return func(l *logrus.Logger) {
		l.SetLevel(logrus.InfoLevel)
	}
}

func WithTextFormatter() func(*logrus.Logger) {
	return func(l *logrus.Logger) {
		l.SetFormatter(&logrus.TextFormatter{
			FieldMap: fieldMap,
		})
	}
}
