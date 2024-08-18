package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()

	logger.SetLevel(logrus.InfoLevel)

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logger.SetOutput(os.Stdout)

	return logger
}
