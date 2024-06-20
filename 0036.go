package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 36
func DoubleBasePalindromes() {
	defer util.TimeProcess(time.Now())

	sum := 0

	for i := 1; i < 1000000; i++ {
		// convert number from base10 to base2
		base2 := strconv.FormatInt(int64(i), 2)
		if util.IsPalindrome(fmt.Sprintf("%v", i)) && util.IsPalindrome(base2) {
			sum += i
		}
	}

	fmt.Println("answer:", sum)
}
