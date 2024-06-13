package main

import (
	"fmt"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 6
func SumSquareDifference() {
	defer util.TimeProcess(time.Now())
	squareSum := 0
	sum := 0
	for i := 1; i <= 100; i++ {
		squareSum += i * i
		sum += i
	}

	fmt.Println("answer:", (sum*sum)-squareSum)
}
