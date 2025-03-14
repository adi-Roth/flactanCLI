package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger instance for FlactanCLI
var Logger = logrus.New()

func InitLogger(level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.InfoLevel
	}

	Logger.SetLevel(lvl)

	// Detect if running in CI/CD or structured logging is needed
	if os.Getenv("FLACTAN_CLI_JSON_LOGS") == "true" {
		Logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		Logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}
}
