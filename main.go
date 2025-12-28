package main

import (
	"github.com/charmbracelet/log"
	"main/handlers"
	"main/helpers"
	"sync"
)

func main() {
	targets, err := helpers.ReadTargets("./config/targets.yaml")
	if err != nil {
		log.Error("Error during read yaml file")
		return
	}
	client, err := helpers.GetTorClient()
	if err != nil {
		log.Error("Tor Connection fail", err)
		return
	}
	var wg sync.WaitGroup
	for _, url := range targets {
		wg.Add(1)
		go func(targetURL string) {
			defer wg.Done()
			err := handlers.Screaper(client, url)
			if err != nil {
				log.Error("Unsuccessful request", err)

			}

		}(url)

	}
	wg.Wait()
	handlers.SaveReport()
	log.Info("Completed scanning. Report saved to scan_report.log")
}
