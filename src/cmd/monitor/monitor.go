package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gerifield/mini-monitor/src/cache"
	"github.com/gerifield/mini-monitor/src/checker/debug"
	"github.com/gerifield/mini-monitor/src/checker/docker"
	"github.com/gerifield/mini-monitor/src/config"
	"github.com/gerifield/mini-monitor/src/loader"
	"github.com/gerifield/mini-monitor/src/server"
)

var availableCheckers = map[string]func() config.Checker{
	"debug":  debug.New,
	"docker": docker.New,
}

func main() {
	listenAddr := flag.String("listen", ":8080", "HTTP endpoint listen")
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

	checkCache := cache.New()
	srv := server.New(*listenAddr, checkCache)
	go srv.Start()

	ticker := time.NewTicker(checks.CheckTime)
	doChecks(loadedModules, checkCache)
	for _ = range ticker.C {
		fmt.Println("Do checks!")
		doChecks(loadedModules, checkCache)
	}
}

func doChecks(loadedModules map[string]config.Checker, cache *cache.Cache) {
	fmt.Println("Module results:")
	for n, m := range loadedModules {
		err := m.Check()

		fmt.Print(n)
		if err != nil {
			fmt.Println("", err)
			cache.Set(n, false)
		} else {
			fmt.Println(" ok")
			cache.Set(n, true)
		}
	}
	fmt.Println()
}
