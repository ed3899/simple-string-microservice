package endpoints

import (
	"context"
	"edca3899/string-service/services"

	"github.com/go-kit/kit/endpoint"
)

func MakeCountEndpoint(svc services.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(CountRequest)
		v := svc.Count(req.S)
		return CountResponse{v}, nil
	}
}
