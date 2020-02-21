package debug

import (
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
		if b, ok := lf.(bool); ok && b {
			return config.ErrLoadFailed
		}
	}
	return nil
}

// Check .
func (d *debugChecker) Check() error {
	fmt.Println("debug check called")
	return nil
}
