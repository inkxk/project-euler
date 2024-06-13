package main

import (
	"fmt"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 3
func LargestPrimeFactor() {
	defer util.TimeProcess(time.Now())
	pfs := util.PrimeFactors(600851475143)

	fmt.Println("answer:", pfs[len(pfs)-1])
}
