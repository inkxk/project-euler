package main

import (
	"fmt"
	"slices"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 29
func DistinctPowers() {
	defer util.TimeProcess(time.Now())

	powerList := []string{}
	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			power := util.Pow(a, b)
			powerList = append(powerList, power.String())
		}
	}

	// remove duplicated
	slices.Sort(powerList)
	powerList = slices.Compact(powerList)

	fmt.Println("powerList:", powerList)
	fmt.Println("answer:", len(powerList))
}
