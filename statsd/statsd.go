package statsd

import (
	"fmt"
	"os"

	dd_statsd "github.com/DataDog/datadog-go/statsd"
	"github.com/VEVO/nfs-go/logger"
)

var (
	Statsd      *dd_statsd.Client
	statsdHost  = "localhost"
	serviceName = ""
	statsdPort  = "8125"
)

func NewStatsd() *dd_statsd.Client {
	if s := os.Getenv("SERVICE"); s != "" {
		serviceName = s
	}
	if h := os.Getenv("STATSD_HOST"); h != "" {
		statsdHost = h
	}
	if p := os.Getenv("STATSD_PORT"); p != "" {
		statsdPort = p
	}

	c, err := dd_statsd.New(fmt.Sprintf("%s:%s", statsdHost, statsdPort))
	if err != nil {
		logger.Log.Fatal(err)
	}

	c.Namespace = fmt.Sprintf("%s.", serviceName)
	c.Tags = append(c.Tags, fmt.Sprintf("pod_name:%s", os.Getenv("POD_NAME")))
	return c
}

func init() {
	Statsd = NewStatsd()
}
