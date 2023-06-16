package main

import (
	"edca3899/string-service/endpoints"
	"edca3899/string-service/services"
	"edca3899/string-service/transports"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc := services.StringService{}
	port := "3333"

	uppercaseHandler := httptransport.NewServer(
		endpoints.MakeUppercaseEndpoint(svc),
		transports.DecodeUppercaseRequest,
		transports.EncodeResponse,
	)

	countHandler := httptransport.NewServer(
		endpoints.MakeCountEndpoint(svc),
		transports.DecodeCountRequest,
		transports.EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)

	log.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
