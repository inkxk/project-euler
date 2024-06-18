package main

import (
	"fmt"
	"math/big"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 15
func LatticePaths() {
	defer util.TimeProcess(time.Now())

	n := 20

	// formula
	// factorial(2*n) / (factorial(n) * factorial(n))
	frist := new(big.Int).MulRange(1, int64(n*2))
	second := new(big.Int).MulRange(1, int64(n))
	third := new(big.Int).Mul(second, second)
	answer := new(big.Int).Div(frist, third)
	// answer := new(big.Int).Div(new(big.Int).MulRange(1, int64(n*2)), new(big.Int).Mul(new(big.Int).MulRange(1, int64(n)), new(big.Int).MulRange(1, int64(n))))

	fmt.Println("answer:", answer)

}
