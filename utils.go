package timewindow

import "time"

// UntilTomorrow returns the amount of time until midnight.
func UntilTomorrow(now time.Time) time.Duration {
	tomorrow := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Add(24 * time.Hour)
	return tomorrow.Sub(now)
}
