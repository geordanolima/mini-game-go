package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (game *Game) drawMenu(screen *ebiten.Image) {
	for _, action := range game.Menu.Actions {
		if action.Visible {
			drawButtonMargin(
				action.TextOptions.Text,
				action.Object.Position,
				action.Object.Size,
				action.TextOptions.TextSize,
				game.Font,
				screen,
			)
		}
	}
}
