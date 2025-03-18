package domain

import (
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewCar(carImage image.Image, carSize Size) Car {
	return Car{
		Image: ebiten.NewImageFromImage(carImage),
		Object: Object{
			Position: Position{
				X: 320,
				Y: GameHeight - 300,
			},
			Size:   carSize,
			Angule: 0,
		},
		Speed:     5,
		SpeedView: 10,
		Fuel: Fuel{
			Percent: 100,
			Time:    time.Now(),
			Color:   Green,
		},
	}
}

func NewGameOver() GameOver {
	return GameOver{
		Flag:      false,
		BoxObject: Object{Position: Position{X: 180, Y: -200}, Size: Size{Width: 450, Height: 120}},
		TextOptions: TextOptions{
			Text:        "Game Over",
			SubText:     "press ''M'' to go to the menu",
			Position:    Position{X: 300 + 105, Y: -200},
			TextSize:    50,
			SubTextSize: 30,
		},
	}
}

func NewMenu() Menu {
	options := map[string]GameState{
		"New game": StateNewGame,
		"Records":  StateRecords,
		"Controls": StateControls,
	}
	actions := make([]Action, len(options))
	i := 0
	for option, state := range options {
		posX := (GameWidth - ButtonWidth) / 2
		posY := (GameHeight+ButtonHeight)/4 + float64(i)*(ButtonHeight+20)
		actions[i] = Action{
			State:   state,
			Visible: true,
			TextOptions: TextOptions{
				Text:     option,
				TextSize: 30,
				Position: Position{
					X: posX + (ButtonWidth / 2),
					Y: posY + (ButtonHeight / 2),
				},
			},
			Object: Object{
				Size:     Size{Width: ButtonWidth, Height: ButtonHeight},
				Position: Position{X: posX, Y: posY},
			},
		}
		i++
	}
	return Menu{Actions: actions}
}
