package main

import (
	"flag"
	"github.com/docker/docker/client"
	"github.com/gonvenience/bunt"
	"log"
	"os"
	"os/signal"
	"time"
)

var (
	container = flag.String("container", "", "-container <container_name>")
	interval  = flag.Duration("interval", time.Millisecond*500, "-duration 10s")
	output    = flag.String("output", "CSV", "-output CSV")
)

func main() {
	flag.Parse()
	if "" == *container {
		flag.Usage()
		_, _ = bunt.Println("Red{-container is required!}")
		return
	}
	mode, err := ParseOutputMode(*output)
	if err != nil {
		_, _ = bunt.Printf("Orange{WARN: %s}\n", err)
	}

	context := DisplayContext{
		Name:       *container,
		OutputMode: mode,
	}

	dockerClient, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("Failed to create environment dockerClient: %s", err)
	}

	signalBus := make(chan os.Signal, 1)
	signal.Notify(signalBus, os.Interrupt, os.Kill)

	timer := time.NewTicker(*interval)
	for {
		select {
		case _ = <-signalBus:
			return
		case _ = <-timer.C:
			DisplayCurrentStats(dockerClient, context)
		}
	}
}
