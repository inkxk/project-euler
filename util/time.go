package util

import (
	"fmt"
	"time"
)

// หาเวลาทั้งหมดที่ใช้ในการ process
func TimeProcess(start time.Time) {
	elapsed := time.Since(start)
	// fmt.Printf("total time process: %v secs\n", elapsed.Seconds())
	fmt.Printf("total time process: %s\n", elapsed)
}
