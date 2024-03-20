package when

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type forecast struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Now struct {
		Temp      float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

type weatherApi struct {
	key string
}

func (api *weatherApi) getForecast() (forecast, error) {
	res, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=M28", api.key))
	if err != nil {
		return forecast{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return forecast{}, err
	}

	var f forecast
	err = json.Unmarshal(body, &f)
	if err != nil {
		return forecast{}, err
	}
	return f, nil
}

func NewWeatherApi() (ForecastRetriever, error) {
	key, found := os.LookupEnv("WEATHER_KEY")
	if !found {
		return &weatherApi{}, fmt.Errorf("WEATHER_KEY not found")
	}
	return &weatherApi{
		key: key,
	}, nil
}

type ForecastRetriever interface {
	getForecast() (forecast, error)
}

func NewWhenCommand(forecastRetriever ForecastRetriever) Command {
	return Command{forecastRetriever: forecastRetriever}
}

type Command struct {
	forecastRetriever ForecastRetriever
}

func (cmd *Command) Run() error {
	forecast, err := cmd.forecastRetriever.getForecast()
	if err != nil {
		return err
	}
	fmt.Println(forecast)
	return nil
}
