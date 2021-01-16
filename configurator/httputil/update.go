package httputil

import (
	"net/http"

	"github.com/pkg/errors"
)

// Update ...
func Update(apiKey, service string, payload interface{}) error {
	_, err := request(Config{
		APIKey:   apiKey,
		Endpoint: service,
		Method:   http.MethodPut,
		Payload:  payload,
	})
	if err != nil {
		return errors.Wrap(err, "MakeRequest")
	}

	return nil
}
