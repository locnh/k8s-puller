package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	cron "github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

var images []string

func main() {
	if jsonLog := getConfigLogFormat("JSONLOG"); jsonLog {
		log.SetFormatter(&log.JSONFormatter{})
	}

	interval := getConfigInterval("INTERVAL")
	args, hasValue := os.LookupEnv("IMAGES")
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

func getConfigInterval(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		intVal, err := strconv.Atoi(value)
		if err != nil {
			return value
		}
		return fmt.Sprintf("%vm", intVal)
	}
	return "60m"
}

func getConfigLogFormat(key string) bool {
	if value, ok := os.LookupEnv(key); ok {
		if value == "true" || value == "1" {
			return true
		}
	}
	return false
}

func pull() {
	for _, image := range images {
		log.WithFields(log.Fields{"image": image}).Info("Start pulling image")
		start := time.Now()
		_, err := exec.Command("docker", "pull", fmt.Sprintf("%v", image)).Output()
		if err != nil {
			log.WithError(err)
		}
		t := time.Now()
		duration := t.Sub(start)
		log.WithFields(log.Fields{"image": image, "duration": duration.Round(1 * time.Second)}).Info("Pull finished")
	}
}
