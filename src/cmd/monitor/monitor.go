package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gerifield/mini-monitor/src/loader"

	"github.com/gerifield/mini-monitor/src/checker/debug"
	"github.com/gerifield/mini-monitor/src/checker/docker"

	"github.com/gerifield/mini-monitor/src/config"
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

	// TODO: do actual checking and create a web API to access the results (?)
	ticker := time.NewTicker(checks.CheckTime)

	doChecks(loadedModules)
	for _ = range ticker.C {
		fmt.Println("Do checks!")
		doChecks(loadedModules)
	}
}

func doChecks(loadedModules map[string]config.Checker) {
	fmt.Println("Module results:")
	for n, m := range loadedModules {
		err := m.Check()
		fmt.Printf("%s: %t\n", n, err == nil)
		if err != nil {
			fmt.Println("\t", err)
		}
	}
	fmt.Println()
}
