package timewindow

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWeekdays(t *testing.T) {
	cases := []struct {
		name string

		weekdays Weekdays
		now      time.Weekday

		nextDayOfWeek          time.Weekday
		daysUntilNextDayOfWeek int
	}{
		{
			name:                   "empty",
			weekdays:               Weekdays(map[time.Weekday]bool{}),
			now:                    time.Thursday,
			nextDayOfWeek:          time.Weekday(-1),
			daysUntilNextDayOfWeek: 0,
		},
		{
			name: "beginning",
			weekdays: Weekdays(map[time.Weekday]bool{
				time.Wednesday: true,
				time.Thursday:  true,
			}),
			now:                    time.Tuesday,
			nextDayOfWeek:          time.Wednesday,
			daysUntilNextDayOfWeek: 1,
		},
		{
			name: "middle",
			weekdays: Weekdays(map[time.Weekday]bool{
				time.Monday:    true,
				time.Wednesday: true,
				time.Thursday:  true,
			}),
			now:                    time.Tuesday,
			nextDayOfWeek:          time.Wednesday,
			daysUntilNextDayOfWeek: 1,
		},
		{
			name: "middle-match",
			weekdays: Weekdays(map[time.Weekday]bool{
				time.Monday:    true,
				time.Wednesday: true,
				time.Thursday:  true,
			}),
			now:                    time.Wednesday,
			nextDayOfWeek:          time.Thursday,
			daysUntilNextDayOfWeek: 1,
		},
		{
			name: "end",
			weekdays: Weekdays(map[time.Weekday]bool{
				time.Monday:    true,
				time.Wednesday: true,
				time.Thursday:  true,
			}),
			now:                    time.Friday,
			nextDayOfWeek:          time.Monday,
			daysUntilNextDayOfWeek: 3,
		},
		{
			name: "same",
			weekdays: Weekdays(map[time.Weekday]bool{
				time.Tuesday: true,
			}),
			now:                    time.Tuesday,
			nextDayOfWeek:          time.Tuesday,
			daysUntilNextDayOfWeek: 7,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.nextDayOfWeek.String(), c.weekdays.NextDayOfWeek(c.now).String())
			require.Equal(t, c.daysUntilNextDayOfWeek, c.weekdays.DaysUntilNextDayOfWeek(c.now))
		})
	}
}
