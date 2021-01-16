package sonarr

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/joshuasprow/media-server/configurator/cmd"
	"github.com/joshuasprow/media-server/configurator/downloadclient"

	"github.com/pkg/errors"
)

// Configurator ...
type Configurator struct {
	args cmd.Arguments
}

// NewConfigurator ...
func NewConfigurator(args cmd.Arguments) Configurator {
	return Configurator{
		args: args,
	}
}

// ping ...
func (c Configurator) ping() (available bool, err error) {
	url := fmt.Sprintf(
		"http://%s:8989/api/health?apikey=%s",
		c.args.Host,
		c.args.APIKey,
	)

	res, err := http.Get(url)
	if err != nil {
		return false, errors.Wrap(err, "http.Get")
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, errors.Wrap(err, "ioutil.ReadAll")
	}

	return cmd.IsSuccessfulResponse(b), nil
}

// Wait ...
func (c Configurator) Wait() error {
	for {
		available, err := c.ping()
		if err == io.EOF {
			continue
		}
		if err != nil {
			log.Println(err)
		}

		if available {
			break
		}

		time.Sleep(time.Second)
	}

	log.Println("sonarr started. time to configure!")

	err := downloadclient.Update(
		1,
		c.args.APIKey,
		downloadclient.Args{
			Username: c.args.User,
			Password: c.args.Pass,
		},
	)
	if err != nil {
		return errors.Wrap(err, "downloadclient.Update")
	}

	return nil
}
