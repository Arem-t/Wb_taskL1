package internal

import (
	"fmt"
)

// Устанавливает i-й бит в значении num в значение bit (0 или 1)
func setBit(num int64, i int, bit int) int64 {
	// Создаем маску с единицей в позиции i и инвертируем ее, чтобы получить 0 везде, кроме позиции i
	mask := ^(int64(1) << i)

	// Очищаем i-й бит в num
	num &= mask

	// Устанавливаем i-й бит в значение bit
	num |= int64(bit) << i

	return num
}

func Task8() {
	var num int64
	var i, bit int

	// Вводим переменную и бит, который хотим установить (0 или 1)
	fmt.Print("Введите целое число (int64): ")
	fmt.Scan(&num)

	fmt.Print("Введите позицию бита: ")
	fmt.Scan(&i)

	fmt.Print("Введите значение бита (0 или 1): ")
	fmt.Scan(&bit)

	// Проверяем корректность введенной позиции бита
	if i < 0 || i > 63 {
		fmt.Println("Позиция бита должна быть от 0 до 63.")
		return
	}

	// Проверяем корректность введенного значения бита
	if bit != 0 && bit != 1 {
		fmt.Println("Значение бита должно быть 0 или 1.")
		return
	}

	// Устанавливаем i-й бит в значение bit
	newNum := setBit(num, i, bit)

	fmt.Printf("Новое значение числа после установки бита: %d\n", newNum)
}
