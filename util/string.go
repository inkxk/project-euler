package util

func PossibleRotation(initial string) []string {
	rotated := make([]string, 0, len(initial))

	for i := 0; i < len(initial); i++ {
		// ตัดตัวหน้าไปใส่ด้านหลัง
		rotatedStr := initial[i:] + initial[:i]
		rotated = append(rotated, rotatedStr)
	}

	return rotated
}
