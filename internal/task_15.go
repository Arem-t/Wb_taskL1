package internal

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = makeCopy(v[:100]) // Создаем копию подстроки и сохраняем ее в justString
}

func Task15() {
	someFunc()
}

// createHugeString возвращает большую строку
func createHugeString(size int) string {
	// Ваша реализация createHugeString
	return "huge_string_content"
}

// makeCopy создает копию строки
func makeCopy(s string) string {
	// Создаем новую строку той же длины, что и s
	copyStr := make([]byte, len(s))
	// Копируем содержимое строки s в новую строку copyStr
	copy(copyStr, s)
	// Возвращаем копию строки в виде string
	return string(copyStr)
}
