package indexer

import (
	"github.com/joshuasprow/media-server/configurator/httputil"
	"github.com/pkg/errors"
)

// Create ...
func Create(apiKey string, args Args) error {
	conf := newConfig(args)

	err := httputil.Create(apiKey, "indexer", &conf)
	if err != nil {
		return errors.Wrap(err, "httputil.Create")
	}

	return nil
}
