package domain

import (
	"image"
	"mini-game-go/domain/entitie"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewCar(carImage image.Image, carSize entitie.Size) entitie.Car {
	return entitie.Car{
		Image: ebiten.NewImageFromImage(carImage),
		Object: entitie.Object{
			Position: entitie.Position{
				X: 320,
				Y: GameHeight - 300,
			},
			Size:   carSize,
			Angule: 0,
		},
		Speed:     5,
		SpeedView: 10,
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

func NewMenu() entitie.Menu {
	options := map[string]entitie.GameState{
		"New game": entitie.StateNewGame,
		"Records":  entitie.StateRecords,
		"Controls": entitie.StateControls,
	}
	actions := make([]entitie.Action, len(options))
	i := 0
	for option, state := range options {
		posX := (GameWidth - ButtonWidth) / 2
		posY := (GameHeight+ButtonHeight)/4 + float64(i)*(ButtonHeight+20)
		actions[i] = entitie.Action{
			State:   state,
			Visible: true,
			TextOptions: entitie.TextOptions{
				Text:     option,
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
	return entitie.Menu{Actions: actions}
}
