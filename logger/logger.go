package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	// Log is the main structure for logging
	Log         *logrus.Entry
	serviceName string
	teamName    string
	version     string
	logLevel    = logrus.InfoLevel
)

// NewLogger creates a new logger instance
func NewLogger() *logrus.Entry {
	if s := os.Getenv("SERVICE"); s != "" {
		fmt.Printf("test: %s\n", s)
		serviceName = s
	}
	if t := os.Getenv("TEAM"); t != "" {
		teamName = t
	}
	if l := os.Getenv("LOG_LEVEL"); l != "" {
		if level, err := logrus.ParseLevel(l); err != nil {
			fmt.Printf("Failed to parse %s as log level: %s\n", l, err)
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
