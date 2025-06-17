package log

import "github.com/sirupsen/logrus"

var Logger *logrus.Logger

func SetupLoger() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	log.Info("Log initiated using logrus")
	Logger = log
}
