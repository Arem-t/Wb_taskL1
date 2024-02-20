package internal

import (
	"fmt"
	"sync"
)

func Task2() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			square := x * x
			fmt.Printf("Квадрат числа %d: %d\n", x, square)
		}(num)
	}

	wg.Wait()
}
