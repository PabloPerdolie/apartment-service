package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()

	logger.SetLevel(logrus.InfoLevel)

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:    true,
		TimestampFormat:  time.RFC822,
		ForceColors:      true,
		DisableColors:    false,
		QuoteEmptyFields: true,
	})

	logger.SetOutput(os.Stdout)

	return logger
}
