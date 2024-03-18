package when

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/urfave/cli/v3"
)

type Forecast struct {
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

func Exec(ctx context.Context, c *cli.Command) error {
	key, found := os.LookupEnv("WEATHER_KEY")
	if !found {
		return fmt.Errorf("WEATHER_KEY not found")
	}
	res, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=M28", key))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("got %d from api", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var forecast Forecast
	err = json.Unmarshal(body, &forecast)
	if err != nil {
		return err
	}
	fmt.Printf("%v", forecast)
	fmt.Println()
	return nil
}
