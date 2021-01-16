package cmd

import "encoding/json"

// Arguments ...
type Arguments struct {
	Host   string `arg:"env:HOST"`
	APIKey string `arg:"env:API_KEY"`
	User   string `arg:"env:USER"`
	Pass   string `arg:"env:PASS"`
}

// Config ...
type Config struct {
}

func (c Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(c)
}
