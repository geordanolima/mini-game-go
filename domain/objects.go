package domain

import (
	_ "image/png"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Position struct {
	X int
	Y int
}

type Size struct {
	Width  int
	Height int
}

type Object struct {
	Size     Size
	Position Position
	Angule   int
	Margin   int
}

type Car struct {
	Object    Object
	Image     *ebiten.Image
	Speed     int
	SpeedView int
	Fuel      Fuel
	Angule    int
}

type Fuel struct {
	Percent int
	Time    time.Time
	Color   string
}

type Score struct {
	Score int
	Time  time.Time
}

type Obstacle struct {
	Object   Object
	Image    *ebiten.Image
	FilePath string
}

type Gasoline struct {
	Percent  int
	Object   Object
	Image    *ebiten.Image
	FilePath string
}

type Game struct {
	car Car
}

type GameOver struct {
	Flag         bool
	BoxObject    Object
	TextPosition Position
	Text         string
}
