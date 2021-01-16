package httputil

import (
	"net/http"

	"github.com/pkg/errors"
)

// Create ...
func Create(apiKey, service string, payload interface{}) error {
	_, err := request(Config{
		APIKey:   apiKey,
		Endpoint: service,
		Method:   http.MethodPost,
		Payload:  payload,
	})
	if err != nil {
		return errors.Wrap(err, "MakeRequest")
	}

	return nil
}
