package internal

import (
	"fmt"
	"strings"
)

// areAllCharactersUnique проверяет, что все символы в строке уникальны
func areAllCharactersUnique(str string) bool {
	// Преобразуем строку в нижний регистр для регистронезависимой проверки
	str = strings.ToLower(str)

	// Создаем map для хранения символов и их количества
	charCount := make(map[rune]int)

	// Проходим по строке и подсчитываем количество каждого символа
	for _, char := range str {
		charCount[char]++
		// Если количество символа больше 1, значит символ не уникален
		if charCount[char] > 1 {
			return false
		}
	}

	// Если все символы уникальны, возвращаем true
	return true
}

func Task26() {
	// Примеры строк для проверки
	stringsToCheck := []string{"abcd", "abCdefAaf", "aabcd"}

	// Проверяем каждую строку и выводим результат
	for _, str := range stringsToCheck {
		result := areAllCharactersUnique(str)
		fmt.Printf("Строка \"%s\": %t\n", str, result)
	}
}
