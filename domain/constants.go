package domain

import (
	"mini-game-go/domain/entitie"
)

// Sizes
var GameWidth float64 = 750
var GameHeight float64 = 900
var LineWidth float64 = 10
var LineHeight float64 = 200
var ButtonWidth float64 = 200
var ButtonHeight float64 = 50

// colores
var Black string = "#000000"
var DarkGray string = "#0F0F0F"
var Gray string = "#3C3C3C"
var Green string = "#07B51E"
var Yellow string = "B57E07"
var Red string = "#B50707"
var White string = "#FFFFFF"

// positions to randomize objects
var PositionsX = []float64{15, 155, 305, 455, 605}

// fuel percent by difficulty
var PercentsGas = []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
)

var PercentGasByLevels = map[Difficulty][]int{
	Easy:   PercentsGas[4:],  // 70, 80, 90, 100
	Medium: PercentsGas[2:6], // 30, 40, 50, 60
	Hard:   PercentsGas[:4],  // 10, 20, 30, 40
}

// obstacles
var ObstacleImages []entitie.Obstacle = []entitie.Obstacle{
	{FilePath: "cone.png", Object: entitie.Object{Size: entitie.Size{Width: 100, Height: 100}, Position: entitie.Position{X: 20}, Margin: -20}, Type: entitie.ObjectObstacle},
	{FilePath: "cone2.png", Object: entitie.Object{Size: entitie.Size{Width: 125, Height: 125}, Position: entitie.Position{X: 10}, Margin: -35}, Type: entitie.ObjectObstacle},
	{FilePath: "hole.png", Object: entitie.Object{Size: entitie.Size{Width: 100, Height: 100}, Position: entitie.Position{X: 25}, Margin: -30}, Type: entitie.ObjectObstacle},
	{FilePath: "truck.png", Object: entitie.Object{Size: entitie.Size{Width: 160, Height: 330}, Position: entitie.Position{X: 5}, Margin: -10}, Type: entitie.ObjectObstacle},
	{FilePath: "bus.png", Object: entitie.Object{Size: entitie.Size{Width: 150, Height: 430}, Position: entitie.Position{X: 0}, Margin: -10}, Type: entitie.ObjectObstacle},
}

var Controls = []string{
	"↑/w   move up",
	"↓/s   move down",
	"←/a   move left",
	"→/d   move right",
	"m     menu",
}
