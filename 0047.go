package main

import (
	"fmt"
	"slices"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 47
func DistinctPrimesFactors() {
	defer util.TimeProcess(time.Now())

	primes := []int{}
	for i := 2; i < 1000000; i++ {
		primeFactors := util.PrimeFactors(i)

		slices.Sort(primeFactors)
		primeFactors = slices.Compact(primeFactors)

		if len(primeFactors) == 4 {
			primes = append(primes, i)
		}

		if len(primes) == 4 {
			if util.IsArrayConsecutive(primes) {
				fmt.Println("primes:", primes)
				fmt.Println("answer:", primes[0])
				break
			} else {
				primes = []int{}
			}
		}
	}

}
