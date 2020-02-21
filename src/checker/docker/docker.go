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
	confDebug      = "debug"
)

type dockerChecker struct {
	ID         string
	nameRegex  *regexp.Regexp
	imageRegex *regexp.Regexp
	debug      bool
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

	// Load bool
	if lf, ok := conf[confDebug]; ok {
		if b, ok := lf.(bool); ok {
			d.debug = b
		} else {
			return config.ErrLoadFailed
		}
	}
	return nil
}

// Check .
func (d *dockerChecker) Check() error {
	cmd := exec.Command("docker", "ps", "--format", "{{json . }}")
	b, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s - %w", string(b), err)
	}

	if d.debug {
		fmt.Println("docker ps output:", string(b))
	}

	var psOut dockerPSOutput
	psOuts := make([]dockerPSOutput, 0)
	decoder := json.NewDecoder(bytes.NewReader(b))
	for {
		err = decoder.Decode(&psOut)
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
			if d.debug {
				fmt.Printf("ID match: %s\n", o.ID)
			}
			return nil
		}

		if d.nameRegex != nil && d.nameRegex.MatchString(o.Names) {
			if d.debug {
				fmt.Printf("Name match: %s (%s)\n", o.Names, o.ID)
			}
			return nil
		}

		if d.imageRegex != nil && d.imageRegex.MatchString(o.Image) {
			if d.debug {
				fmt.Printf("Image match: %s (%s)\n", o.Names, o.ID)
			}
			return nil
		}
	}
	return config.ErrCheckFailed // No match found
}
