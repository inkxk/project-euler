package main

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 16
func LargeSum() {
	defer util.TimeProcess(time.Now())
	sum := 0
	powerFloat := math.Pow(2, 1000)
	// float64 to string
	powerString := strconv.FormatFloat(powerFloat, 'f', 0, 64)
	for _, v := range powerString {
		// rune -> string -> int
		num, _ := strconv.Atoi(string(v))
		sum += num
	}
	fmt.Println("answer:", sum)
}
