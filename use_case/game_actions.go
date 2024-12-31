package usecase

import (
	"image/color"
	_ "image/png"
	domain "mini-game-go/domain"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	car domain.Car
}

func NewGame(car domain.Car) Game {
	return Game{car: car}
}

func (g *Game) Update() error {
	// validar posição do carro
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.car.Position.X -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.car.Position.X += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.car.Position.Y -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.car.Position.Y += 5
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//  criar tela
	screen.Fill(color.RGBA{0xee, 0xee, 0xee, 0xff})

	// criar carro
	carImg := &ebiten.DrawImageOptions{}
	carImg.GeoM.Scale(0.25, 0.25)
	carImg.GeoM.Translate(float64(g.car.Position.X), float64(g.car.Position.Y))
	screen.DrawImage(g.car.Image, carImg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 1000
}
