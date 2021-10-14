package timewindow

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTODWeekWindow(t *testing.T) {
	cases := []struct {
		name string

		window     TODWeekWindow
		now        time.Time
		nowWeekday time.Weekday

		withinWindow  bool
		duration      time.Duration
		startTime     time.Time
		endTime       time.Time
		nextStartTime time.Time
	}{
		{
			name: "one-hour-plus-one-day-until-window",
			window: TODWeekWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
				Weekdays: Weekdays{
					time.Sunday: true,
				},
			},
			now:        time.Date(2000, time.January, 1, 9, 0, 0, 0, time.UTC),
			nowWeekday: time.Saturday,

			withinWindow:  false,
			duration:      time.Hour + 24*time.Hour,
			startTime:     time.Date(2000, time.January, 2, 10, 0, 0, 0, time.UTC),
			endTime:       time.Date(2000, time.January, 2, 20, 0, 0, 0, time.UTC),
			nextStartTime: time.Date(2000, time.January, 9, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "one-hour-until-window",
			window: TODWeekWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
				Weekdays: Weekdays{
					time.Saturday: true,
				},
			},
			now:        time.Date(2000, time.January, 1, 9, 0, 0, 0, time.UTC),
			nowWeekday: time.Saturday,

			withinWindow:  false,
			duration:      time.Hour,
			startTime:     time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:       time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
			nextStartTime: time.Date(2000, time.January, 8, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "on-start",
			window: TODWeekWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
				Weekdays: Weekdays{
					time.Saturday: true,
				},
			},
			now:        time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			nowWeekday: time.Saturday,

			withinWindow:  true,
			duration:      0,
			startTime:     time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:       time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
			nextStartTime: time.Date(2000, time.January, 8, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "one-hour-after-window",
			window: TODWeekWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
				Weekdays: Weekdays{
					time.Saturday: true,
					time.Sunday:   true,
				},
			},
			now:        time.Date(2000, time.January, 1, 21, 0, 0, 0, time.UTC),
			nowWeekday: time.Saturday,

			withinWindow:  false,
			duration:      13 * time.Hour,
			startTime:     time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:       time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
			nextStartTime: time.Date(2000, time.January, 2, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "two-hours-into-window-of-next-day",
			window: TODWeekWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
				Weekdays: Weekdays{
					time.Sunday: true,
				},
			},
			now:        time.Date(2000, time.January, 1, 12, 0, 0, 0, time.UTC),
			nowWeekday: time.Saturday,

			withinWindow:  false,
			duration:      24*time.Hour - 2*time.Hour,
			startTime:     time.Date(2000, time.January, 2, 10, 0, 0, 0, time.UTC),
			endTime:       time.Date(2000, time.January, 2, 20, 0, 0, 0, time.UTC),
			nextStartTime: time.Date(2000, time.January, 9, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "one-hour-within-window",
			window: TODWeekWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
				Weekdays: Weekdays{
					time.Saturday: true,
				},
			},
			now:        time.Date(2000, time.January, 1, 11, 0, 0, 0, time.UTC),
			nowWeekday: time.Saturday,

			withinWindow:  true,
			duration:      24*7*time.Hour - 1*time.Hour,
			startTime:     time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:       time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
			nextStartTime: time.Date(2000, time.January, 8, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "on-end",
			window: TODWeekWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
				Weekdays: Weekdays{
					time.Saturday: true,
				},
			},
			now:        time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
			nowWeekday: time.Saturday,

			withinWindow:  false,
			duration:      7*24*time.Hour - 10*time.Hour,
			startTime:     time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:       time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
			nextStartTime: time.Date(2000, time.January, 8, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "one-hour-after-end",
			window: TODWeekWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
				Weekdays: Weekdays{
					time.Saturday: true,
				},
			},
			now:        time.Date(2000, time.January, 1, 21, 0, 0, 0, time.UTC),
			nowWeekday: time.Saturday,

			withinWindow:  false,
			duration:      7*24*time.Hour - 11*time.Hour,
			startTime:     time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:       time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
			nextStartTime: time.Date(2000, time.January, 8, 10, 0, 0, 0, time.UTC),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.nowWeekday, c.now.Weekday())

			require.Equal(t, c.startTime.String(), c.window.StartTime(c.now).String())
			require.Equal(t, c.endTime.String(), c.window.EndTime(c.now).String())
			require.Equal(t, c.nextStartTime.String(), c.window.NextStartTime(c.now).String())

			withinWindow, untilStart := c.window.WithinWindow(c.now)
			require.Equal(t, c.withinWindow, withinWindow)
			require.Equal(t, c.duration.String(), untilStart.String())
		})
	}
}
