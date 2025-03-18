package engine

import (
	"mini-game-go/domain"
	"mini-game-go/domain/entitie"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (game *Game) drawGame(screen *ebiten.Image) {
	DrawRoad(screen, game)
	drawGas(screen, game)
	for _, obstacle := range game.Obstacles {
		if obstacle.Image != nil {
			img := &ebiten.DrawImageOptions{}
			img.GeoM.Translate(float64(obstacle.Object.Position.X), float64(obstacle.Object.Position.Y))
			screen.DrawImage(obstacle.Image, img)
		}
	}
	DrawRect64(0, 0, domain.GameWidth, 55, screen, domain.DarkGray)
	drawHeader(game, screen, game.Font)
	if game.Car.Image != nil {
		carImage := &ebiten.DrawImageOptions{}
		carImage.GeoM.Translate(float64(game.Car.Object.Position.X), float64(game.Car.Object.Position.Y))
		screen.DrawImage(game.Car.Image, carImage)
	}
	//load Gameover box
	DrawRect64(
		game.GameOver.BoxObject.Position.X,
		game.GameOver.BoxObject.Position.Y,
		game.GameOver.BoxObject.Size.Width,
		game.GameOver.BoxObject.Size.Height,
		screen,
		domain.DarkGray,
	)
	//load Gameover text
	LoadText(
		game.GameOver.TextOptions.Text,
		game.GameOver.TextOptions.Position,
		game.GameOver.TextOptions.TextSize,
		game.Font,
		screen,
		domain.Red,
		text.AlignCenter,
		text.AlignStart,
	)
	LoadText(
		game.GameOver.TextOptions.SubText,
		entitie.Position{X: game.GameOver.TextOptions.Position.X, Y: game.GameOver.TextOptions.Position.Y + 70},
		game.GameOver.TextOptions.SubTextSize,
		game.Font,
		screen,
		domain.Red,
		text.AlignCenter,
		text.AlignStart,
	)
}

func drawHeader(game *Game, screen *ebiten.Image, fontDraw *text.GoTextFaceSource) {
	textFuel := "Fuel: " + strconv.Itoa(game.Car.Fuel.Percent) + "%"
	textScore := "Score: " + strconv.Itoa(game.Score.Score)
	textSpeed := "Speed: " + strconv.Itoa(game.Car.SpeedView)
	LoadText(textFuel, entitie.Position{X: 20, Y: 20}, 30, fontDraw, screen, game.Car.Fuel.Color, text.AlignStart, text.AlignStart)
	LoadText(textScore, entitie.Position{X: (domain.GameWidth / 2), Y: 20}, 30, fontDraw, screen, domain.White, text.AlignCenter, text.AlignStart)
	LoadText(textSpeed, entitie.Position{X: domain.GameWidth - 20, Y: 20}, 30, fontDraw, screen, domain.White, text.AlignEnd, text.AlignStart)
}
