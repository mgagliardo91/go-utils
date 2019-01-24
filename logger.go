package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var registeredLoggers map[string]*logrus.Logger

func NewLogger(name string) *logrus.Logger {
	if registeredLoggers[name] != nil {
		return registeredLoggers[name]
	}

	logger := &logrus.Logger{
		Out: os.Stdout,
	}

	registeredLoggers[name] = logger
	return logger
}
