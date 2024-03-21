package when_test

import (
	"encoding/json"
	"testing"

	"github.com/Jimmeh/run/cmd/when"
)

type forecasts struct {
}

func (f forecasts) GetForecast() (when.Forecast, error) {
	forecastJson := `{"location":{"name":"Manchester", "country":"UK"}, "current":{"temp_c":11.5, "condition":{"text":"Windy"}}}`
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
	err := cmd.Run()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{
		"Location: Manchester, UK",
		"Currently: 11C - Windy",
	}

	for i, line := range expected {
		if lines[i] != line {
			t.Fatalf("expected: '%s'; got: '%s'", line, lines[i])
		}
	}
}
