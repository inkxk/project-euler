package main

import (
	"fmt"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 41
func PandigitalPrime() {
	defer util.TimeProcess(time.Now())

	max := 0
	for i := 1; i < 10000000; i++ {
		if util.IsPandigital(i) && util.IsPrimeNumber(i) && i > max {
			max = i
		}
	}

	fmt.Println("answer", max)
}
