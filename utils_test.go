package timewindow

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestUntilTomorrow(t *testing.T) {
	cases := []struct {
		now   time.Time
		until time.Duration
	}{
		{
			now:   time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			until: 24 * time.Hour,
		},
		{
			now:   time.Date(2000, time.January, 1, 1, 0, 0, 0, time.UTC),
			until: 23 * time.Hour,
		},
		{
			now:   time.Date(2000, time.January, 1, 1, 1, 0, 0, time.UTC),
			until: 22*time.Hour + 59*time.Minute,
		},
		{
			now:   time.Date(2000, time.January, 1, 23, 59, 0, 0, time.UTC),
			until: 1 * time.Minute,
		},
	}

	for _, c := range cases {
		t.Run(c.until.String(), func(t *testing.T) {
			require.Equal(t, c.until, UntilTomorrow(c.now))
		})
	}
}
