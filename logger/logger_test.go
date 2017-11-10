package logger_test

import (
	"testing"

	"github.com/VEVO/nfs-go/logger"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	assert.IsType(t, &logrus.Entry{}, logger.Log)
	logger.Log.Info("test log message")
}
