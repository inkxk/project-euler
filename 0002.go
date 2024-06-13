package main

import (
	"fmt"
	"time"
)

// problem 2
func EvenFibonacciNumbers() {
	defer TimeIt(time.Now())
	fiboArray := []int{1, 2}
	sum := 2
	// loop while last element of fiboArray < 4,000,000
	for i := 0; fiboArray[len(fiboArray)-1] < 4000000; i++ {
		fibo := fiboArray[i] + fiboArray[i+1]
		fiboArray = append(fiboArray, fibo)
		// sum only even-valued fibo
		if fibo%2 == 0 {
			sum += fibo
		}
	}

	fmt.Println("answer:", sum)
}
