package main

import (
	"fmt"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 1
func MultiplesOf3Or5() {
	defer util.TimeProcess(time.Now())
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Println("answer:", sum)
}
