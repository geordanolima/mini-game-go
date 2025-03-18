package engine

import (
	"fmt"
	"mini-game-go/domain"
	"mini-game-go/domain/entitie"

	"github.com/hajimehoshi/ebiten/v2"
)

func (game *Game) drawRecords(screen *ebiten.Image) {
	title("Records", "When          | Score           |  Name", game.Font, screen)
	posX := ((domain.GameWidth - domain.ButtonWidth) / 6) - 25
	for i, record := range listRecords() {
		posY := (domain.GameHeight+domain.ButtonHeight)/8 + float64(i)*(domain.ButtonHeight+20)
		drawButton(
			fmt.Sprintf("%s | %d |  %s", record.CreatedAt.Format("06-01-02 15:04:05"), record.Score, record.Name),
			entitie.Position{X: posX, Y: posY},
			game.Font,
			screen,
			25,
		)
	}
	drawButtonMargin(
		"Back",
		entitie.Position{X: 20, Y: domain.GameHeight - domain.ButtonHeight - 20},
		entitie.Size{Width: domain.ButtonWidth, Height: domain.ButtonHeight},
		30,
		game.Font,
		screen,
	)
}
