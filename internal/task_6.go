package internal

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Task6() {
	// Используем канал для сигнализации о завершении
	done := make(chan bool)
	go func() {
		defer close(done)
		fmt.Println("Горутина 1 начала выполнение")
		time.Sleep(2 * time.Second)
		fmt.Println("Горутина 1 завершила выполнение")
	}()

	// Используем контекст для отмены выполнения
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		fmt.Println("Горутина 2 начала выполнение")
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("Горутина 2 завершила выполнение")
		case <-ctx.Done():
			fmt.Println("Горутина 2 была остановлена")
		}
	}()

	// Используем мьютекс и переменную состояния для управления выполнением
	var wg sync.WaitGroup
	stop := false
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Горутина 3 начала выполнение")
		for {
			if stop {
				break
			}
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Горутина 3 завершила выполнение")
	}()

	// Ждем некоторое время, затем останавливаем горутины
	time.Sleep(2 * time.Second)
	fmt.Println("Останавливаем горутины...")

	// Остановка первой горутины через канал
	<-done

	// Отмена выполнения второй горутины через контекст
	cancel()

	// Остановка третьей горутины через мьютекс и переменную состояния
	stop = true
	wg.Wait()

	fmt.Println("Основная горутина завершила выполнение")
}
