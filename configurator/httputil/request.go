package httputil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func request(conf Config) ([]byte, error) {
	url := fmt.Sprintf(
		"http://sonarr:8989/api/%s?apikey=%s",
		conf.Endpoint,
		conf.APIKey,
	)

	var body []byte

	if conf.Payload != nil {
		p, err := json.Marshal(conf.Payload)
		if err != nil {
			return nil, errors.Wrap(err, "json.Marshal")
		}
		body = p
	}

	req, err := http.NewRequest(conf.Method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.Wrap(err, "http.NewRequest")
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "client.Do")
	}
	defer res.Body.Close()

	p, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll")
	}

	err = getResponseErrors(p)
	if err != nil {
		return nil, errors.Wrap(err, "error in response")
	}

	return p, nil
}
