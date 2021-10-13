package timewindow

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseTimeOfDayHappyPath(t *testing.T) {
	cases := []struct {
		s   string
		tod TimeOfDay
	}{
		{
			s:   "00:00",
			tod: TimeOfDay{Hour: 0, Minute: 0},
		},
		{
			s:   "01:02",
			tod: TimeOfDay{Hour: 1, Minute: 2},
		},
		{
			s:   "11:22",
			tod: TimeOfDay{Hour: 11, Minute: 22},
		},
	}

	for _, c := range cases {
		t.Run(c.s, func(t *testing.T) {
			tod, err := ParseTimeOfDay(c.s)
			require.NoError(t, err)
			require.Equal(t, c.tod, tod)
		})
	}
}

func TestParseTimeOfDaySadPath(t *testing.T) {
	cases := []struct {
		name    string
		s       string
		errText string
	}{
		{
			name:    "empty",
			s:       "",
			errText: "invalid format",
		},
		{
			name:    "negative-hour",
			s:       "-12:00",
			errText: "invalid hour",
		},
		{
			name:    "negative-minutes",
			s:       "00:-12",
			errText: "invalid minute",
		},
		{
			name:    "too-many-hours",
			s:       "25:00",
			errText: "invalid hour",
		},
		{
			name:    "too-many-minutes",
			s:       "00:61",
			errText: "invalid minute",
		},
	}

	for _, c := range cases {
		t.Run(c.s, func(t *testing.T) {
			_, err := ParseTimeOfDay(c.s)
			require.Contains(t, err.Error(), c.errText)
		})
	}
}
