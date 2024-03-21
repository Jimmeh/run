package when

import "fmt"

func NewConsoleOutput() output {
	return consoleOutput{}
}

type consoleOutput struct{}

func (out consoleOutput) Println(line string) {
	fmt.Println(line)
}
