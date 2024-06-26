package util

import (
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

// เช็คว่าเป็นจำนวนเฉพาะหรือไม่
func IsPrimeNumber(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// หาตัวหารทั้งของจำนวนเฉพาะ
func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

// สร้างเลข fibonaci ที่มีหลักตามเลข digit เช่น 100 หลัก
// และบอก index ของเลข fibonaci ว่าเป็นตัวที่เท่าไร นับตั้งแต่ fibonaci ตัวแรก
func GenFibonacciByDigit(digit int) (*big.Int, int) {
	a := big.NewInt(0)
	b := big.NewInt(1)

	// Initialize limit as 10^99, the smallest integer with {digit} digits.
	var limit big.Int

	limit.Exp(big.NewInt(10), big.NewInt(int64(digit-1)), nil)

	// index of a
	count := 0

	// Loop while a is smaller than 1e100.
	for a.Cmp(&limit) < 0 {
		// Compute the next Fibonacci number, storing it in a.
		a.Add(a, b)
		// Swap a and b so that b is the next number in the sequence.
		a, b = b, a
		count++
	}

	return a, count // a = {digit}-digit Fibonacci number
}

func RootOf(value int, root int) float64 {
	return math.Round(math.Pow((float64(value)), (1.0 / float64(root))))
}

// a ยกกำลัง b
func Pow(a, b int) *big.Int {
	mul := big.NewInt(int64(a))
	result := big.NewInt(int64(a))
	// a * a เป็นจำนวน b ครั้ง
	for i := 1; i < b; i++ {
		result.Mul(result, mul)
	}
	return result
}

// ใช้สำหรับ function GenPermutations
func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

// ใช้สำหรับ function GenPermutations
func getPerm(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

// สลับตำแหน่งตัวใน array โดย return ทุกตัวที่เป็นไปได้
func GenLexicographicPermutations(orig []int) (perms []string) {
	for p := make([]int, len(orig)); p[0] < len(p); nextPerm(p) {
		perm := getPerm(orig, p)
		// array of int -> string ex. [1,2,3] -> 123
		permStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(perm)), ""), "[]")
		perms = append(perms, permStr)
	}

	return
}

func IsPandigital(num int) bool {
	numStr := strconv.Itoa(num)

	digits := strings.Split(strconv.Itoa(num), "")
	sort.Slice(digits, func(i, j int) bool { return digits[i] < digits[j] })

	sorted := strings.Join(digits, "")

	return sorted == "123456789"[:len(numStr)]
}

// เลขนั้นเป็นเลขติดต่อกันหรือไม่ เช่น 123, 6789
func IsNumberConsecutive(num int) bool {
	digits := strings.Split(strconv.Itoa(num), "")
	isConsecutive := true

	for index := range digits {
		if index < len(digits)-1 {
			currentValue, _ := strconv.Atoi(digits[index])
			nextValue, _ := strconv.Atoi(digits[index+1])

			if currentValue+1 != nextValue {
				isConsecutive = false
				break
			}
		}
	}

	return isConsecutive
}

// เลขใน array เรียงต่อกันหรือไม่ เช่น [1234, 1235, 1236, 1237]
func IsArrayConsecutive(digits []int) bool {
	isConsecutive := true

	for index := range digits {
		if index < len(digits)-1 && digits[index]+1 != digits[index+1] {
			isConsecutive = false
			break
		}
	}

	return isConsecutive
}
