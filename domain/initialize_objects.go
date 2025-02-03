package domain

import (
	"mini-game-go/helpers"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewCar() Car {
	carSize := Size{Height: 290, Width: 120}
	carImage, _ := helpers.LoadImageResize("car.png", carSize.Width, carSize.Height)
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
		Speed: 5,
		Fuel: Fuel{
			Percent: 100,
			Time:    time.Now(),
			Color:   ColorGreen,
		},
	}
}

func NewGameOver() GameOver {
	return GameOver{
		Flag:         false,
		BoxObject:    Object{Position: Position{X: 280, Y: -100}, Size: Size{Width: 210, Height: 60}},
		TextPosition: Position{X: 300, Y: -100},
		Text:         "Game Over",
	}
}
