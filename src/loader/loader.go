package loader

import (
	"log"

	"github.com/gerifield/mini-monitor/src/config"
)

// LoadModules .
func LoadModules(checkInitFns map[string]func() config.Checker, checks config.Conf) map[string]config.Checker {
	loadedModules := make(map[string]config.Checker, 0)
	var module config.Checker
	var err error

	log.Println("Loading modules:")
	for _, c := range checks.Configs {
		log.Printf("%s (%s)", c.Name, c.Type)
		if n, ok := checkInitFns[c.Type]; ok {
			module = n()
			err = module.Init(c.Config)
			if err != nil {
				log.Printf("%s module load failure: %s", c.Name, err)
			} else {
				loadedModules[c.Name] = module
			}
		} else {
			log.Println("Not supported module")
		}
	}
	return loadedModules
}

// ConfigString reads and converts a value from the config map as string
func ConfigString(conf map[string]interface{}, key string) (string, error) {
	if val, ok := conf[key]; ok {
		// If it's found try to cast it or fail
		if valStr, ok := val.(string); ok {
			return valStr, nil
		}
		return "", config.ErrLoadFailed
	}

	// Ignore if it's not found to init the variable with the default/empty value
	return "", nil
}

// ConfigBool reads and converts a value from config map as bool
func ConfigBool(conf map[string]interface{}, key string) (bool, error) {
	if val, ok := conf[key]; ok {
		// If it's found try to cast it or fail
		if valBool, ok := val.(bool); ok {
			return valBool, nil
		}
		return false, config.ErrLoadFailed
	}

	// Ignore if it's not found to init the variable with the default/empty value
	return false, nil
}

// ConfigInt reads and converts a value from config map as integer
func ConfigInt(conf map[string]interface{}, key string) (int, error) {
	if val, ok := conf[key]; ok {
		// If it's found try to cast it or fail
		switch v := val.(type) {
		case float32:
			return int(v), nil
		case int64:
			return int(v), nil
		case int32:
			return int(v), nil
		case float64:
			return int(v), nil
		case int:
			return int(v), nil
		}
		return 0, config.ErrLoadFailed
	}

	// Ignore if it's not found to init the variable with the default/empty value
	return 0, nil
}
