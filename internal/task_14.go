package internal

import (
	"fmt"
)

func Task14() {
	checkType(10)
	checkType("hello")
	checkType(true)
	checkType(make(chan int))
}

func checkType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Переменная типа int")
	case string:
		fmt.Println("Переменная типа string")
	case bool:
		fmt.Println("Переменная типа bool")
	case chan int:
		fmt.Println("Переменная типа chan int")
	default:
		fmt.Println("Неизвестный тип")
	}
}
