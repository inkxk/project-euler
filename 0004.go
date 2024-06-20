package main

import (
	"fmt"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 4
func LargestPalindromeProduct() {
	defer util.TimeProcess(time.Now())
	max := 0
	for i := 1; i < 1000; i++ {
		for j := 1; j < 1000; j++ {
			product := i * j
			if util.IsPalindrome(fmt.Sprintf("%v", product)) && product > max {
				max = product
			}
		}
	}

	fmt.Println("answer:", max)
}
