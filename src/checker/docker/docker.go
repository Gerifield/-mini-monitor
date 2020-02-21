package docker

import (
	"regexp"

	"github.com/gerifield/mini-monitor/src/checker/config"
)

type dockerPSOutput struct {
	Command      string `json:"Command"`
	CreatedAt    string `json:"CreatedAt"`
	ID           string `json:"ID"`
	Image        string `json:"Image"`
	Labels       string `json:"Labels"`
	LocalVolumes string `json:"LocalVolumes"`
	Mounts       string `json:"Mounts"`
	Names        string `json:"Names"`
	Networks     string `json:"Networks"`
	Ports        string `json:"Ports"`
	RunningFor   string `json:"RunningFor"`
	Size         string `json:"Size"`
	Status       string `json:"Status"`
}

var (
	confID         = "ID"
	confNameRegex  = "confNameRegex"
	confImageRegex = "confImageRegex"
)

type dockerChecker struct {
	ID         string
	nameRegex  *regexp.Regexp
	imageRegex *regexp.Regexp
}

// New .
func New() config.Checker {
	return &dockerChecker{}
}

// Init .
func (d *dockerChecker) Init(conf map[string]interface{}) error {
	var err error
	// Load the ID matcher
	if lf, ok := conf[confID]; ok {
		if id, ok := lf.(string); ok {
			d.ID = id
		} else {
			return config.ErrLoadFailed
		}
	}

	// Load the name regex
	if lf, ok := conf[confNameRegex]; ok {
		if regex, ok := lf.(string); ok {
			d.nameRegex, err = regexp.Compile(regex)
			if err != nil {
				return config.ErrLoadFailed
			}
		} else {
			return config.ErrLoadFailed
		}
	}

	// Load image regex
	if lf, ok := conf[confImageRegex]; ok {
		if regex, ok := lf.(string); ok {
			d.imageRegex, err = regexp.Compile(regex)
			if err != nil {
				return config.ErrLoadFailed
			}
		} else {
			return config.ErrLoadFailed
		}
	}
	return nil
}

// Check .
func (d *dockerChecker) Check() error {
	return nil
}
