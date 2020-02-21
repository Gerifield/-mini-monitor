package debug

import (
	"errors"
	"fmt"

	"github.com/gerifield/mini-monitor/src/checker/config"
)

const loadFail = "loadFail"

type debugChecker struct {
}

// New .
func New() config.Checker {
	return &debugChecker{}
}

// Init .
func (d *debugChecker) Init(conf map[string]interface{}) error {
	if lf, ok := conf[loadFail]; ok {
		if lf.(bool) {
			return errors.New("load failed")
		}
	}
	return nil
}

// Check .
func (d *debugChecker) Check() error {
	fmt.Println("debug check called")
	return nil
}
