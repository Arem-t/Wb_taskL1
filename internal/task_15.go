package internal

import (
	"fmt"
	"strings"
)

func createHugeString(size int) string {
	var result strings.Builder
	for i := 0; i < size; i++ {
		result.WriteRune(' ') // Добавляем пробел в строку
	}
	return result.String()
}

var justString string

func someFunc() {
	v := createHugeString(1000)
	justString = v[:100]
}

func Task15() {
	someFunc()
	fmt.Println(justString)
}
