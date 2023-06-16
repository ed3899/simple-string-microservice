package endpoints

import (
	"context"
	"edca3899/string-service/services"

	"github.com/go-kit/kit/endpoint"
)

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

func makeCountEndpoint(svc services.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{v}, nil
	}
}
