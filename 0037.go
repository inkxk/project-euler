package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/inkxk/project-euler/util"
)

// problem 37
func TruncatablePrimes() {
	defer util.TimeProcess(time.Now())

	sum := 0
	primes := []int{}

out:
	for i := 10; len(primes) < 11; i++ {
		if util.IsPrimeNumber(i) {
			iStr := fmt.Sprintf("%v", i)

			isTruncatablePrimes := true
			for j := 1; j < len(iStr); j++ {
				// ตัด string left and right ทีละ 1 ตัว
				subLeft, _ := strconv.Atoi(iStr[:len(iStr)-j])
				subRight, _ := strconv.Atoi(iStr[j:])

				// ถ้าไม่เป็น prime, ข้าม i ตัวนี้ไปทำตัวอื่นต่อ
				if !util.IsPrimeNumber(subLeft) || !util.IsPrimeNumber(subRight) {
					isTruncatablePrimes = false
					continue out
				}
			}

			// ถ้า loop ผ่านบรรทัดที่ 30 มาได้ทุกตัว จะมาเข้าบรรทัดนี้
			if isTruncatablePrimes {
				primes = append(primes, i)
				sum += i
			}
		}
	}

	fmt.Println("primes", primes)
	fmt.Println("answer:", sum)
}
