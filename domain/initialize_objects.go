package domain

import (
	"image"
	"mini-game-go/domain/entitie"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func getInitialSpeedByDifficulty(difficulty Difficulty) float64 {
	switch difficulty {
	case Hard:
		return 10
	case Medium:
		return 5
	default:
		return 3
	}
}

func NewCar(carImage image.Image, carSize entitie.Size, difficulty Difficulty) entitie.Car {
	return entitie.Car{
		Image: ebiten.NewImageFromImage(carImage),
		Object: entitie.Object{
			Position: entitie.Position{
				X: 320,
				Y: GameHeight - 300,
			},
			Size: carSize,
		},
		Speed:     getInitialSpeedByDifficulty(difficulty),
		SpeedView: int(getInitialSpeedByDifficulty(difficulty)),
		Fuel: entitie.Fuel{
			Percent: 100,
			Time:    time.Now(),
			Color:   Green,
		},
	}
}

func NewGameOver() entitie.GameOver {
	return entitie.GameOver{
		Flag:      false,
		BoxObject: entitie.Object{Position: entitie.Position{X: 180, Y: -200}, Size: entitie.Size{Width: 450, Height: 120}},
		TextOptions: entitie.TextOptions{
			Text:        "Game Over",
			SubText:     "press ''M'' to go to the menu",
			Position:    entitie.Position{X: 300 + 105, Y: -200},
			TextSize:    50,
			SubTextSize: 30,
		},
	}
}

func createSelectors(options map[string][]any) entitie.Selector {
	actions := make([]entitie.Action, len(options))
	i := 0
	for key, value := range options {
		posX := (GameWidth - ButtonWidth) / 2
		posY := (GameHeight+ButtonHeight)/4 + float64(i)*(ButtonHeight+20)
		actions[i] = entitie.Action{
			State:   value[0].(int),
			Visible: true,
			TextOptions: entitie.TextOptions{
				Text:     key,
				SubText:  value[1].(string),
				TextSize: 30,
				Position: entitie.Position{
					X: posX + (ButtonWidth / 2),
					Y: posY + (ButtonHeight / 2),
				},
			},
			Object: entitie.Object{
				Size:     entitie.Size{Width: ButtonWidth, Height: ButtonHeight},
				Position: entitie.Position{X: posX, Y: posY},
			},
		}
		i++
	}
	return entitie.Selector{Actions: actions}
}

func NewMenu() entitie.Selector {
	options := map[string][]any{
		"New game": {int(entitie.StateDifficulty), ""},
		"Records":  {int(entitie.StateRecords), ""},
		"Controls": {int(entitie.StateControls), ""},
	}
	return createSelectors(options)
}

func NewDifficultys() entitie.Selector {
	options := map[string][]any{
		"Easy":   {int(Easy), "In this difficulty, initial speed 3, only 2 obstacles at a time, \nand more fuel on the way. \nRefuel your car with 70% to 100% fuel."},
		"Medium": {int(Medium), "In this difficulty, initial speed 5, 3 obstacles at a time, \nand more fuel on the way. \nRefuel your car with 30% to 60% fuel."},
		"Hard":   {int(Hard), "In this difficulty, initial speed 10, 4 obstacles at a time, \nand more fuel on the way. \nRefuel your car with 10% to 40% fuel."},
	}
	return createSelectors(options)
}
