package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func InitLogger() {
	Logger.SetLevel(logrus.InfoLevel)

	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	Logger.SetOutput(os.Stdout)
}
