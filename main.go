package main

import (
	"github.com/lynxplay/carre/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

var version string

func main() {
	app := &cli.App{
		Name:    "carre",
		Usage:   "Docker process data point collection",
		Version: getVersion(),
		Action:  cmd.Execute,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "container",
				Aliases:  []string{"C"},
				Usage:    "unique container name",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"F"},
				Value:   "JSON",
				Usage:   "output format of caree",
			},
			&cli.DurationFlag{
				Name:    "interval",
				Aliases: []string{"I"},
				Value:   500 * time.Millisecond,
				Usage:   "interval between data point collection",
			},
		},
	}
	app.EnableBashCompletion = true

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getVersion() string {
	if version == "" {
		return "develop"
	}
	return version
}
