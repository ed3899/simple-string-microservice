package transports

import (
	"context"
	"edca3899/string-service/endpoints"
	"encoding/json"
	"net/http"
)

func DecodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoints.UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}


