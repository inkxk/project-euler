package main

import (
	"fmt"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 5
func SmallestMultiple() {
	defer util.TimeProcess(time.Now())
	dividerRange := 20
	smallest := dividerRange
	for {
		// plus value with highest number to divided by
		smallest += dividerRange
		found := true
		// divided by i from biggest to smallest, so it's faster
		// last number to divided by is 2
		for i := dividerRange; i >= 2; i-- {
			if smallest%i != 0 {
				found = false
				// if can't divided, break loop
				break
			}
		}
		// if value can divided by 20 to 2, it will pass through the loop without break
		if found {
			fmt.Println("answer:", smallest)
			break
		}
	}
}
