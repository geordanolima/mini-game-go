package main

import (
	_ "image/png"
	"log"
	"mini-game-go/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// create game
	game := engine.NewGame()
	ebiten.SetWindowSize(game.Layout(0, 0))
	ebiten.SetWindowTitle("Car race")
	// run game
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
