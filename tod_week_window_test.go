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

	within, untilStart := window.WithinWindow(time.Now())
	fmt.Println("within: ", within)
	fmt.Println("untilStart: ", untilStart)
}
