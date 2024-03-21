package when

import (
	"fmt"
)

func NewWhenCommand(forecasts forecastRetriever, out output) command {
	return command{
		forecasts: forecasts,
		out:       out,
	}
}

type command struct {
	forecasts forecastRetriever
	out       output
}

func (cmd command) Run() error {
	forecast, err := cmd.forecasts.GetForecast()
	if err != nil {
		return err
	}
	cmd.out.Println(fmt.Sprintf("Location: %s, %s", forecast.Location.Name, forecast.Location.Country))
	cmd.out.Println(fmt.Sprintf("Currently: %dC - %s", int(forecast.Now.Temp), forecast.Now.Condition.Text))
	return nil
}

type output interface {
	Println(line string)
}

type forecastRetriever interface {
	GetForecast() (Forecast, error)
}
