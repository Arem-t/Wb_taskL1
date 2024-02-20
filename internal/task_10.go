package internal

import (
	"fmt"
	"math"
	"sort"
)

func groupTemperatures(temperatures []float64) map[int][]float64 {
	groupedTemperatures := make(map[int][]float64)

	for _, temp := range temperatures {
		// Округляем температуру до ближайшего кратного 10
		key := int(math.Round(temp/10.0) * 10)

		// Добавляем температуру в соответствующую группу
		groupedTemperatures[key] = append(groupedTemperatures[key], temp)
	}

	return groupedTemperatures
}

func Task10() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Группируем температуры
	groupedTemperatures := groupTemperatures(temperatures)

	// Получаем отсортированные ключи (значения температур)
	var keys []int
	for key := range groupedTemperatures {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	// Выводим группы температур
	for _, key := range keys {
		fmt.Printf("%d градусов: %v\n", key, groupedTemperatures[key])
	}
}
