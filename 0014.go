package main

import (
	"fmt"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 14
func LongestCollatzSequence() {
	defer util.TimeProcess(time.Now())

	max := 0
	bestStart := 0

	for i := 2; i < 1000000; i++ {
		start := i
		maxRound := 0
		for start > 1 {
			if start%2 == 0 {
				start = start / 2
			} else {
				start = (3 * start) + 1
			}
			maxRound += 1
		}

		if maxRound > max {
			max = maxRound
			bestStart = i
		}
	}

	fmt.Println("answer:", bestStart)
}
