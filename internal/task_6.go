package internal

import (
	"fmt"
	"sync"
	"time"
)

func worker1(wg *sync.WaitGroup, quit chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-quit:
			fmt.Println("Горутина 1 получила сигнал выхода")
			return
		default:
			fmt.Println("Горутина 1 работает")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func worker2(wg *sync.WaitGroup, stop chan bool) {
	defer wg.Done()
	for {
		select {
		case <-stop:
			fmt.Println("Горутина 2 получила сигнал остановки")
			return
		default:
			fmt.Println("Горутина 2 работает")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func Task6() {
	var wg sync.WaitGroup

	// Пример с использованием канала с типом struct{}
	quit := make(chan struct{})
	wg.Add(1)
	go func() {
		defer close(quit)
		defer fmt.Println("Goroutine 1 stopped")
		worker1(&wg, quit)
	}()

	// Пример с использованием канала с типом bool
	stop := make(chan bool)
	wg.Add(1)
	go func() {
		defer fmt.Println("Goroutine 2 stopped")
		defer wg.Done()
		worker2(&wg, stop)
	}()

	// Пример с использованием ожидания группы горутин с таймаутом
	time.Sleep(3 * time.Second) // Подождем некоторое время
	fmt.Println("Главная горутина ждет окончание остальных...")
	wg.Wait() // Ожидание завершения всех горутин
	fmt.Println("Все горутины остановлены")
}
