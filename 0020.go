package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 20
func FactorialDigitSum() {
	defer util.TimeProcess(time.Now())

	x := new(big.Int)
	sum := 0

	// factorial
	x.MulRange(1, 100)

	for _, v := range fmt.Sprintf("%d", x) {
		vInt, _ := strconv.Atoi(string(v))
		sum += vInt
	}

	fmt.Println("answer:", sum)
}
