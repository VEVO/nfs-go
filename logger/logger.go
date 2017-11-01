package logger

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"os"
)

var (
	Log         *logrus.Entry
	serviceName = ""
	teamName    = ""
	version     = ""
	logLevel    = logrus.InfoLevel
)

func NewLogger() *logrus.Entry {
	if s := os.Getenv("SERVICE"); s != "" {
		fmt.Println(fmt.Sprintf("test: %s", s))
		serviceName = s
	}
	if t := os.Getenv("TEAM"); t != "" {
		teamName = t
	}
	if l := os.Getenv("LOG_LEVEL"); l != "" {
		if level, err := logrus.ParseLevel(l); err != nil {
			fmt.Println(fmt.Sprintf("Failed to parse %s as log level: %s", l, err))
		} else {
			logLevel = level
		}
	}
	if v := os.Getenv("VERSION"); v != "" {
		version = v
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logLevel)
	context := logrus.WithFields(logrus.Fields{
		"service": serviceName,
		"team":    teamName,
		"version": version,
	})

	return context
}

func init() {
	Log = NewLogger()
}
