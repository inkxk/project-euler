package util

import (
	"sort"
	"strings"
)

// ค่าที่ส่งมา สามารถสลับตำแหน่งเป็นค่าอะไรได้บ้าง
// จำนวนทั้งหมดใน array มีค่าเท่ากับ หลักของค่าที่ส่งมา
// 123 => [123, 231, 312], สามารถเป็นได้ 3 ตัวตามหลักของค่าที่ส่งมา
// ไม่สามารถสลับเป็น 321 ได้ อิงจากค่าที่ส่งมา 3 ต้องอยู่ด้านขวาของ 2 ไม่ใช่ซ้าย
// คนละเรื่องกับความน่าจะเป็น ที่สามารถมีได้ n! = 3! = 6 ตัว
// อันนี้แค่ rotate position
func PossibleRotation(initial string) []string {
	rotated := make([]string, 0, len(initial))

	for i := 0; i < len(initial); i++ {
		// ตัดตัวหน้าไปใส่ด้านหลัง
		rotatedStr := initial[i:] + initial[:i]
		rotated = append(rotated, rotatedStr)
	}

	return rotated
}

// เช็คว่า original กับ reverse มีค่าเท่ากันไหม
func IsPalindrome(original string) bool {
	reversed := ""
	for _, c := range original {
		reversed = string(c) + reversed
	}
	return original == reversed
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
