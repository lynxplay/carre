package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docker/docker/client"
	"log"
	"strings"
	"time"
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
	Name         string
	OutputFormat OutputMode
}

func DisplayCurrentStats(dockerClient *client.Client, ctx DisplayContext) {
	result, err := dockerClient.ContainerTop(ContextTodo, ctx.Name, DockerTopArgs)
	if err != nil {
		fmt.Printf("Failed to fetch top data for %s (%s)!\n", ctx.Name, err)
		return
	}

	switch ctx.OutputFormat {
	case JSON:
		for i := range result.Processes {
			process := result.Processes[i]
			processValues := make(map[string]string)
			processValues["TIMESTAMP"] = time.Now().String()
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
			fmt.Println(strings.Join(result.Titles, ",") + ",TIMESTAMP") // Avoid slice reallocation for csv print
			PrintedCSVHeader = true
		}

		for i := range result.Processes {
			fmt.Println(strings.Join(result.Processes[i], ",") + "," + time.Now().String())
		}
	}
}

func ParseOutputFormat(str string) (OutputMode, error) {
	switch strings.ToUpper(str) {
	case "CSV":
		return CSV, nil
	case "JSON":
		return JSON, nil // Different case than the default case, JSON was parsed correctly!
	default:
		return JSON, errors.New("failed to parse output format; defaulting to JSON")
	}
}
