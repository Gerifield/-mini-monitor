package docker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
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

const (
	confID         = "id"
	confNameRegex  = "nameRegex"
	confImageRegex = "imageRegex"
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

	// TODO: Maybe add a debug flag to help the checking of the matching containers
	return nil
}

// Check .
func (d *dockerChecker) Check() error {
	cmd := exec.Command("docker", "ps", "--format", "{{json . }}")
	b, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s - %w", string(b), err)
	}

	var psOut dockerPSOutput
	psOuts := make([]dockerPSOutput, 0)
	decoder := json.NewDecoder(bytes.NewReader(b))
	for {
		err = decoder.Decode(&psOut) // TODO: Call this multiple times until we have data!!! (\n separated jsons)
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		psOuts = append(psOuts, psOut)
	}

	// Do the checks on the output
	return d.doChecks(psOuts)
}

func (d *dockerChecker) doChecks(psOuts []dockerPSOutput) error {
	for _, o := range psOuts {
		if o.ID == d.ID {
			return nil
		}

		if d.nameRegex != nil && d.nameRegex.MatchString(o.Names) {
			return nil
		}

		if d.imageRegex != nil && d.imageRegex.MatchString(o.Image) {
			return nil
		}
	}
	return config.ErrCheckFailed // No match found
}
