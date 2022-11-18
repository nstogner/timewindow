package timewindow

import "time"

// WithinWindow returns true if within a window. It also returns the time until the next
// window starts.
func WithinWindow(now, start, end, followingStart time.Time) WindowResult {
	// BOD    = Beginning of Day
	// EOD    = End of Day
	// start  = Start of current window
	// end    = End of current window
	// fStart = Start of following window
	// D      = Day
	// N      = Now
	// TTS    = Caluculated time until start
	// TTE    = Caluculated time until end
	//
	// BOD------------start----------end---------EOD------D+1--D+N----fStart----
	//         N-(TTS)->|
	//                  |    N-------------------(TTS)------------------>|
	//                  |    N-(TTE)->|
	//                                     N------------(TTS)----------->|

	// -|-------------------
	// ---start-----end-----
	if now.Before(start) {
		return WindowResult{
			Within:  false,
			TTStart: start.Sub(now),
		}
	}

	// -----|---------------
	// ---start-----end-----
	if now.Equal(start) {
		return WindowResult{
			Within:  true,
			TTStart: 0,
			TTEnd:   end.Sub(now),
		}
	}

	// ----------|----------
	// ---start-----end-----
	if now.After(start) && now.Before(end) {
		return WindowResult{
			Within:  true,
			TTStart: followingStart.Sub(now),
			TTEnd:   end.Sub(now),
		}
	}

	// --------------|------
	// ------------------|--
	// ---start-----end-----
	if now.Equal(end) || now.After(end) {
		return WindowResult{
			Within:  false,
			TTStart: followingStart.Sub(now),
		}
	}

	// This should never be reached.
	return WindowResult{}
}

type WindowResult struct {
	Within  bool
	TTStart time.Duration
	TTEnd   time.Duration
}

// TTWithinChange is the Time Til there is a change in the .Within window result.
func (r WindowResult) TTWithinChange() time.Duration {
	if r.Within {
		return r.TTEnd
	} else {
		return r.TTStart
	}
}
