package internal

import (
	"fmt"
	"sync"
)

func Task3() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	sum := 0
	for _, num := range numbers {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			square := x * x
			mu.Lock()
			sum = sum + square
			mu.Unlock()
		}(num)
	}
	wg.Wait()
	fmt.Printf("Сумма квадратов: %d\n", sum)

}
