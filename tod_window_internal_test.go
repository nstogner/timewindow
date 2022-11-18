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

		result               WindowResult
		resultTTWindowChange time.Duration
		sameDay              bool
		startTime            time.Time
		endTime              time.Time
		followingStartTime   time.Time
	}{
		{
			name: "one-hour-until-same-day-window",
			window: TODWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
			},
			now: time.Date(2000, time.January, 1, 9, 0, 0, 0, time.UTC),

			result: WindowResult{
				Within:  false,
				TTStart: time.Hour,
				TTEnd:   0,
			},
			resultTTWindowChange: time.Hour,
			sameDay:              true,
			startTime:            time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:              time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
			followingStartTime:   time.Date(2000, time.January, 2, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "one-hour-until-non-same-day-window",
			window: TODWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 8, Minute: 0},
			},
			now: time.Date(2000, time.January, 1, 9, 0, 0, 0, time.UTC),

			result: WindowResult{
				Within:  false,
				TTStart: time.Hour,
				TTEnd:   0,
			},
			resultTTWindowChange: time.Hour,
			sameDay:              false,
			startTime:            time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:              time.Date(2000, time.January, 2, 8, 0, 0, 0, time.UTC),
			followingStartTime:   time.Date(2000, time.January, 2, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "20-minutes-before-window",
			window: TODWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
			},
			now: time.Date(2000, time.January, 1, 9, 40, 0, 0, time.UTC),

			result: WindowResult{
				Within:  false,
				TTStart: 20 * time.Minute,
				TTEnd:   0,
			},
			resultTTWindowChange: 20 * time.Minute,
			sameDay:              true,
			startTime:            time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:              time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
			followingStartTime:   time.Date(2000, time.January, 2, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "on-start",
			window: TODWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
			},
			now: time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),

			result: WindowResult{
				Within:  true,
				TTStart: 0,
				TTEnd:   10 * time.Hour,
			},
			resultTTWindowChange: 10 * time.Hour,
			sameDay:              true,
			startTime:            time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:              time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
			followingStartTime:   time.Date(2000, time.January, 2, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "one-hour-within-same-day-window",
			window: TODWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
			},
			now: time.Date(2000, time.January, 1, 11, 0, 0, 0, time.UTC),

			result: WindowResult{
				Within:  true,
				TTStart: 23 * time.Hour,
				TTEnd:   9 * time.Hour,
			},
			resultTTWindowChange: 9 * time.Hour,
			sameDay:              true,
			startTime:            time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:              time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
			followingStartTime:   time.Date(2000, time.January, 2, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "one-hour-within-non-same-day-window",
			window: TODWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 8, Minute: 0},
			},
			now: time.Date(2000, time.January, 1, 11, 0, 0, 0, time.UTC),

			result: WindowResult{
				Within:  true,
				TTStart: 23 * time.Hour,
				TTEnd:   21 * time.Hour,
			},
			resultTTWindowChange: 21 * time.Hour,
			sameDay:              false,
			startTime:            time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:              time.Date(2000, time.January, 2, 8, 0, 0, 0, time.UTC),
			followingStartTime:   time.Date(2000, time.January, 2, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "one-hour-after-same-day-window",
			window: TODWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
			},
			now: time.Date(2000, time.January, 1, 21, 0, 0, 0, time.UTC),

			result: WindowResult{
				Within:  false,
				TTStart: 13 * time.Hour,
				TTEnd:   0,
			},
			resultTTWindowChange: 13 * time.Hour,
			sameDay:              true,
			startTime:            time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:              time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
			followingStartTime:   time.Date(2000, time.January, 2, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "one-hour-after-non-same-day-window",
			window: TODWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 3, Minute: 0},
			},
			now: time.Date(2000, time.January, 1, 6, 0, 0, 0, time.UTC),

			result: WindowResult{
				Within:  false,
				TTStart: 4 * time.Hour,
				TTEnd:   0,
			},
			resultTTWindowChange: 4 * time.Hour,
			sameDay:              false,
			startTime:            time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:              time.Date(2000, time.January, 2, 3, 0, 0, 0, time.UTC),
			followingStartTime:   time.Date(2000, time.January, 2, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "on-end",
			window: TODWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
			},
			now: time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),

			result: WindowResult{
				Within:  false,
				TTStart: 14 * time.Hour,
				TTEnd:   0,
			},
			resultTTWindowChange: 14 * time.Hour,
			sameDay:              true,
			startTime:            time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:              time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
			followingStartTime:   time.Date(2000, time.January, 2, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "one-hour-after-end",
			window: TODWindow{
				Start: TOD{Hour: 10, Minute: 0},
				End:   TOD{Hour: 20, Minute: 0},
			},
			now: time.Date(2000, time.January, 1, 21, 0, 0, 0, time.UTC),

			result: WindowResult{
				Within:  false,
				TTStart: 13 * time.Hour,
				TTEnd:   0,
			},
			resultTTWindowChange: 13 * time.Hour,
			sameDay:              true,
			startTime:            time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:              time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
			followingStartTime:   time.Date(2000, time.January, 2, 10, 0, 0, 0, time.UTC),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := c.window.WithinWindow(c.now)
			require.Equal(t, c.startTime.String(), c.window.StartTime(c.now).String())
			require.Equal(t, c.endTime.String(), c.window.EndTime(c.now).String())
			require.Equal(t, c.followingStartTime.String(), c.window.FollowingStartTime(c.now).String())

			require.Equal(t, c.result.Within, result.Within)
			require.Equal(t, c.result.TTStart.String(), result.TTStart.String())
			require.Equal(t, c.result.TTEnd.String(), result.TTEnd.String())
			require.Equal(t, c.result.TTWithinChange().String(), result.TTWithinChange().String())
			require.Equal(t, c.sameDay, c.window.sameDay())
		})
	}
}
