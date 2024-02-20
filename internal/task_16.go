package internal

import (
	"fmt"
)

func Task16() {
	// Наш исходный массив для сортировки
	arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}

	// Вызываем функцию быстрой сортировки
	quickSort(arr)

	// Выводим отсортированный массив
	fmt.Println("Отсортированный массив:", arr)
}

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// Выбираем опорный элемент
	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	// Разбиваем массив на две части вокруг опорного элемента
	left := 0
	right := len(arr) - 1

	for left <= right {
		for arr[left] < pivot {
			left++
		}

		for arr[right] > pivot {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	// Рекурсивно сортируем подмассивы
	quickSort(arr[:right+1])
	quickSort(arr[left:])
}
