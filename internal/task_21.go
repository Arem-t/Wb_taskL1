package internal

import "fmt"

// Football имитирует футбольный мяч
type Football struct{}

// Kick выполняет удар по футбольному мячу
func (f *Football) Kick() {
	fmt.Println("Kick the football!")
}

// BasketballBall интерфейс для мяча в баскетболе
type BasketballBall interface {
	Bounce()
	Shoot()
}

// BasketballBallAdapter адаптирует футбольный мяч к интерфейсу мяча в баскетболе
type BasketballBallAdapter struct {
	football *Football
}

// Bounce эмулирует отскок мяча в баскетболе
func (b *BasketballBallAdapter) Bounce() {
	fmt.Println("Bounce the basketball!")
}

// Shoot эмулирует бросок мяча в баскетболе
func (b *BasketballBallAdapter) Shoot() {
	// Превращаем удар по футбольному мячу в бросок мяча в баскетболе
	b.football.Kick()
	fmt.Println("Shoot the basketball!")
}

func Task21() {
	// Создаем футбольный мяч
	football := &Football{}

	// Создаем адаптер для футбольного мяча
	basketballAdapter := &BasketballBallAdapter{football}

	// Играем в баскетбол с использованием футбольного мяча через адаптер
	basketballAdapter.Bounce()
	basketballAdapter.Shoot()
}
