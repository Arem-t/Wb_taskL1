package internal

import "fmt"

func Task23() {
	// Исходный слайс
	slice := []int{1, 2, 3, 4, 5}

	// Индекс элемента, который нужно удалить
	index := 2

	// Удаление элемента из слайса
	if index >= 0 && index < len(slice) {
		// Перемещаем элементы после удаляемого элемента влево
		copy(slice[index:], slice[index+1:])
		// Обрезаем слайс, удаляя последний элемент
		slice = slice[:len(slice)-1]
		fmt.Println("Элемент удален из слайса:", slice)
	} else {
		fmt.Println("Недопустимый индекс для удаления элемента")
	}
}
