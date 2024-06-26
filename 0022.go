package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/inkxk/project-euler/util"
)

func getNameScore(name string) int {
	sum := 0
	for _, c := range name {
		// convert alphabet to index (a=0, b=1)
		score := int(c) - 64
		sum += score
	}

	return sum
}

// problem 22
func NamesScores() {
	defer util.TimeProcess(time.Now())

	// load file
	file := util.LoadFile("0022_names.txt")

	// sorting file into alphabetical order.
	slices.Sort(file)

	sum := 0
	for index, v := range file {
		name := strings.Split(v, `"`)[1]
		nameScore := getNameScore(name)
		sum += nameScore * (index + 1)
	}

	fmt.Println("answer:", sum)
}
