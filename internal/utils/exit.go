package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

// ExitHandler provides a way to override os.Exit for testing
var ExitHandler = func(code int, err error) {
	if err != nil && code != 0 {
		Logger.WithFields(logrus.Fields{
			"exit_code": code,
			"error":     err.Error(),
		}).Error("Exiting FlactanCLI due to error")
	} else if code != 0 {
		Logger.WithField("exit_code", code).Debug("Exiting FlactanCLI")
	}
	os.Exit(code)
}
