package main

import (
	"image"
	_ "image/png"
	"log"
	domain "mini-game-go/domain"
	usecase "mini-game-go/use_case"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

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

	game := usecase.NewGame(
		domain.Car{
			Image: ebiten.NewImageFromImage(img),
			Position: domain.Position{
				X: 350,
				Y: 700,
			},
			Speed:  1,
			Fuel:   100,
			Angule: 0,
		},
	)
	ebiten.SetWindowSize(game.Layout(0, 0))
	ebiten.SetWindowTitle("Car race")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
