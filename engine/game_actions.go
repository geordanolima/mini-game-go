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
	allanFont, err := helpers.LoadFont("AllanRegular.ttf")
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
		font:      allanFont,
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

func (game *Game) drawMenu(screen *ebiten.Image) {
	bacgoundColor, err := helpers.HexToRGBA(domain.Gray)
	if err != nil {
		log.Fatal(err)
	}
	screen.Fill(bacgoundColor)
	// DrawRoad(screen, game)
	for _, action := range game.menu.Actions {
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

func (game *Game) drawGame(screen *ebiten.Image) {
	bacgoundColor, err := helpers.HexToRGBA(domain.Gray)
	if err != nil {
		log.Fatal(err)
	}
	screen.Fill(bacgoundColor)
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
}

func (game *Game) Draw(screen *ebiten.Image) {
	switch game.state {
	case domain.StateMenu:
		game.drawMenu(screen)
	case domain.StateNewGame:
		game.drawGame(screen)
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
