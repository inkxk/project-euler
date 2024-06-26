package main

import (
	"fmt"
	"slices"
	"strconv"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 32
func PandigitalProducts() {
	defer util.TimeProcess(time.Now())

	sum := 0
	productList := []int{}

	for i := 2; i < 10000; i++ {
		// ใช้ j := i เพราะจะได้ไม่ต้องคูณตัวเดิมซ้ำ เช่น 12x483 = 483x12
		// i จะคูณได้แค่ตัวที่มากกว่า i
		for j := i; j < 10000; j++ {
			numberStr := fmt.Sprintf(`%v%v%v`, i, j, i*j)

			number, _ := strconv.Atoi(numberStr)
			if len(numberStr) == 9 && util.IsPandigital(number) {
				productList = append(productList, i*j)
			}
		}
	}

	slices.Sort(productList)
	for _, number := range slices.Compact(productList) {
		sum += number
	}

	fmt.Println("answer", sum)
}
