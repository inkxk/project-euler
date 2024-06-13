package main

import "fmt"

// problem 9
func SpecialPythagoreanTriplet() {
	for a := 0; a < 1000; a++ {
		for b := 0; b < 1000; b++ {
			c := 1000 - a - b
			if (a*a)+(b*b) == (c * c) {
				fmt.Println(a, b, c)
				fmt.Println("answer:", a*b*c)
			}
		}
	}
}
