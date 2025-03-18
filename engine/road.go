package engine

import (
	"mini-game-go/domain"
	"mini-game-go/domain/entitie"

	"github.com/hajimehoshi/ebiten/v2"
)

func loadRoad() []entitie.Object {
	// create start positions to dashed lines by road
	lines := make([]entitie.Object, 21)
	positionsX := []float64{150, 300, 450, 600}
	index := 0
	for line := 0; line <= 4; line++ {
		for column := 0; column <= 3; column++ {
			lines[index] = entitie.Object{
				Position: entitie.Position{X: positionsX[column], Y: float64(int(domain.LineHeight+50) * line)},
				Size:     entitie.Size{Height: domain.LineHeight, Width: 0},
			}
			index++
		}
	}
	return lines
}

func DrawRoad(screen *ebiten.Image, game *Game) {
	// draw continuous line on the sides
	DrawRect64(0, 0, domain.LineWidth, domain.GameHeight, screen, domain.White)
	DrawRect64(domain.GameWidth-domain.LineWidth, 0, domain.LineWidth, domain.GameHeight, screen, domain.White)
	// draw dashed lines
	for i := 0; i < len(game.Road); i++ {
		posY := game.Road[i].Position.Y + game.Car.Speed
		DrawRect64(game.Road[i].Position.X, posY, domain.LineWidth, domain.LineHeight, screen, domain.White)
	}
}
