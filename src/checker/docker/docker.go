package docker

import "github.com/gerifield/mini-monitor/src/checker/config"

type dockerChecker struct {
}

// New .
func New() config.Checker {
	return &dockerChecker{}
}

// Init .
func (d *dockerChecker) Init(conf map[string]interface{}) error {
	return nil
}

// Check .
func (d *dockerChecker) Check() error {
	return nil
}
