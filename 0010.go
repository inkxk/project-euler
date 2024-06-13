package main

import (
	"fmt"
	"math"
	"time"
)

func isPrimeOptimized(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// problem 10
func SummationOfPrimes() {
	defer TimeIt(time.Now())
	sum := 0
	for i := 1; i < 2000000; i++ {
		if isPrimeOptimized(i) {
			sum += i
		}
	}

	fmt.Println("answer:", sum)
}
