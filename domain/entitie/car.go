package entitie

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

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
