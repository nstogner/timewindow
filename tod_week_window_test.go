package timewindow_test

import (
	"fmt"
	"log"
	"time"

	"github.com/nstogner/timewindow"
)

func ExampleTODWeekWindow() {
	window, err := timewindow.ParseTODWeekWindow("10:30", "14:00", []string{"Wed", "Tue", "Thu"})
	if err != nil {
		log.Fatal(err)
	}

	// NOTE: Time in playground (docs) is fixed.
	now := time.Now()
	fmt.Println("now: ", now)

	result := window.WithinWindow(now)
	fmt.Println("within: ", result.Within)
	fmt.Println("untilStart: ", result.TTStart)
}
