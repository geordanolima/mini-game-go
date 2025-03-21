package engine

import (
	"fmt"
	"log"
	"mini-game-go/domain"
	"mini-game-go/domain/entitie"
	"mini-game-go/helpers"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

func getTextCaracters(text, char string, amount int) string {
	if len(text) > amount {
		return text[:amount]
	}
	return text + strings.Repeat(char, amount-len(text))
}

func (game *Game) drawRecords(screen *ebiten.Image) {
	font, err := helpers.LoadFont("RobotoMono.ttf")
	if err != nil {
		log.Fatal(err)
	}
	titleSize("Records", " Diff. | Score | Name         | When              ", font, screen, 25, 20)
	posX := ((domain.GameWidth - domain.ButtonWidth) / 6) - 25
	for i, record := range listRecords() {
		posY := (domain.GameHeight+domain.ButtonHeight)/8 + float64(i)*(domain.ButtonHeight+20)
		var difficulty string
		switch record.Difficulty {
		case domain.Hard:
			difficulty = "Hard"
		case domain.Medium:
			difficulty = "Medium"
		default:
			difficulty = "Easy"
		}
		difficulty = getTextCaracters(difficulty, " ", 6)
		name := getTextCaracters(record.Name, " ", 15)
		score := getTextCaracters(fmt.Sprint(record.Score), " ", 5)

		drawButton(
			fmt.Sprintf("%s | %s | %s |  %s", difficulty, score, name, record.CreatedAt.Format("06-01-02 15:04:05")),
			entitie.Position{X: posX, Y: posY},
			font,
			screen,
			18,
		)
	}
	backButon(game, screen)
}
