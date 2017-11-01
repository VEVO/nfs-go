# nfs-go

## Description

The nfs-go package provides helpers that abide by the N-Factor Service specification.

## Usage Examples

### Logging

Expose the following environment variables:
 * `SERVICE`
 * `TEAM`
 * `LOG_LEVEL`
 * `VERSION`

```go
import "github.com/VEVO/nfs-go/logger"

func main() {
  logger.Log.Info("An informational message")
  logger.Log.Error("An error message")
  logger.Log.Fatal("An fatal message")
}
```

### Statsd

Expose the following environment variables:
 * `SERVICE`
 * `STATSD_HOST`
 * `STATSD_PORT`

```go
import "github.com/VEVO/nfs-go/statsd"

func main() {
  statsd.Statsd.Gauge("some.metric", float64(GetSomeMetricValue()), nil, 1)
}
```
