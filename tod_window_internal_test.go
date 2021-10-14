package timewindow

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTODWindow(t *testing.T) {
	cases := []struct {
		name string

		window TODWindow
		now    time.Time

		withinWindow bool
		duration     time.Duration
		sameDay      bool
		startTime    time.Time
		endTime      time.Time
	}{
		{
			name: "one-hour-until-same-day-window",
			window: TODWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
			},
			now: time.Date(2000, time.January, 1, 9, 0, 0, 0, time.UTC),

			withinWindow: false,
			duration:     time.Hour,
			sameDay:      true,
			startTime:    time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:      time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
		},
		{
			name: "one-hour-until-non-same-day-window",
			window: TODWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 8, Minute: 0},
			},
			now: time.Date(2000, time.January, 1, 9, 0, 0, 0, time.UTC),

			withinWindow: false,
			duration:     time.Hour,
			sameDay:      false,
			startTime:    time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:      time.Date(2000, time.January, 2, 8, 0, 0, 0, time.UTC),
		},
		{
			name: "one-hour-within-same-day-window",
			window: TODWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
			},
			now: time.Date(2000, time.January, 1, 11, 0, 0, 0, time.UTC),

			withinWindow: true,
			duration:     23 * time.Hour,
			sameDay:      true,
			startTime:    time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:      time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
		},
		{
			name: "one-hour-after-same-day-window",
			window: TODWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
			},
			now: time.Date(2000, time.January, 1, 21, 0, 0, 0, time.UTC),

			withinWindow: false,
			duration:     13 * time.Hour,
			sameDay:      true,
			startTime:    time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:      time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
		},
		{
			name: "one-hour-within-non-same-day-window",
			window: TODWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 8, Minute: 0},
			},
			now: time.Date(2000, time.January, 1, 11, 0, 0, 0, time.UTC),

			withinWindow: true,
			duration:     23 * time.Hour,
			sameDay:      false,
			startTime:    time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:      time.Date(2000, time.January, 2, 8, 0, 0, 0, time.UTC),
		},
		{
			name: "one-hour-after-non-same-day-window",
			window: TODWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 3, Minute: 0},
			},
			now: time.Date(2000, time.January, 1, 6, 0, 0, 0, time.UTC),

			withinWindow: false,
			duration:     4 * time.Hour,
			sameDay:      false,
			startTime:    time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:      time.Date(2000, time.January, 2, 3, 0, 0, 0, time.UTC),
		},
		{
			name: "on-start",
			window: TODWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
			},
			now: time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),

			withinWindow: true,
			duration:     0,
			sameDay:      true,
			startTime:    time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:      time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
		},
		{
			name: "on-end",
			window: TODWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
			},
			now: time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),

			withinWindow: false,
			duration:     14 * time.Hour,
			sameDay:      true,
			startTime:    time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:      time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
		},
		{
			name: "20-minutes-before",
			window: TODWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
			},
			now: time.Date(2000, time.January, 1, 9, 40, 0, 0, time.UTC),

			withinWindow: false,
			duration:     20 * time.Minute,
			sameDay:      true,
			startTime:    time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:      time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			withinWindow, nextWindow := c.window.WithinWindow(c.now)
			require.Equal(t, c.withinWindow, withinWindow)
			require.Equal(t, c.duration.String(), nextWindow.String())
			require.Equal(t, c.sameDay, c.window.sameDay())
			require.Equal(t, c.startTime.String(), c.window.StartTime(c.now).String())
			require.Equal(t, c.endTime.String(), c.window.EndTime(c.now).String())
		})
	}
}
