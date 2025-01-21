package domain

import (
	_ "image/png"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Car struct {
	Position Position
	Image    *ebiten.Image
	Speed    int
	Fuel     Fuel
	Angule   int
}

type Fuel struct {
	Percent int
	Time    time.Time
	Color   string
}

type Obstacle struct {
	Position Position
	Image    *ebiten.Image
}

type Game struct {
	car Car
}

type GameOver struct {
	Flag         bool
	BoxPosition  Position
	TextPosition Position
}
