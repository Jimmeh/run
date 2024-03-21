package when_test

import (
	"encoding/json"
	"testing"

	"github.com/Jimmeh/run/cmd/when"
)

type forecasts struct {
}

func (f forecasts) GetForecast() (when.Forecast, error) {
	forecastJson := `{"location":{"name":"Manchester"}}`
	var result when.Forecast
	err := json.Unmarshal([]byte(forecastJson), &result)
	if err != nil {
		return when.Forecast{}, err
	}
	return result, nil
}

var lines = []string{}

type console struct {
}

func (o console) Println(line string) {
	lines = append(lines, line)
}

func TestWhenOutput(t *testing.T) {
	cmd := when.NewWhenCommand(forecasts{}, console{})
	cmd.Run()

	if len(lines) != 1 {
		t.Fatalf("expected %d lines; got %d", 1, len(lines))
	}
	expected := "location: Manchester"
	if lines[0] != expected {
		t.Fatalf("expected: '%s'; got: '%s'", expected, lines[0])
	}
}
