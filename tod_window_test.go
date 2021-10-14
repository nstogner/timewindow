package timewindow_test

import (
	"fmt"
	"log"
	"time"

	"github.com/nstogner/timewindow"
)

func ExampleTODWindow() {
	window, err := timewindow.ParseTODWindow("10:30", "14:00")
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	fmt.Println("now: ", now)

	within, untilStart := window.WithinWindow(now)
	fmt.Println("within: ", within)
	fmt.Println("untilStart: ", untilStart)
}
