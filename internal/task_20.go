package internal

import (
	"fmt"
	"strings"
)

// reverseWordsOrder переворачивает порядок слов в строке
func reverseWordsOrder(s string) string {
	// Разбиваем строку на слова
	words := strings.Fields(s)

	// Переворачиваем порядок слов в массиве
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Собираем перевернутую строку из слов
	reversedStr := strings.Join(words, " ")

	return reversedStr
}

func Task20() {
	// Исходная строка
	var str string
	println("Введите строку:")
	_, err := fmt.Scanf("%d", &str)
	if err != nil {
		return
	}

	// Переворачиваем порядок слов в строке
	reversedStr := reverseWordsOrder(str)

	// Выводим результат
	fmt.Println("Перевернутая строка:", reversedStr)
}
