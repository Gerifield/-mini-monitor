package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gerifield/mini-monitor/src/checker/loader"

	"github.com/gerifield/mini-monitor/src/checker/debug"
	"github.com/gerifield/mini-monitor/src/checker/docker"

	"github.com/gerifield/mini-monitor/src/checker/config"
)

var availableCheckers = map[string]func() config.Checker{
	"debug":  debug.New,
	"docker": docker.New,
}

func main() {
	configFile := flag.String("config", "config.json", "Config file")
	flag.Parse()

	f, err := os.Open(*configFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() { _ = f.Close() }()

	checks, err := config.ReadConfig(f)
	if err != nil {
		log.Fatalln(err)
	}

	loadedModules := loader.LoadModules(availableCheckers, checks)
	fmt.Printf("Loaded %d modules\n", len(loadedModules))
}
