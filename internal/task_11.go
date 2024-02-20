package internal

import (
	"fmt"
)

// Функция для нахождения пересечения двух множеств
func intersection(set1, set2 []int) []int {
	// Создаем хэш-таблицы для обоих множеств
	hashMap1 := make(map[int]bool)
	hashMap2 := make(map[int]bool)

	// Заполняем хэш-таблицы элементами из соответствующих множеств
	for _, v := range set1 {
		hashMap1[v] = true
	}
	for _, v := range set2 {
		hashMap2[v] = true
	}

	// Создаем слайс для хранения пересечения
	var intersect []int

	// Перебираем элементы первого множества
	// Если элемент присутствует и во втором множестве, добавляем его в пересечение
	for key := range hashMap1 {
		if hashMap2[key] {
			intersect = append(intersect, key)
		}
	}

	return intersect
}

func Task11() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{4, 5, 6, 7, 8}

	// Находим пересечение множеств
	intersect := intersection(set1, set2)

	// Выводим результат
	fmt.Println("Пересечение множеств:", intersect)
}
