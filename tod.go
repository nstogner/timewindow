package timewindow

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// TOD is Time Of Day.
type TOD struct {
	Hour   int
	Minute int
}

func ParseTOD(s string) (TOD, error) {
	invalidErr := errors.New("invalid format (expected 12:34): " + s)
	split := strings.Split(s, ":")
	if len(split) != 2 {
		return TOD{}, invalidErr
	}

	hour, err := strconv.Atoi(split[0])
	if err != nil {
		return TOD{}, fmt.Errorf("%s: parsing hour: %w", invalidErr, err)
	}
	minute, err := strconv.Atoi(split[1])
	if err != nil {
		return TOD{}, fmt.Errorf("%s: parsing minute: %w", invalidErr, err)
	}

	if hour < 0 || hour >= 24 {
		return TOD{}, fmt.Errorf("invalid hour: %v", hour)
	}
	if minute < 0 || minute >= 60 {
		return TOD{}, fmt.Errorf("invalid minute: %v", minute)
	}

	return TOD{Hour: hour, Minute: minute}, nil
}
