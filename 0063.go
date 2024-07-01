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
		for pow := 1; pow < 100; pow++ {
			powResult := util.Pow(i, pow)
			length := len(powResult.String())
			if length == pow {
				count++
			}
		}
	}

	fmt.Println("answer", count)
}
