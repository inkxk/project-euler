package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 34
func DigitFactorials() {
	defer util.TimeProcess(time.Now())

	totalSum := 0

	for i := 3; i < 1000000; i++ {
		sum := 0
		for _, v := range fmt.Sprintf("%d", i) {
			vInt, _ := strconv.Atoi(string(v))
			// sum += factorial of each cha in i, convert big.Int to int
			sum += int(new(big.Int).MulRange(1, int64(vInt)).Int64())
		}

		if sum == i {
			totalSum += sum
		}
	}

	fmt.Println("answer:", totalSum)

}
