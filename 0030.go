package main

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 30
func DigitFifthPowers() {
	defer util.TimeProcess(time.Now())

	total := 0
	list := []int{}

	// answer of 1000000 == 1000000000, so use 1000000
	for i := 2; i < 1000000; i++ {
		sum := 0
		for _, v := range fmt.Sprintf("%v", i) {
			value, _ := strconv.Atoi(string(v))
			sum += int(math.Pow(float64(value), 5))
		}
		// fmt.Printf("sum:%v, i: %v\n", sum, i)
		// if sum of each cha**4 == i
		if sum == i {
			total += i
			list = append(list, i)
		}
	}

	fmt.Println(list)
	fmt.Println("answer:", total)
}
