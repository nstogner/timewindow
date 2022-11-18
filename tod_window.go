package timewindow

import (
	"fmt"
	"time"
)

func ParseTODWindow(start, end string) (*TODWindow, error) {
	s, err := ParseTOD(start)
	if err != nil {
		return nil, fmt.Errorf("start: %w", err)
	}

	e, err := ParseTOD(end)
	if err != nil {
		return nil, fmt.Errorf("end: %w", err)
	}

	return &TODWindow{Start: s, End: e}, nil
}

type TODWindow struct {
	Start TOD
	End   TOD
}

// WithinWindow returns true if within the window. It also returns the time until
// the next window.
func (w *TODWindow) WithinWindow(now time.Time) WindowResult {
	return WithinWindow(now, w.StartTime(now), w.EndTime(now), w.FollowingStartTime(now))
}

func (w *TODWindow) StartTime(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), w.Start.Hour, w.Start.Minute, 0, 0, now.Location())
}

func (w *TODWindow) FollowingStartTime(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), w.Start.Hour, w.Start.Minute, 0, 0, now.Location()).Add(24 * time.Hour)
}

func (w *TODWindow) EndTime(now time.Time) time.Time {
	end := time.Date(now.Year(), now.Month(), now.Day(), w.End.Hour, w.End.Minute, 0, 0, now.Location())
	if !w.sameDay() {
		end = end.Add(24 * time.Hour)
	}
	return end
}

func (w *TODWindow) sameDay() bool {
	return 60*w.Start.Hour+w.Start.Minute <= 60*w.End.Hour+w.End.Minute
}
