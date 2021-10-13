package timewindow

import "time"

func WithinWindow(now, start, end time.Time) (bool, time.Duration) {
	// -----|---------------
	// ---start-----end-----
	if now.Equal(start) {
		return true, 0
	}
	// --------------|------
	// ---start-----end-----
	if now.Equal(end) {
		return false, 24 * time.Hour
	}

	startMinusNow := start.Sub(now)

	// ----------|----------
	// ---start-----end-----
	if now.After(start) && now.Before(end) {
		return true, startMinusNow
	}

	// ------------------|--
	// ---start-----end-----
	if now.After(end) {
		return false, 24*time.Hour + startMinusNow
	}

	// -|-------------------
	// ---start-----end-----
	return false, startMinusNow
}
