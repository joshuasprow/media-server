package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

type response struct {
	ID int `json:"id"`
}

type responseError struct {
	ErrorMessage string `json:"errorMessage"`
}

// Successful responses are JSON objects {}, while errors are returned as an array [].
// Using this strategy to determine the datatype: https://stackoverflow.com/a/55014220
func hasResponseError(j []byte) error {
	is, err := IsJSONObject(j)
	if err != nil {
		return errors.Wrap(err, "IsJSONObject")
	}
	// An object means we have a successful response
	if is {
		return nil
	}

	is, err = IsJSONArray(j)
	if err != nil {
		return errors.Wrap(err, "IsJSONObject")
	}

	errs := []responseError{}

	err = json.Unmarshal(j, &errs)
	if err != nil {
		return errors.Wrap(err, "json.Unmarshal")
	}

	msgs := []string{}

	for _, resErr := range errs {
		msgs = append(msgs, resErr.ErrorMessage)
	}

	return errors.New(strings.Join(msgs, ": "))
}

// ConfigureService ...
func ConfigureService(apiKey string, name string, conf json.Marshaler) (id int, err error) {
	url := fmt.Sprintf(
		"http://localhost:8989/api/%s?apikey=%s",
		name,
		apiKey,
	)

	b, err := conf.MarshalJSON()
	if err != nil {
		return 0, errors.Wrap(err, "conf.Marshal")
	}

	fmt.Println(string(b))

	// return len(b), nil

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return 0, errors.Wrap(err, "http.NewRequest")
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return 0, errors.Wrap(err, "client.Do")
	}
	defer res.Body.Close()

	j, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, errors.Wrap(err, "ioutil.ReadAll")
	}

	err = hasResponseError(j)
	if err != nil {
		return 0, errors.Wrap(err, "hasResponseError")
	}

	body := struct {
		ID int `json:"id"`
	}{}

	err = json.Unmarshal(j, &body)
	if err != nil {
		return 0, errors.Wrapf(err, "json.Unmarshal: %s", string(j))
	}

	if body.ID == 0 {
		return 0, fmt.Errorf("returned ID is 0: %s", string(j))
	}

	return body.ID, nil
}
