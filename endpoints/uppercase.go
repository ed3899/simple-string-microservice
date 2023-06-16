package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"edca3899/string-service/services"
)

func MakeUppercaseEndpoint(svc services.StringService) endpoint.Endpoint {
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