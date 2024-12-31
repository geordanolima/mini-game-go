package domain

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Car struct {
	Position Position
	Image    *ebiten.Image
	Speed    int
	Fuel     int
	Angule   int
}

type Obstacle struct {
	Position Position
	Image    *ebiten.Image
}

type Game struct {
	car Car
}
