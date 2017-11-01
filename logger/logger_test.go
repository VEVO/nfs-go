package logger_test

import (
	"testing"

	"github.com/VEVO/nfs-go/logger"
	"github.com/stretchr/testify/assert"
	"github.com/Sirupsen/logrus"
)

func TestNewLogger(t *testing.T) {
	assert.IsType(t, &logrus.Entry{}, logger.Log)
	logger.Log.Info("test log message")
}
