package main

import (
	"edca3899/string-service/endpoints"
	"edca3899/string-service/middleware"
	"edca3899/string-service/services"
	"edca3899/string-service/transports"

	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func main() {
	port := "3333"

	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of request received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{})

	// Service logging
	var svc services.StringServiceI
	svc = services.StringService{}
	svc = middleware.AppLoggingMiddleware{
		Logger: logger,
		Next:   svc,
	}
	svc = middleware.PrometheusMiddleware{
		RequestCount:   requestCount,
		RequestLatency: requestLatency,
		CountResult:    countResult,
		Next:           svc,
	}

	// Endpoint logging
	var uppercase endpoint.Endpoint
	uppercase = endpoints.MakeUppercaseEndpoint(svc)
	uppercase = middleware.EnpointLogging(log.With(logger, "method", "uppercase"))(uppercase)

	var count endpoint.Endpoint
	count = endpoints.MakeCountEndpoint(svc)
	count = middleware.EnpointLogging(log.With(logger, "method", "count"))(count)

	uppercaseHandler := httptransport.NewServer(
		uppercase,
		transports.DecodeUppercaseRequest,
		transports.EncodeResponse,
	)

	countHandler := httptransport.NewServer(
		count,
		transports.DecodeCountRequest,
		transports.EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)

	logger.Log("msg", "HTTP", "addr", port)
	logger.Log("err", http.ListenAndServe(":"+port, nil))
}
