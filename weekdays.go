package timewindow

import (
	"sort"
	"time"
)

type Weekdays map[time.Weekday]bool

// NextDayOfWeek returns the next matching day of the week.
// It will return the current day if the current day matches AND
// there are no other matching days.
func (w Weekdays) NextDayOfWeek(today time.Weekday) time.Weekday {
	if len(w) == 0 {
		return -1
	}

	var days []int
	for day := range w {
		days = append(days, int(day))
	}
	sort.Ints(days)

	// Wrap around.
	// 2 5 (6) --> 2
	if int(today) > days[len(days)-1] {
		return time.Weekday(days[0])
	}

	// (1) 2 5 --> 2
	// 2 (3) 5 --> 5
	for _, day := range days {
		if day > int(today) {
			return time.Weekday(day)
		}
	}

	return today
}

// DaysUntilNextDayOfWeek calculates the next day of the week that matches
// and returns the number of days until then.
func (w Weekdays) DaysUntilNextDayOfWeek(today time.Weekday) int {
	if len(w) == 0 {
		return 0
	}

	next := w.NextDayOfWeek(today)
	if next > today {
		return int(next - today)
	} else if next < today {
		return (7 - int(today)) + int(next)
	} else {
		return 7
	}
}
