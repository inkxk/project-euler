package main

import (
	"fmt"
	"math"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 12
func HighlyDivisibleTriangularNumber() {
	defer util.TimeProcess(time.Now())

	triangular := 0
	// infinity loop (no condition)
	for i := 1; ; i++ {
		dividedCount := 0
		triangular += i
		// user sqrt root and dividedCount += 2 instead of += 1 for performance
		// faster than normal loop as fuckkk
		for j := 1; float64(j) < math.Sqrt(float64(triangular)); j++ {
			if triangular%j == 0 {
				dividedCount += 2
			}
		}

		if dividedCount >= 10 {
			fmt.Println("answer:", triangular)
			break
		}
	}
}
