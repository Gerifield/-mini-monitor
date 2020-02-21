package debug

import (
	"github.com/gerifield/mini-monitor/src/config"
)

const (
	loadFail  = "loadFail"
	checkFail = "checkFail"
)

type debugChecker struct {
	checkFail bool
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

	if lf, ok := conf[checkFail]; ok {
		if b, ok := lf.(bool); ok {
			d.checkFail = b
		} else {
			return config.ErrLoadFailed
		}
	}
	return nil
}

// Check .
func (d *debugChecker) Check() error {
	if d.checkFail {
		return config.ErrCheckFailed
	}
	return nil
}
