package logger

import (
	"os"

	logger "github.com/sirupsen/logrus"
)

func init() {
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logger.JSONFormatter{})
	logger.SetLevel(logger.InfoLevel)
}
