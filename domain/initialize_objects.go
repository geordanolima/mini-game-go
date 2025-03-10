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
				Y: 700,
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
		BoxObject: Object{Position: Position{X: 180, Y: -100}, Size: Size{Width: 450, Height: 120}},
		TextOptions: TextOptions{
			Text:     "Game Over\n\npress ''N'' to go to the menu",
			Position: Position{X: 300 + 105, Y: -100},
			TextSize: 50,
		},
	}
}

func NewMenu() Menu {
	buttonWidth := float64(200)
	buttonHeight := float64(50)
	options := []string{"New game", "Records", "Controls"}
	states := []GameState{StateNewGame, StateRecords, StateControls}
	actions := make([]Action, len(options))
	for i := 0; i < len(options); i++ {
		posX := (GameWidth - buttonWidth) / 2
		posY := (GameHeight+buttonHeight)/4 + float64(i)*(buttonHeight+20)
		actions[i] = Action{
			State: states[i],
			TextOptions: TextOptions{
				Text:     options[i],
				TextSize: 30,
				Position: Position{
					X: posX + (buttonWidth / 2),
					Y: posY + (buttonHeight / 2),
				},
			},
			Object: Object{
				Size:     Size{Width: buttonWidth, Height: buttonHeight},
				Position: Position{X: posX, Y: posY},
			},
		}
	}
	return Menu{Actions: actions}
}
