package domain

import (
	_ "image/png"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Position struct {
	X float64
	Y float64
}

type Size struct {
	Width  float64
	Height float64
}

type Object struct {
	Size     Size
	Position Position
	Angule   int
	Margin   float64
}

type Car struct {
	Object    Object
	Image     *ebiten.Image
	Speed     float64
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
	Flag        bool
	BoxObject   Object
	TextOptions TextOptions
}

type TextOptions struct {
	Text        string
	SubText     string
	TextSize    float64
	SubTextSize float64
	Position    Position
}

type Action struct {
	TextOptions TextOptions
	Object      Object
	Visible     bool
	State       GameState
}

type Menu struct {
	Actions []Action
}
