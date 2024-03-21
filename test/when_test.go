package when_test

import (
	"testing"

	"github.com/Jimmeh/run/cmd/when"
)

type forecasts struct {
}

func (f forecasts) GetForecast() (when.Forecast, error) {
	return when.Forecast{}, nil
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
}
