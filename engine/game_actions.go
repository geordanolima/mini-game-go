package engine

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"mini-game-go/domain"
	"mini-game-go/helpers"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	menu      domain.Menu
	car       domain.Car
	obstacles []domain.Obstacle
	score     domain.Score
	level     int
	road      []domain.Object
	roadMove  time.Time
	gameOver  domain.GameOver
	font      *text.GoTextFaceSource
	dificulty int
	objectGas domain.Gasoline
	state     domain.GameState
}

func NewGame() Game {
	font, err := helpers.LoadFont("Outfit.ttf")
	if err != nil {
		log.Fatal(err)
	}
	LoadImageObstacleImages()
	obstaclesGame, _ := loadObstacles(5, nil)
	carSize := domain.Size{Height: 290, Width: 120}
	carImage, _ := helpers.LoadImageResize("car.png", carSize.Width, carSize.Height)
	return Game{
		car:       domain.NewCar(carImage, carSize),
		score:     domain.Score{Score: 0, Time: time.Now()},
		level:     01,
		obstacles: obstaclesGame,
		road:      loadRoad(),
		gameOver:  domain.NewGameOver(),
		state:     domain.StateMenu,
		font:      font,
		menu:      domain.NewMenu(),
	}
}

func (game *Game) Update() error {
	updateFuel(game)
	updateScore(game)
	updateRoad(game)
	updateCar(game)
	return nil
}

func (game *Game) initBackgound(screen *ebiten.Image) {
	bacgoundColor, err := helpers.HexToRGBA(domain.Gray)
	if err != nil {
		log.Fatal(err)
	}
	screen.Fill(bacgoundColor)
}

func (game *Game) drawMenu(screen *ebiten.Image) {
	for _, action := range game.menu.Actions {
		if action.Visible {
			vector.DrawFilledRect(
				screen,
				float32(action.Object.Position.X),
				float32(action.Object.Position.Y),
				float32(action.Object.Size.Width),
				float32(action.Object.Size.Height),
				color.Gray{},
				false,
			)
			LoadText(
				action.TextOptions.Text,
				action.TextOptions.Position,
				action.TextOptions.TextSize,
				game.font,
				screen,
				domain.White,
				text.AlignCenter,
				text.AlignCenter,
			)
		}
	}
}

func (game *Game) drawControls(screen *ebiten.Image) {
	posX := ((domain.GameWidth - domain.ButtonWidth) / 2) - 25
	i := 0
	for _, action := range domain.Controls {
		posY := (domain.GameHeight+domain.ButtonHeight)/5 + float64(i)*(domain.ButtonHeight+20)
		vector.DrawFilledRect(
			screen,
			float32(posX),
			float32(posY),
			float32(domain.ButtonWidth+50),
			float32(domain.ButtonHeight),
			color.Gray{},
			false,
		)
		LoadText(
			action,
			domain.Position{X: posX + 20, Y: posY + 25},
			30,
			game.font,
			screen,
			domain.White,
			text.AlignStart,
			text.AlignCenter,
		)
		i++
	}
	i++
	vector.DrawFilledRect(
		screen,
		float32(20),
		float32(domain.GameHeight-domain.ButtonHeight-20),
		float32(domain.ButtonWidth),
		float32(domain.ButtonHeight),
		color.Gray{},
		false,
	)
	LoadText(
		"Voltar",
		domain.Position{X: 120, Y: domain.GameHeight - domain.ButtonHeight + 5},
		30,
		game.font,
		screen,
		domain.White,
		text.AlignCenter,
		text.AlignCenter,
	)
}

func (game *Game) drawGame(screen *ebiten.Image) {
	DrawRoad(screen, game)
	drawGas(screen, game)
	for _, obstacle := range game.obstacles {
		if obstacle.Image != nil {
			img := &ebiten.DrawImageOptions{}
			img.GeoM.Translate(float64(obstacle.Object.Position.X), float64(obstacle.Object.Position.Y))
			screen.DrawImage(obstacle.Image, img)
		}
	}
	DrawRectGame(0, 0, domain.GameWidth, 55, screen, domain.DarkGray)
	drawHeader(game, screen, game.font)
	if game.car.Image != nil {
		carImage := &ebiten.DrawImageOptions{}
		carImage.GeoM.Translate(float64(game.car.Object.Position.X), float64(game.car.Object.Position.Y))
		screen.DrawImage(game.car.Image, carImage)
	}
	//load gameover box
	DrawRectGame(
		float64(game.gameOver.BoxObject.Position.X),
		float64(game.gameOver.BoxObject.Position.Y),
		float64(game.gameOver.BoxObject.Size.Width),
		float64(game.gameOver.BoxObject.Size.Height),
		screen,
		domain.DarkGray,
	)
	//load gameover text
	LoadText(
		game.gameOver.TextOptions.Text,
		game.gameOver.TextOptions.Position,
		game.gameOver.TextOptions.TextSize,
		game.font,
		screen,
		domain.Red,
		text.AlignCenter,
	)
	LoadText(
		game.gameOver.TextOptions.SubText,
		domain.Position{X: game.gameOver.TextOptions.Position.X, Y: game.gameOver.TextOptions.Position.Y + 70},
		game.gameOver.TextOptions.SubTextSize,
		game.font,
		screen,
		domain.Red,
		text.AlignCenter,
	)
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.initBackgound(screen)
	switch game.state {
	case domain.StateMenu:
		game.drawMenu(screen)
	case domain.StateNewGame:
		game.drawGame(screen)
	case domain.StateControls:
		game.drawControls(screen)
	}
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(domain.GameWidth), int(domain.GameHeight)
}

func drawHeader(game *Game, screen *ebiten.Image, fontDraw *text.GoTextFaceSource) {
	textFuel := "Fuel: " + strconv.Itoa(game.car.Fuel.Percent) + "%"
	textScore := "Score: " + strconv.Itoa(game.score.Score)
	textSpeed := "Speed: " + strconv.Itoa(game.car.SpeedView)
	LoadText(textFuel, domain.Position{X: 20, Y: 20}, 30, fontDraw, screen, game.car.Fuel.Color, text.AlignStart)
	LoadText(textScore, domain.Position{X: (domain.GameWidth / 2), Y: 20}, 30, fontDraw, screen, domain.White, text.AlignCenter)
	LoadText(textSpeed, domain.Position{X: domain.GameWidth - 20, Y: 20}, 30, fontDraw, screen, domain.White, text.AlignEnd)
}

func LoadImageObstacleImages() error {
	for i, obs := range domain.ObstacleImages {
		img, err := helpers.LoadImageResize(obs.FilePath, obs.Object.Size.Width, obs.Object.Size.Height)
		if err != nil {
			return fmt.Errorf("error loading obstacle image %s: %w", obs.FilePath, err)
		}
		domain.ObstacleImages[i].Image = ebiten.NewImageFromImage(img)
	}
	return nil
}
