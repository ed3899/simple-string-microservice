package main

import (
	"edca3899/string-service/endpoints"
	"edca3899/string-service/middleware"
	"edca3899/string-service/services"
	"edca3899/string-service/transports"

	native_log "log"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
)

func main() {
	svc := services.StringService{}
	port := "3333"

	// Instrumentation
	logger := log.NewLogfmtLogger(os.Stderr)

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

	native_log.Printf("Server listening on port %s\n", port)
	native_log.Fatal(http.ListenAndServe(":"+port, nil))
}
