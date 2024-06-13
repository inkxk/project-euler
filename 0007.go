package main

import (
	"fmt"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 7
func TenThousandOnePrime() {
	defer util.TimeProcess(time.Now())
	count := 0
	primeArray := []int{}
	for i := 1; count < 10001; i++ {
		if util.IsPrimeOptimized(i) {
			primeArray = append(primeArray, i)
			count++
		}

	}

	fmt.Println("answer:", primeArray[len(primeArray)-1])
}
