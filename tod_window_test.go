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

	within, untilStart := window.WithinWindow(time.Now())
	fmt.Println("within: ", within)
	fmt.Println("untilStart: ", untilStart)
}