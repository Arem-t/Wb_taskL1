package internal

import "fmt"

func Task13() {
	a, b := 5, 7

	fmt.Println("До обмена:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// Обмен значениями без временной переменной
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println("После обмена:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
