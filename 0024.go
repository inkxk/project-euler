package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 24
func LexicographicPermutations() {
	defer util.TimeProcess(time.Now())

	permutations := util.GenLexicographicPermutations([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	// sort น้อย -> มาก
	sort.Strings(permutations)

	// index start from 0, so no.1000000 is index 999999
	fmt.Println("answer:", permutations[999999])
}
