package timewindow

import "time"

// WithinWindow returns true if within a window. It also returns the time until the next
// window starts.
func WithinWindow(now, start, end, followingStart time.Time) (bool, time.Duration) {
	// BOD    = Beginning of Day
	// EOD    = End of Day
	// start  = Start of current window
	// end    = End of current window
	// fStart = Start of following window
	// D      = Day
	// N      = Now
	// T      = Caluculated time until start
	//
	// BOD------------start----------end---------EOD------D+1--D+N----fStart----
	//         N--(T)-->|
	//                  |       N-----------------(T)------------------->|
	//                                     N------------(T)------------->|

	// -|-------------------
	// ---start-----end-----
	if now.Before(start) {
		return false, start.Sub(now)
	}

	// -----|---------------
	// ---start-----end-----
	if now.Equal(start) {
		return true, 0
	}

	// ----------|----------
	// ---start-----end-----
	if now.After(start) && now.Before(end) {
		return true, followingStart.Sub(now)
	}

	// --------------|------
	// ------------------|--
	// ---start-----end-----
	if now.Equal(end) || now.After(end) {
		return false, followingStart.Sub(now)
	}

	// This should never be reached.
	return false, 0
}
