package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 38
func PandigitalMultiples() {
	defer util.TimeProcess(time.Now())

	largest := 0

	for i := 1; i < 1000000; i++ {
		var numberStr string
		for j := 1; len(numberStr) < 9; j++ {
			numberStr += fmt.Sprintf(`%v`, i*j)
		}

		if len(numberStr) == 9 {
			number, _ := strconv.Atoi(numberStr)
			if util.IsPandigital(number) && number > largest {
				largest = number
			}
		}
	}

	fmt.Println("answer", largest)
}
