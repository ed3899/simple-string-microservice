package simplestringservice

import (
	"edca3899/string-service/endpoints"
	"edca3899/string-service/services"
	"edca3899/string-service/transports"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc := services.StringService{}

	uppercaseHandler := httptransport.NewServer(
		endpoints.MakeUppercaseEndpoint(svc),
		transports.DecodeUppercaseRequest,
		transports.EncodeResponse,
	)
}


