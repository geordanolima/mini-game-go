package engine

import (
	"mini-game-go/domain"
	"mini-game-go/domain/entitie"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (game *Game) drawDifficulty(screen *ebiten.Image) {
	for _, action := range game.DifficultySelector.Actions {
		if action.Visible {
			if action.Active {
				drawButtonMarginColor(
					action.TextOptions.Text,
					action.Object.Position,
					action.Object.Size,
					action.TextOptions.TextSize,
					game.Font,
					screen,
					domain.White,
					domain.Gray,
					domain.Black,
					text.AlignStart,
					text.AlignCenter,
				)
				drawButtonMarginColor(
					action.TextOptions.SubText,
					entitie.Position{X: 50, Y: (domain.GameHeight - (domain.GameHeight/4)*2)},
					entitie.Size{Width: domain.GameWidth - 100, Height: domain.GameHeight / 4},
					25,
					game.Font,
					screen,
					domain.White,
					domain.Black,
					domain.Black,
					text.AlignStart,
					text.AlignStart,
				)
			} else {
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
	backButon(game, screen)
	confirmButon(game, screen)
}
