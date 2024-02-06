package logging

import (
	"github.com/sirupsen/logrus"
	"os"
	"productApp/app/config"
)

func NewLoggerConsole() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)

	return log
}

// function log to file
func NewLoggerFile(cfg config.IConfig) *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.DebugLevel)
	log.SetOutput(os.Stdout)

	// open file
	file, err := os.OpenFile(cfg.GetConfig().Logging.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to open file, using default stdout")
	}

	return log
}
