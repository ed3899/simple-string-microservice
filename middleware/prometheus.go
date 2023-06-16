package middleware

import (
	"edca3899/string-service/services"

	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

type PrometheusMiddleware struct {
	request metrics.Counter
	requestLatency metrics.Histogram
	countResult metrics.Histogram
	next services.StringServiceI
}