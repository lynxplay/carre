package cmd

import (
	"github.com/docker/docker/client"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/signal"
	"time"
)

func Execute(ctx *cli.Context) error {
	outputFormat, err := ParseOutputFormat(ctx.String("format"))

	context := DisplayContext{
		Name:         ctx.String("container"),
		OutputFormat: outputFormat,
	}

	dockerClient, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("Failed to create environment dockerClient: %s", err)
		return err
	}

	signalBus := make(chan os.Signal, 1)
	signal.Notify(signalBus, os.Interrupt, os.Kill)

	timer := time.NewTicker(ctx.Duration("interval"))
	for {
		select {
		case _ = <-signalBus:
			return nil
		case _ = <-timer.C:
			DisplayCurrentStats(dockerClient, context)
		}
	}
}
