package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docker/docker/client"
	"log"
	"strings"
)

var (
	ContextTodo      = context.TODO()
	DockerTopArgs    = []string{"auxww"}
	PrintedCSVHeader = false
)

type OutputMode = uint

const (
	CSV  OutputMode = iota
	JSON OutputMode = iota
)

type DisplayContext struct {
	Name       string
	OutputMode OutputMode
}

func DisplayCurrentStats(dockerClient *client.Client, ctx DisplayContext) {
	result, err := dockerClient.ContainerTop(ContextTodo, ctx.Name, DockerTopArgs)
	if err != nil {
		fmt.Printf("Failed to fetch top data for %s (%s)!\n", ctx.Name, err)
		return
	}

	switch ctx.OutputMode {
	case JSON:
		for i := range result.Processes {
			process := result.Processes[i]
			processValues := make(map[string]string)
			for j := range process {
				processValues[result.Titles[j]] = process[j]
			}
			marshal, err := json.Marshal(processValues)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(marshal))
		}
	case CSV:
		if !PrintedCSVHeader {
			fmt.Println(strings.Join(result.Titles, ","))
			PrintedCSVHeader = true
		}

		for i := range result.Processes {
			fmt.Println(strings.Join(result.Processes[i], ","))
		}
	}
}

func ParseOutputMode(str string) (OutputMode, error) {
	switch strings.ToUpper(str) {
	case "JSON":
		return JSON, nil
	default:
		return CSV, errors.New("failed to parse output mode; defaulting to CSV")
	}
}
