package internal

import "fmt"

func Task22() {
	// Задаем значения переменных a и b
	a := 2 << 20 // a = 2^21
	b := 2 << 21 // b = 2^22

	// Умножение
	multiplication := a * b
	fmt.Printf("Умножение: %d * %d = %d\n", a, b, multiplication)

	// Деление
	division := b / a
	fmt.Printf("Деление: %d / %d = %d\n", b, a, division)

	// Сложение
	addition := a + b
	fmt.Printf("Сложение: %d + %d = %d\n", a, b, addition)

	// Вычитание
	subtraction := b - a
	fmt.Printf("Вычитание: %d - %d = %d\n", b, a, subtraction)
}
