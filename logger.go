package utils

import (
	"github.com/sirupsen/logrus"
)

type LogWrapper struct {
	*logrus.Logger
}

var registeredLogLevels map[string]*logrus.Level
var registeredLoggers map[string]*LogWrapper

func NewLogger(name string) *LogWrapper {
	logger := GetLogger(name)

	if logger == nil {
		logger = &LogWrapper{
			logrus.New(),
		}

		registeredLoggers[name] = logger

		if registeredLogLevels[name] != nil {
			logger.Logger.SetLevel(*registeredLogLevels[name])
		}
	}

	return logger
}

func GetLogger(name string) *LogWrapper {
	if registeredLoggers == nil {
		registeredLoggers = make(map[string]*LogWrapper)
	}

	return registeredLoggers[name]
}

func SetLoggerLevel(logger string, level string) {
	if registeredLogLevels == nil {
		registeredLogLevels = make(map[string]*logrus.Level)
	}

	levelValue, e := logrus.ParseLevel(level)
	if e != nil {
		panic(e)
	}

	currentLevel := registeredLogLevels[logger]

	if currentLevel != &levelValue {
		registeredLogLevels[logger] = &levelValue

		if currentLevel != nil {
			GetLogger(logger).Logger.SetLevel(levelValue)
		}
	}
}

func (l *LogWrapper) SetLevel(level string) *LogWrapper {
	if level, e := logrus.ParseLevel(level); e == nil {
		l.Logger.SetLevel(level)
	}

	return l
}
