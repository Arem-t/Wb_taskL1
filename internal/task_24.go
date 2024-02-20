package internal

import (
	"fmt"
	"math"
)

// Point представляет точку в двумерном пространстве
type Point struct {
	x float64
	y float64
}

// NewPoint создает новую точку с заданными координатами
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// DistanceTo вычисляет расстояние от текущей точки до указанной точки
func (p *Point) DistanceTo(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func Task24() {
	// Создаем две точки
	point1 := NewPoint(1, 2)
	point2 := NewPoint(4, 6)

	// Вычисляем расстояние между точками
	distance := point1.DistanceTo(point2)

	// Выводим результат
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
