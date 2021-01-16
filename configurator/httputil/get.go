package httputil

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Get ...
func Get(id int, apiKey, service string, out interface{}) error {
	endpoint := fmt.Sprintf("%s/%d", service, id)

	p, err := request(Config{
		APIKey:   apiKey,
		Endpoint: endpoint,
		Method:   http.MethodGet,
	})
	if err != nil {
		return errors.Wrap(err, "MakeRequest")
	}

	err = json.Unmarshal(p, out)
	if err != nil {
		return errors.Wrap(err, "json.Unmarshal")
	}

	return nil
}

// GetAll ...
func GetAll(apiKey, service string, out interface{}) error {
	p, err := request(Config{
		APIKey:   apiKey,
		Endpoint: service,
		Method:   http.MethodGet,
	})
	if err != nil {
		errors.Wrap(err, "MakeRequest")
	}

	err = json.Unmarshal(p, out)
	if err != nil {
		errors.Wrap(err, "json.Unmarshal")
	}

	return nil
}
