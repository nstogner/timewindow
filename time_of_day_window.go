package timewindow

import (
	"fmt"
	"time"
)

func ParseTimeOfDayWindow(start, end string) (*TimeOfDayWindow, error) {
	s, err := ParseTimeOfDay(start)
	if err != nil {
		return nil, fmt.Errorf("start: %w", err)
	}

	e, err := ParseTimeOfDay(end)
	if err != nil {
		return nil, fmt.Errorf("end: %w", err)
	}

	return &TimeOfDayWindow{Start: s, End: e}, nil
}

type TimeOfDayWindow struct {
	Start TimeOfDay
	End   TimeOfDay
}

// WithinWindow returns true if within the window (and the time since the window started
// expressed as a negative number). Otherwise, it returns false and the time
// until the window starts.
func (w *TimeOfDayWindow) WithinWindow(now time.Time) (bool, time.Duration) {
	return WithinWindow(now, w.StartTime(now), w.EndTime(now))
}

func (w *TimeOfDayWindow) StartTime(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), w.Start.Hour, w.Start.Minute, 0, 0, now.Location())
}

func (w *TimeOfDayWindow) EndTime(now time.Time) time.Time {
	end := time.Date(now.Year(), now.Month(), now.Day(), w.End.Hour, w.End.Minute, 0, 0, now.Location())
	if !w.sameDay() {
		end = end.Add(24 * time.Hour)
	}
	return end
}

func (w *TimeOfDayWindow) sameDay() bool {
	return 60*w.Start.Hour+w.Start.Minute <= 60*w.End.Hour+w.End.Minute
}
