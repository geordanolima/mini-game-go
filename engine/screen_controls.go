package engine

import (
	"mini-game-go/domain"
	"mini-game-go/domain/entitie"

	"github.com/hajimehoshi/ebiten/v2"
)

func (game *Game) drawControls(screen *ebiten.Image) {
	title("Controls", "", game.Font, screen)
	posX := ((domain.GameWidth - domain.ButtonWidth) / 2) - 25
	i := 0
	for _, action := range domain.Controls {
		posY := (domain.GameHeight+domain.ButtonHeight)/5 + float64(i)*(domain.ButtonHeight+20)
		drawButton(action, entitie.Position{X: posX, Y: posY}, game.Font, screen, 30)
		i++
	}
	backButon(game, screen)
}
