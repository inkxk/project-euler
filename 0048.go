package main

import (
	"fmt"
	"math/big"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 48
func SelfPowers() {
	defer util.TimeProcess(time.Now())

	sum := new(big.Int)
	for i := 1; i <= 1000; i++ {
		sum = sum.Add(sum, util.Pow(i, i))
	}

	sumStr := sum.String()
	answer := sumStr[len(sumStr)-10:]

	fmt.Println("answer", answer)
}
