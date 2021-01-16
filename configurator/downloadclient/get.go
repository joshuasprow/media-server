package downloadclient

import (
	"github.com/joshuasprow/media-server/configurator/httputil"
	"github.com/pkg/errors"
)

// Get ...
func Get(id int, apiKey string) (Config, error) {
	conf := Config{}

	err := httputil.Get(id, apiKey, "downloadclient", &conf)
	if err != nil {
		return Config{}, errors.Wrap(err, "httputil.Get")
	}

	return conf, nil
}

// GetAll ...
func GetAll(apiKey string) ([]Config, error) {
	confs := []Config{}

	err := httputil.GetAll(apiKey, "downloadclient", &confs)
	if err != nil {
		return nil, errors.Wrap(err, "httputil.MakeRequest")
	}

	return confs, nil
}
