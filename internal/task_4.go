package internal

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, ch <-chan int, wg *sync.WaitGroup, done <-chan struct{}) {
	for {
		select {
		case data, ok := <-ch:
			if !ok {
				return // Канал закрыт, завершаем работу
			}
			fmt.Printf("Воркер %d принял: %d\n", id, data)
		case <-done:
			fmt.Printf("Воркер %d устал работать\n", id)
			return // Получен сигнал о завершении, завершаем работу
		}

	}
}

func Task4() {
	fmt.Println("Введите количество воркеров:")
	var isdone = true
	var numWorkers int
	_, err := fmt.Scanf("%d", &numWorkers)
	if err != nil {
		return
	}

	// Создаем канал и группу WaitGroup
	dataChannel := make(chan int)
	done := make(chan struct{})
	var wg sync.WaitGroup

	// Создаем канал для обработки сигнала Ctrl+C
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	// Создаем N воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChannel, &wg, done)
	}

	go func() {
		<-sigCh
		// Закрываем канал и ожидаем завершения всех воркеров
		isdone = false
		close(done)
		close(dataChannel)

	}()

	// Постоянная запись данных в канал
	for i := 1; ; i++ {
		if isdone == true {
			dataChannel <- i
			time.Sleep(time.Second) // Имитация задержки
		} else {
			return
		}
	}

	wg.Wait()

}
