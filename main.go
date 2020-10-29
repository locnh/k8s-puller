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
var isValid bool

func main() {
	if jsonLog := getConfigLogFormat("JSONLOG"); jsonLog {
		log.SetFormatter(&log.JSONFormatter{})
	}

	interval := getConfigInterval("INTERVAL")

	images, isValid = getImages("IMAGES")
	if !isValid {
		log.Error("Missing Env IMAGES")
		os.Exit(1)
	}

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
	if result, ok := os.LookupEnv(key); ok {
		intVal, err := strconv.Atoi(result)
		if err != nil {
			return result
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

func getImages(key string) (result []string, hasValue bool) {
	if value, ok := os.LookupEnv(key); ok {
		result = strings.Split(value, ",")
		return result, true
	}
	return result, false
}

func pull() {
	for _, image := range images {
		log.WithFields(log.Fields{"image": image}).Info("Start pulling image")
		start := time.Now()
		_, err := exec.Command("docker", "pull", fmt.Sprintf("%v", image)).Output()
		t := time.Now()
		duration := t.Sub(start)
		if err != nil {
			log.WithError(err)
		} else {
			log.WithFields(log.Fields{"image": image, "duration": duration.Round(1 * time.Second)}).Info("Pull finished")
		}
	}
}
