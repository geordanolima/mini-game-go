package main

import (
	"image"
	"image/color"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Coordinate struct {
	X float64
	Y float64
}

type Car struct {
	image    *ebiten.Image
	Position Coordinate
	speed    int
	fuel     int
}

type Obstacle struct {
	image    *ebiten.Image
	Position Coordinate
	size     int
}

type Game struct {
	imageCar *ebiten.Image
	x        int
	y        int
}

func (g *Game) Update() error {
	// validar posição do carro
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.x -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.x += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.y -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.y += 5
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//  criar tela
	screen.Fill(color.RGBA{0xee, 0xee, 0xee, 0xff})

	// criar carro
	carImg := &ebiten.DrawImageOptions{}
	carImg.GeoM.Translate(float64(g.x), float64(g.y))
	screen.DrawImage(g.imageCar, carImg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 900, 600
}

func main() {
	// get car image
	file, err := os.Open("assets/car.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// decodificar imagem
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	imageCar := ebiten.NewImageFromImage(img)

	game := &Game{
		imageCar: imageCar,
		x:        100,
		y:        500,
	}
	ebiten.SetWindowSize(game.Layout(0, 0))
	ebiten.SetWindowTitle("Car race")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
