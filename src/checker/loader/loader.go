package loader

import (
	"log"

	"github.com/gerifield/mini-monitor/src/checker/config"
)

func LoadModules(checkInitFns map[string]func() config.Checker, checks config.Checks) map[string]config.Checker {
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
