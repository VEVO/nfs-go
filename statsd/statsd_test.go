package statsd_test

import (
	"testing"

	dd_statsd "github.com/DataDog/datadog-go/statsd"
	"github.com/VEVO/nfs-go/statsd"
	"github.com/stretchr/testify/assert"
)

func TestNewStatsd(t *testing.T) {
	s := statsd.NewStatsd()
	assert.IsType(t, &dd_statsd.Client{}, s)
	err := statsd.Statsd.Gauge("metric.test", 1.2, nil, 1)
	assert.NoError(t, err)
}
