package engine

import (
	"fmt"
	"image/color"
	"mini-game-go/domain"
	"mini-game-go/domain/entitie"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (game *Game) drawEnterNameScreen(screen *ebiten.Image) {
	title(fmt.Sprintf("Your score are %d", game.Score.Score), "", game.Font, screen)

	posX := ((domain.GameWidth - domain.ButtonWidth) / 5) - 25
	vector.DrawFilledRect(screen, float32(posX), 200, float32(domain.GameWidth-(2*posX))+5, 55, color.Gray{}, false)
	vector.DrawFilledRect(screen, float32(posX), 200, float32(domain.GameWidth-(2*posX)), 50, color.White, false)

	LoadText(
		"Type your name and press Enter to save your record:",
		entitie.Position{X: posX + 10, Y: 202},
		20,
		game.Font,
		screen,
		domain.Black,
		text.AlignStart,
		text.AlignStart,
	)

	LoadText(
		game.User.Name,
		entitie.Position{X: posX + 10, Y: 235},
		25,
		game.Font,
		screen,
		domain.Black,
		text.AlignStart,
		text.AlignCenter,
	)

}
