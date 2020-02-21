package config

import (
	"encoding/json"
	"errors"
	"io"
)

var (
	// ErrLoadFailed .
	ErrLoadFailed = errors.New("load failed")
)

// Checker .
type Checker interface {
	Init(conf map[string]interface{}) error
	Check() error
}

// Checks .
type Checks struct {
	Configs []CheckConfig `json:"configs"`
}

// CheckConfig has the structure of the config file
type CheckConfig struct {
	// Name should be unique in the config
	Name   string                 `json:"name"`
	Type   string                 `json:"type"`
	Config map[string]interface{} `json:"config"`
}

// ReadConfig form an io.Reader
func ReadConfig(r io.Reader) (Checks, error) {
	var checks Checks
	err := json.NewDecoder(r).Decode(&checks)
	return checks, err
}
