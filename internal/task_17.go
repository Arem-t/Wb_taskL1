package internal

import (
	"fmt"
	"sort"
)

func Task17() {
	// Исходный отсортированный массив
	arr := []int{2, 5, 7, 11, 13, 17, 19, 23, 29, 31}

	// Элемент, который мы ищем
	target := 17

	// Используем стандартную библиотеку для выполнения бинарного поиска
	index := sort.SearchInts(arr, target)

	// Проверяем, был ли найден элемент
	if index < len(arr) && arr[index] == target {
		fmt.Printf("Элемент %d найден в позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}
