package main

import (
	"fmt"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 63
func PowerfulDigitCounts() {
	defer util.TimeProcess(time.Now())

	count := 0
	for i := 1; i < 10; i++ {
		for j := 1; j < 100; j++ {
			pow := util.Pow(i, j)
			length := len(pow.String())
			if length == j {
				count++
			}
		}

	}

	fmt.Println("answer", count)
}
