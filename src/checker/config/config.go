package config

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

var (
	// ErrLoadFailed .
	ErrLoadFailed = errors.New("load failed")

	// ErrCheckFailed .
	ErrCheckFailed = errors.New("check failed")
)

// Checker .
type Checker interface {
	Init(conf map[string]interface{}) error
	Check() error
}

// Conf .
type Conf struct {
	CheckTime time.Duration
	Configs   []CheckConfig
}

type confFile struct {
	// Duration miss the unmarshaler so we need a bit workaround
	CheckTime string        `json:"checkTime"`
	Configs   []CheckConfig `json:"configs"`
}

// CheckConfig has the structure of the config file
type CheckConfig struct {
	// Name should be unique in the config
	Name   string                 `json:"name"`
	Type   string                 `json:"type"`
	Config map[string]interface{} `json:"config"`
}

// ReadConfig form an io.Reader
func ReadConfig(r io.Reader) (Conf, error) {
	var checks Conf

	var confFile confFile
	err := json.NewDecoder(r).Decode(&confFile)
	if err != nil {
		return checks, err
	}

	// Parse the check duration
	d, err := time.ParseDuration(confFile.CheckTime)
	if err != nil {
		return checks, err
	}

	checks.CheckTime = d
	// Copy the rest
	checks.Configs = confFile.Configs
	return checks, err
}
