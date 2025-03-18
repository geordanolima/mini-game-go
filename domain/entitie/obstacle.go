package entitie

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type TypeObtacle int

const (
	ObjectObstacle TypeObtacle = iota
	ObjectGasoline
)

type Obstacle struct {
	TextValue string
	Value     int
	Object    Object
	Image     *ebiten.Image
	FilePath  string
	Type      TypeObtacle
}
