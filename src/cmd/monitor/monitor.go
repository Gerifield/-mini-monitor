package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/Gerifield/mini-monitor/src/cache"
	"github.com/Gerifield/mini-monitor/src/checker/debug"
	"github.com/Gerifield/mini-monitor/src/checker/docker"
	"github.com/Gerifield/mini-monitor/src/checker/http"
	"github.com/Gerifield/mini-monitor/src/config"
	"github.com/Gerifield/mini-monitor/src/loader"
	"github.com/Gerifield/mini-monitor/src/server"
)

var availableCheckers = map[string]func() config.Checker{
	"debug":  debug.New,
	"docker": docker.New,
	"http":   http.New,
}

func main() {
	listenAddr := flag.String("listen", "127.0.0.1:8080", "HTTP endpoint listen")
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
	log.Printf("Loaded %d modules\n", len(loadedModules))

	checkCache := cache.New()
	srv := server.New(*listenAddr, checkCache)
	go func() { _ = srv.Start() }()

	ticker := time.NewTicker(checks.CheckTime)
	doChecks(loadedModules, checkCache)
	for _ = range ticker.C {
		doChecks(loadedModules, checkCache)
	}
}

func doChecks(loadedModules map[string]config.Checker, cache *cache.Cache) {
	//fmt.Println("Module results:")
	for n, m := range loadedModules {
		err := m.Check()

		//fmt.Print(n)
		if err != nil {
			//fmt.Println("", err)
			cache.Set(n, false)
		} else {
			//fmt.Println(" ok")
			cache.Set(n, true)
		}
	}
}
