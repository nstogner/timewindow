package timewindow

import (
	"fmt"
	"time"
)

func ParseTODWeekWindow(start, end string, weekdays []string) (*TODWeekWindow, error) {
	s, err := ParseTOD(start)
	if err != nil {
		return nil, fmt.Errorf("start: %w", err)
	}

	e, err := ParseTOD(end)
	if err != nil {
		return nil, fmt.Errorf("end: %w", err)
	}

	w, err := ParseWeekdays(weekdays)
	if err != nil {
		return nil, fmt.Errorf("weekdays: %w", err)
	}

	return &TODWeekWindow{Start: s, End: e, Weekdays: w}, nil
}

type TODWeekWindow struct {
	Weekdays
	Start TOD
	End   TOD
}

// WithinWindow returns true if within the window. It also returns the time until
// the next window.
func (w *TODWeekWindow) WithinWindow(now time.Time) WindowResult {
	return WithinWindow(now, w.StartTime(now), w.EndTime(now), w.FollowingStartTime(now))
}

func (w *TODWeekWindow) StartTime(now time.Time) time.Time {
	return w.accountForWeekday(time.Date(now.Year(), now.Month(), now.Day(), w.Start.Hour, w.Start.Minute, 0, 0, now.Location()))
}

func (w *TODWeekWindow) FollowingStartTime(now time.Time) time.Time {
	return w.accountForWeekday(w.StartTime(now).Add(24 * time.Hour))
}

func (w *TODWeekWindow) EndTime(now time.Time) time.Time {
	end := w.accountForWeekday(time.Date(now.Year(), now.Month(), now.Day(), w.End.Hour, w.End.Minute, 0, 0, now.Location()))

	if !w.sameDay() {
		end = end.Add(24 * time.Hour)
	}

	return end
}

func (w *TODWeekWindow) accountForWeekday(t time.Time) time.Time {
	if !w.Weekdays[t.Weekday()] {
		addDays := w.Weekdays.DaysUntilNextDayOfWeek(t.Weekday())
		t = t.Add(24 * time.Hour * time.Duration(addDays))
	}
	return t
}

func (w *TODWeekWindow) sameDay() bool {
	return 60*w.Start.Hour+w.Start.Minute <= 60*w.End.Hour+w.End.Minute
}
