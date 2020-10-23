package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	cron "github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

var images []string

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	args, hasValue := os.LookupEnv("IMAGES")
	interval := getEnv("INTERVAL", "2m")

	if !hasValue {
		log.Error("Missing IMAGES")
		os.Exit(1)
	}

	images = strings.Split(args, ",")

	for _, image := range images {
		log.WithFields(log.Fields{"image": image}).Info("Added image to watched list")
	}
	log.WithFields(log.Fields{"interval": interval}).Info("Set scheduled interval")

	wg := &sync.WaitGroup{}
	wg.Add(1)
	c := cron.New()
	_, err := c.AddFunc(fmt.Sprintf("@every %v", interval), pull)
	if err != nil {
		log.WithError(err).Fatal("Failed to add cron")
		os.Exit(1)
	}
	c.Start()
	wg.Wait()
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func pull() {
	for _, image := range images {
		log.WithFields(log.Fields{"image": image}).Info("Start pulling image")
		_, err := exec.Command("docker", "pull", fmt.Sprintf("%v", image)).Output()
		if err != nil {
			log.WithError(err)
		}
		log.WithFields(log.Fields{"image": image}).Info("Pull finished")
	}
}
