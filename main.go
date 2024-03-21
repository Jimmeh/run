package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Jimmeh/run/cmd/when"
)

func showUsageMessage() {
	fmt.Println("Usage: run <cmd>")
	fmt.Println("Commands:")
	fmt.Println("    when    when is the best time to run?")
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		showUsageMessage()
		return
	}

	cmd := args[0]
	if strings.ToLower(cmd) == "when" {
		runWhenCmd()
	} else {
		showUsageMessage()
	}
}

func runWhenCmd() {
	api, err := when.NewWeatherApi()
	if err != nil {
		fmt.Println("ERR: ", err)
		return
	}
	cmd := when.NewWhenCommand(api, when.NewConsoleOutput())
	cmd.Run()
}
