package internal

import (
	"fmt"
	"sync"
)

// Counter - структура счетчика
type Counter struct {
	mu    sync.Mutex // мьютекс для защиты доступа к счетчику
	value int        // значение счетчика
}

// Inc - метод для инкрементирования счетчика на 1
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value - метод для получения текущего значения счетчика
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func Task18() {
	// Создаем экземпляр счетчика
	counter := Counter{}

	// Создаем WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup

	// Количество горутин, которые будут инкрементировать счетчик
	numRoutines := 100

	// Добавляем количество горутин в WaitGroup
	wg.Add(numRoutines)

	// Запускаем горутины для инкрементирования счетчика
	for i := 0; i < numRoutines; i++ {
		go func() {
			defer wg.Done()
			// Инкрементируем счетчик 100 раз
			for j := 0; j < 100; j++ {
				counter.Inc()
			}
		}()
	}

	// Ждем завершения всех горутин
	wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Println("Итоговое значение счетчика:", counter.Value())
}
