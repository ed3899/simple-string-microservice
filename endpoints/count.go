package endpoints

import (
	"context"
	"edca3899/string-service/services"

	"github.com/go-kit/kit/endpoint"
)
type CountRequest struct {
	S string `json:"s"`
}
type CountResponse struct {
	V int `json:"v"`
}

func MakeCountEndpoint(svc services.StringServiceI) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(CountRequest)
		v := svc.Count(req.S)
		return CountResponse{v}, nil
	}
}
