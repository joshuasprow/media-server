package httputil

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type responseError struct {
	ErrorMessage string `json:"errorMessage"`
	PropertyName string `json:"propertyName"`
}

// API errors are always returned as an array of objects. Each with an
// "errorMessage" key
func getResponseErrors(p []byte) error {
	resErrs := []responseError{}

	err := json.Unmarshal(p, &resErrs)
	if err != nil {
		// any non-array response is a success
		if strings.Contains(err.Error(), "cannot unmarshal object") {
			return nil
		}

		return errors.Wrap(err, "json.Unmarshal")
	}

	if len(resErrs) == 0 {
		return nil
	}

	msgs := []string{}

	for _, resErr := range resErrs {
		em := resErr.ErrorMessage
		pn := resErr.PropertyName

		if em == "" && pn == "" {
			continue
		}

		msgs = append(msgs, fmt.Sprintf("(%s) %s", pn, em))
	}

	if len(msgs) == 0 {
		return nil
	}

	return errors.New(strings.Join(msgs, "; "))
}
