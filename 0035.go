package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 35
func CircularPrimes() {
	defer util.TimeProcess(time.Now())

	circularPrimes := []int{}
	for i := 1; i < 1000000; i++ {
		stingArray := util.PossibleRotation(fmt.Sprintf("%d", i))
		IsCircularPrimes := true

	in:
		// check every element in array
		for _, v := range stingArray {
			vInt, _ := strconv.Atoi(v)

			// if 1 of them doesn't prime, break
			if !util.IsPrimeNumber(vInt) {
				IsCircularPrimes = false
				break in
			}
		}

		if IsCircularPrimes {
			circularPrimes = append(circularPrimes, i)
		}
	}

	fmt.Println("circularPrimes:", circularPrimes)
	fmt.Println("answer:", len(circularPrimes))
}
