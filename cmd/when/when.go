package when

import (
	"fmt"
)

type output interface {
	Println(line string)
}

type ConsoleOutput struct{}

func (out ConsoleOutput) Println(line string) {
	fmt.Println(line)
}

type forecastRetriever interface {
	getForecast() (forecast, error)
}

func NewWhenCommand(forecasts forecastRetriever) command {
	return command{
		forecasts: forecasts,
		out:       ConsoleOutput{},
	}
}

type command struct {
	forecasts forecastRetriever
	out       output
}

func (cmd command) Run() error {
	forecast, err := cmd.forecasts.getForecast()
	if err != nil {
		return err
	}
	cmd.out.Println(fmt.Sprint(forecast))
	return nil
}
