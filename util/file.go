package util

import (
	"fmt"
	"os"
	"strings"
)

func LoadFile(fileName string) []string {
	b, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}

	return strings.Split(string(b), `,`)
}
