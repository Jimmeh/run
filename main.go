package main

import (
	"context"
	"log"
	"os"

	"github.com/Jimmeh/run/cmd/when"
	"github.com/urfave/cli/v3"
)

func main() {
	root := &cli.Command{
		Name:  "run",
		Usage: "a cli tool for running utilities",
		Commands: []*cli.Command{
			{
				Name:   "when",
				Usage:  "when is the best time to run?",
				Action: when.Exec,
			},
		},
	}
	if err := root.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
