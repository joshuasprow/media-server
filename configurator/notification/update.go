package notification

import (
	"github.com/joshuasprow/media-server/configurator/httputil"
	"github.com/pkg/errors"
)

// Update ...
func Update(id int, apiKey string, args Args) error {
	conf := newConfig(args)
	conf.ID = 1

	err := httputil.Update(apiKey, "notification", &conf)
	if err != nil {
		return errors.Wrap(err, "httputil.MakeRequest")
	}

	return nil
}
