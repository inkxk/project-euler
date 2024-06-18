package main

import (
	"fmt"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 10
func SummationOfPrimes() {
	defer util.TimeProcess(time.Now())
	sum := 0
	for i := 1; i < 2000000; i++ {
		if util.IsPrimeNumber(i) {
			sum += i
		}
	}

	fmt.Println("answer:", sum)
}
