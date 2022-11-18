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

		result               WindowResult
		resultTTWindowChange time.Duration
		startTime            time.Time
		endTime              time.Time
		followingStartTime   time.Time
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

			result: WindowResult{
				Within:  false,
				TTStart: time.Hour + 24*time.Hour,
				TTEnd:   0,
			},
			resultTTWindowChange: time.Hour + 24*time.Hour,
			startTime:            time.Date(2000, time.January, 2, 10, 0, 0, 0, time.UTC),
			endTime:              time.Date(2000, time.January, 2, 20, 0, 0, 0, time.UTC),
			followingStartTime:   time.Date(2000, time.January, 9, 10, 0, 0, 0, time.UTC),
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

			result: WindowResult{
				Within:  false,
				TTStart: time.Hour,
				TTEnd:   0,
			},
			resultTTWindowChange: time.Hour,
			startTime:            time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:              time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
			followingStartTime:   time.Date(2000, time.January, 8, 10, 0, 0, 0, time.UTC),
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

			result: WindowResult{
				Within:  true,
				TTStart: 0,
				TTEnd:   10 * time.Hour,
			},
			resultTTWindowChange: 10 * time.Hour,
			startTime:            time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:              time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
			followingStartTime:   time.Date(2000, time.January, 8, 10, 0, 0, 0, time.UTC),
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

			result: WindowResult{
				Within:  false,
				TTStart: 13 * time.Hour,
				TTEnd:   0,
			},
			resultTTWindowChange: 13 * time.Hour,
			startTime:            time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:              time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
			followingStartTime:   time.Date(2000, time.January, 2, 10, 0, 0, 0, time.UTC),
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

			result: WindowResult{
				Within:  false,
				TTStart: 24*time.Hour - 2*time.Hour,
				TTEnd:   0,
			},
			resultTTWindowChange: 24*time.Hour - 2*time.Hour,
			startTime:            time.Date(2000, time.January, 2, 10, 0, 0, 0, time.UTC),
			endTime:              time.Date(2000, time.January, 2, 20, 0, 0, 0, time.UTC),
			followingStartTime:   time.Date(2000, time.January, 9, 10, 0, 0, 0, time.UTC),
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

			result: WindowResult{
				Within:  true,
				TTStart: 24*7*time.Hour - 1*time.Hour,
				TTEnd:   9 * time.Hour,
			},
			resultTTWindowChange: 9 * time.Hour,
			startTime:            time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:              time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
			followingStartTime:   time.Date(2000, time.January, 8, 10, 0, 0, 0, time.UTC),
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

			result: WindowResult{
				Within:  false,
				TTStart: 24*7*time.Hour - 10*time.Hour,
				TTEnd:   0,
			},
			resultTTWindowChange: 24*7*time.Hour - 10*time.Hour,
			startTime:            time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:              time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
			followingStartTime:   time.Date(2000, time.January, 8, 10, 0, 0, 0, time.UTC),
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

			result: WindowResult{
				Within:  false,
				TTStart: 24*7*time.Hour - 11*time.Hour,
				TTEnd:   0,
			},
			resultTTWindowChange: 24*7*time.Hour - 11*time.Hour,
			startTime:            time.Date(2000, time.January, 1, 10, 0, 0, 0, time.UTC),
			endTime:              time.Date(2000, time.January, 1, 20, 0, 0, 0, time.UTC),
			followingStartTime:   time.Date(2000, time.January, 8, 10, 0, 0, 0, time.UTC),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.nowWeekday, c.now.Weekday())

			require.Equal(t, c.startTime.String(), c.window.StartTime(c.now).String())
			require.Equal(t, c.endTime.String(), c.window.EndTime(c.now).String())
			require.Equal(t, c.followingStartTime.String(), c.window.FollowingStartTime(c.now).String())

			result := c.window.WithinWindow(c.now)
			require.Equal(t, c.result.Within, result.Within)
			require.Equal(t, c.result.TTStart.String(), result.TTStart.String())
			require.Equal(t, c.result.TTEnd.String(), result.TTEnd.String())
			require.Equal(t, c.result.TTWithinChange().String(), result.TTWithinChange().String())
		})
	}
}
