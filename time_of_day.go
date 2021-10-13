package timewindow

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type TimeOfDay struct {
	Hour   int
	Minute int
}

func ParseTimeOfDay(s string) (TimeOfDay, error) {
	invalidErr := errors.New("invalid format (expected 12:34): " + s)
	split := strings.Split(s, ":")
	if len(split) != 2 {
		return TimeOfDay{}, invalidErr
	}

	hour, err := strconv.Atoi(split[0])
	if err != nil {
		return TimeOfDay{}, fmt.Errorf("%s: parsing hour: %w", invalidErr, err)
	}
	minute, err := strconv.Atoi(split[1])
	if err != nil {
		return TimeOfDay{}, fmt.Errorf("%s: parsing minute: %w", invalidErr, err)
	}

	if hour < 0 || hour >= 24 {
		return TimeOfDay{}, fmt.Errorf("invalid hour: %v", hour)
	}
	if minute < 0 || minute >= 60 {
		return TimeOfDay{}, fmt.Errorf("invalid minute: %v", minute)
	}

	return TimeOfDay{Hour: hour, Minute: minute}, nil
}
