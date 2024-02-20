package internal

import (
	"fmt"
)

// Функция для создания множества из последовательности строк
func createSet(strings []string) map[string]bool {
	set := make(map[string]bool)

	for _, str := range strings {
		set[str] = true
	}

	return set
}

func Task12() {
	// Заданная последовательность строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество
	set := createSet(strings)

	// Выводим элементы множества
	fmt.Println("Множество:")
	for key := range set {
		fmt.Println(key)
	}
}
