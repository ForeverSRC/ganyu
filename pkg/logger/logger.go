package logger

import "github.com/sirupsen/logrus"

type Logger interface {
	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
}

func DefaultLogger() Logger {
	return NewLogrusLogger(logrus.InfoLevel)
}

func NewLogrusLogger(level logrus.Level) *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(level)

	return logger
}
