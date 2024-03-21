package when

import (
	"fmt"
)

type forecastRetriever interface {
	getForecast() (forecast, error)
}

func NewWhenCommand(forecasts forecastRetriever) command {
	return command{forecasts: forecasts}
}

type command struct {
	forecasts forecastRetriever
}

func (cmd command) Run() error {
	forecast, err := cmd.forecasts.getForecast()
	if err != nil {
		return err
	}
	fmt.Println(forecast)
	return nil
}
