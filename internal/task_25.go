package internal

import (
	"fmt"
	"time"
)

// sleep функция задержки, которая приостанавливает выполнение программы на заданное количество секунд
func sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func Task25() {
	fmt.Println("Начало программы")

	// Вызываем функцию sleep для задержки на 3 секунды
	sleep(3)

	fmt.Println("Прошло 3 секунды")
}
