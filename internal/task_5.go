package internal

import (
	"fmt"
	"time"
)

func Task5() {
	var sec = 100 * time.Second
	fmt.Println("Введите время в секундах:")
	_, err := fmt.Scanf("%d", &sec)
	if err != nil {
		return
	}

	ch := make(chan int)
	done := make(chan struct{})

	// Отправка значений в канал
	go func() {
		for i := 1; ; i++ {
			ch <- i
			time.Sleep(500 * time.Millisecond) // Задержка между отправками значений
		}
	}()

	// Чтение значений из канала
	go func() {
		for {
			select {
			case val := <-ch:
				fmt.Println("Принято:", val)
			case <-done:
				return
			}
		}
	}()

	// Завершение программы через N секунд
	duration := sec * time.Second
	fmt.Printf("Время работы: %s\n", duration)
	time.Sleep(duration)
	close(done)
}
