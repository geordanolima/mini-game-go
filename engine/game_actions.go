package engine

import (
	"fmt"
	_ "image/png"
	"log"
	"mini-game-go/domain"
	"mini-game-go/helpers"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	car       domain.Car
	obstacles []domain.Obstacle
	score     int
	level     int
	road      []domain.Object
	roadMove  time.Time
	gameOver  domain.GameOver
	font      text.GoTextFaceSource
	dificulty int
	objectGas domain.Gasoline
}

func NewGame() Game {
	LoadImageObstacleImages()
	obstaclesGame, _ := loadObstacles(5, nil)
	carSize := domain.Size{Height: 290, Width: 120}
	carImage, _ := helpers.LoadImageResize("car.png", carSize.Width, carSize.Height)
	return Game{
		car:       domain.NewCar(carImage, carSize),
		score:     0,
		level:     01,
		obstacles: obstaclesGame,
		road:      loadRoad(),
		gameOver:  domain.NewGameOver(),
	}
}

func (game *Game) Update() error {
	updateFuel(game)
	updateRoad(game)
	updateCar(game)
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	bacgoundColor, err := helpers.HexToRGBA(domain.BacgoundColor)
	if err != nil {
		log.Fatal(err)
	}
	screen.Fill(bacgoundColor)
	allanFont, err := helpers.LoadFont("AllanRegular.ttf")
	if err != nil {
		log.Fatal(err)
	}
	DrawRoad(screen, game)
	drawGas(screen, game)
	for _, obstacle := range game.obstacles {
		img := &ebiten.DrawImageOptions{}
		img.GeoM.Translate(float64(obstacle.Object.Position.X), float64(obstacle.Object.Position.Y))
		screen.DrawImage(obstacle.Image, img)
		// LoadText(
		// 	"y:["+strconv.Itoa(obstacle.Object.Position.Y)+"] yh:["+strconv.Itoa(obstacle.Object.Position.Y+obstacle.Object.Size.Height)+"]",
		// 	float64(obstacle.Object.Position.X),
		// 	float64(obstacle.Object.Position.Y+obstacle.Object.Size.Height-obstacle.Object.Margin),
		// 	30,
		// 	allanFont,
		// 	screen,
		// 	domain.ColorDarkGray,
		// )
	}
	DrawRectGame(0, 0, domain.GameWidth, 55, screen, domain.ColorDarkGray)
	drawHeader(game, screen, allanFont)
	//load gameover box
	DrawRectGame(
		float64(game.gameOver.BoxObject.Position.X),
		float64(game.gameOver.BoxObject.Position.Y),
		float64(game.gameOver.BoxObject.Size.Width),
		float64(game.gameOver.BoxObject.Size.Height),
		screen,
		domain.ColorDarkGray,
	)
	//load gameover text
	LoadText(
		game.gameOver.Text,
		float64(game.gameOver.TextPosition.X),
		float64(game.gameOver.TextPosition.Y),
		50,
		allanFont,
		screen,
		domain.ColorRed,
	)
	if game.car.Image != nil {
		carImage := &ebiten.DrawImageOptions{}
		carImage.GeoM.Translate(float64(game.car.Object.Position.X), float64(game.car.Object.Position.Y))
		screen.DrawImage(game.car.Image, carImage)
		// LoadText("x:["+strconv.Itoa(game.car.Object.Position.X)+"] y:["+strconv.Itoa(game.car.Object.Position.Y)+"]", float64(game.car.Object.Position.X), float64(game.car.Object.Position.Y), 30, allanFont, screen, domain.ColorDarkGray)
	}
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return domain.GameWidthInt(), domain.GameHeightInt()
}

func drawHeader(game *Game, screen *ebiten.Image, fontDraw *text.GoTextFaceSource) {
	LoadText("Fuel: "+strconv.Itoa(game.car.Fuel.Percent)+"%", 20, 20, 30, fontDraw, screen, game.car.Fuel.Color)
	LoadText("Score: "+strconv.Itoa(game.score), domain.GameWidth/2, 20, 30, fontDraw, screen, domain.ColorWhite)
	LoadText("Speed: "+strconv.Itoa(game.car.SpeedView), domain.GameWidth-120, 20, 30, fontDraw, screen, domain.ColorWhite)
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
