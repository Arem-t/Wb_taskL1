package internal

import "fmt"

// reverseString переворачивает строку, включая символы Unicode
func reverseString(s string) string {
	// Преобразуем строку в массив рун
	runes := []rune(s)
	// Получаем длину строки
	length := len(runes)

	// Переворачиваем символы в массиве рун
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Возвращаем перевернутую строку
	return string(runes)
}

func Task19() {
	// Тестовая строка с символами Unicode
	str := "главрыба — абырвалг"

	// Переворачиваем строку
	reversedStr := reverseString(str)

	// Выводим результат
	fmt.Println("Перевернутая строка:", reversedStr)
}
