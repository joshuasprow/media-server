package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"log"

	"github.com/pkg/errors"
)

const (
	array  = "array"
	object = "object"
)

// IsJSONArray ...
func IsJSONArray(b []byte) (bool, error) {
	t, err := findJSONType(bytes.NewBuffer(b))
	if err != nil {
		return false, errors.Wrap(err, "findJSONType")
	}
	return t == array, nil
}

// IsJSONObject ...
func IsJSONObject(b []byte) (bool, error) {
	t, err := findJSONType(bytes.NewBuffer(b))
	if err != nil {
		return false, errors.Wrap(err, "findJSONType")
	}
	return t == array, nil
}

// IsSuccessfulResponse ...
func IsSuccessfulResponse(b []byte) bool {
	is, err := IsJSONArray(b)
	if err != nil {
		log.Println(errors.Wrap(err, "IsJSONArray"))
		return false
	}

	is, err = IsJSONObject(b)
	if err != nil {
		log.Println(errors.Wrap(err, "IsJSONObject"))
		return false
	}

	return is
}

// Thanks! https://stackoverflow.com/a/55017470
func findJSONType(in io.Reader) (string, error) {
	dec := json.NewDecoder(in)
	// Get just the first valid JSON token from input
	t, err := dec.Token()
	if err != nil {
		return "", err
	}
	if d, ok := t.(json.Delim); ok {
		// The first token is a delimiter, so this is an array or an object
		switch d {
		case '[':
			return array, nil
		case '{':
			return object, nil
		default: // ] or }
			return "", errors.New("Unexpected delimiter")
		}
	}
	return "", errors.New("Input does not represent a JSON object or array")
}
