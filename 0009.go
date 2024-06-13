package main

import (
	"fmt"
	"time"
)

// problem 9
func SpecialPythagoreanTriplet() {
	defer TimeIt(time.Now())
	for a := 1; a < 1000; a++ {
		for b := 1; b < 1000; b++ {
			c := 1000 - a - b
			if (a*a)+(b*b) == (c * c) {
				fmt.Printf("a: %d, b: %d, c: %d\n", a, b, c)
				fmt.Println("answer:", a*b*c)
			}
		}
	}
}
