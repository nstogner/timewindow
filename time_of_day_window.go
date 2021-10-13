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
	start, end := w.StartTime(now), w.EndTime(now)

	// -----|---------------
	// ---start-----end-----
	if now.Equal(start) {
		return true, 0
	}
	// --------------|------
	// ---start-----end-----
	if now.Equal(end) {
		return false, 24 * time.Hour
	}

	startMinusNow := start.Sub(now)

	// ----------|----------
	// ---start-----end-----
	if now.After(start) && now.Before(end) {
		return true, startMinusNow
	}

	// ------------------|--
	// ---start-----end-----
	if now.After(end) {
		return false, 24*time.Hour + startMinusNow
	}

	// -|-------------------
	// ---start-----end-----
	return false, startMinusNow
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
	return w.Start.Hour <= w.End.Hour
}
