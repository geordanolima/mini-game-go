package engine

import (
	"fmt"
	"math/rand"
	"mini-game-go/domain"
	"mini-game-go/domain/entitie"
	"mini-game-go/helpers"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func getColorFuel(percent int) string {
	switch {
	case percent > 75:
		return domain.Green
	case percent > 25:
		return domain.Yellow
	default: // percent > 0
		return domain.Red
	}
}

func getPosYFuelByDifficulty(game *Game) float64 {
	switch game.Dificulty {
	case domain.Hard:
		return -2000
	case domain.Medium:
		return -1000
	default: // easy
		return -700
	}
}

func defineObjectFuel(game *Game) entitie.Obstacle {
	value := helpers.GetProportionalPercent(domain.PercentGasByDifficulty[game.Dificulty])
	return entitie.Obstacle{
		TextValue: strconv.Itoa(value) + "%",
		Value:     value,
		Type:      entitie.ObjectGasoline,
		FilePath:  "gasoline.png",
		Object: entitie.Object{
			Size: entitie.Size{Width: 125, Height: 125},
			Position: entitie.Position{
				X: domain.PositionsX[rand.Intn(len(domain.PositionsX))],
				Y: getPosYFuelByDifficulty(game),
			},
			Margin: -30,
		},
	}
}

func drawGas(screen *ebiten.Image, game *Game) error {
	if game.ObjectGas.Image == nil {
		game.ObjectGas = defineObjectFuel(game)
		for {
			if !verifyConflictInObstacles(game.ObjectGas.Object, game.Obstacles) {
				break
			}
			game.ObjectGas.Object.Position.X = domain.PositionsX[rand.Intn(len(domain.PositionsX))]
		}
		img, err := helpers.LoadImageResize(
			game.ObjectGas.FilePath,
			game.ObjectGas.Object.Size.Width,
			game.ObjectGas.Object.Size.Height,
		)
		if err != nil {
			return fmt.Errorf("error loading obstacle image %s: %w", game.ObjectGas.FilePath, err)
		}
		game.ObjectGas.Image = ebiten.NewImageFromImage(img)
	}
	gasImage := &ebiten.DrawImageOptions{}
	gasImage.GeoM.Translate(game.ObjectGas.Object.Position.X, game.ObjectGas.Object.Position.Y)
	screen.DrawImage(game.ObjectGas.Image, gasImage)
	LoadText(
		game.ObjectGas.TextValue,
		entitie.Position{X: game.ObjectGas.Object.Position.X + 45, Y: game.ObjectGas.Object.Position.Y + 55},
		30,
		game.Font,
		screen,
		domain.White,
		text.AlignStart,
		text.AlignStart,
	)
	return nil
}
