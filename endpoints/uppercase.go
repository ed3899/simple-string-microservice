package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"edca3899/string-service/services"
)

type UppercaseRequest struct {
	S string `json:"s"`
}

type UppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func MakeUppercaseEndpoint(svc services.StringServiceI) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return UppercaseResponse{
				v,
				err.Error(),
			}, nil
		}
		return UppercaseResponse{
				v,
				"",
			},
			nil
	}
}