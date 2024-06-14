package main

import (
	"fmt"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 25
func ThousandDigitFibonacciNumber() {
	defer util.TimeProcess(time.Now())

	_, index := util.GenFibonacciByDigit(1000)
	fmt.Println("answer:", index)
}
