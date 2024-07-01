package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 52
func PermutedMultiples() {
	defer util.TimeProcess(time.Now())

out:
	for i := 100000; i < 1000000; i++ {
		isPermuted := true
		for j := 2; j <= 6; j++ {

			iSort := util.SortString(strconv.Itoa(i))
			iMulJSort := util.SortString(strconv.Itoa(i * j))

			if iSort != iMulJSort {
				isPermuted = false
				continue out
			}
		}

		if isPermuted {
			fmt.Println("answer", i)
			break out
		}
	}
}
