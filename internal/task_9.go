package internal

import (
	"fmt"
)

// Функция, которая принимает входной канал, читает числа из него,
// умножает каждое число на 2 и отправляет результат в выходной канал.
func multiplyByTwo(input <-chan int, output chan<- int) {
	for x := range input {
		result := x * 2
		output <- result
	}
	close(output)
}

func Task9() {
	// Создаем каналы для передачи данных
	input := make(chan int)
	output := make(chan int)

	// Запускаем горутину для выполнения функции multiplyByTwo
	go multiplyByTwo(input, output)

	// Заполняем входной канал числами из массива
	numbers := []int{1, 2, 3, 4, 5}
	go func() {
		for _, num := range numbers {
			input <- num
		}
		close(input)
	}()

	// Читаем результаты из выходного канала и выводим их в stdout
	for doubled := range output {
		fmt.Println(doubled)
	}
}
