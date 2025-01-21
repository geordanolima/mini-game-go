package engine

import (
	_ "image/png"
	"log"
	"mini-game-go/domain"
	"mini-game-go/helpers"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	car       domain.Car
	obstacles []domain.Obstacle
	score     int
	level     int
	road      []domain.Position
	roadMove  time.Time
	gameOver  domain.GameOver
}

func NewGame() (Game, error) {
	carImage, err := helpers.LoadImage("car.png")
	if err != nil {
		return Game{}, err
	}
	obstaclesGame, err := loadObstacles(5, nil)
	if err != nil {
		return Game{}, err
	}
	return Game{
		car: domain.Car{
			Image: ebiten.NewImageFromImage(carImage),
			Position: domain.Position{
				X: 325,
				Y: 700,
			},
			Speed: 15,
			Fuel: domain.Fuel{
				Percent: 100,
				Time:    time.Now(),
				Color:   domain.ColorGreen,
			},
			Angule: 0,
		},
		score:     0,
		level:     01,
		obstacles: obstaclesGame,
		road:      loadRoad(),
		gameOver: domain.GameOver{
			Flag:         false,
			BoxPosition:  domain.Position{X: 280, Y: -100, Width: 210, Height: 60},
			TextPosition: domain.Position{X: 300, Y: -100},
		},
	}, nil
}

func (g *Game) Update() error {
	updateFuel(g)
	updateRoad(g)
	updateCar(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//  create screen
	bacgoundColor, err := helpers.HexToRGBA(domain.BacgoundColor)
	if err != nil {
		log.Fatal(err)
	}
	screen.Fill(bacgoundColor)
	DrawRoad(screen, g)
	for _, obstacle := range g.obstacles {
		img := &ebiten.DrawImageOptions{}
		img.GeoM.Scale(0.6, 0.6)
		img.GeoM.Translate(float64(obstacle.Position.X), float64(obstacle.Position.Y))
		screen.DrawImage(obstacle.Image, img)
	}
	// create header
	DrawRectGame(0, 0, domain.GameWidth, 55, screen, domain.ColorDarkGray)
	allanFont, err := helpers.LoadFont("AllanRegular.ttf")
	if err != nil {
		log.Fatal(err)
	}
	LoadTextHeader("Fuel: "+strconv.Itoa(g.car.Fuel.Percent)+"%", 20, 20, 30, allanFont, screen, g.car.Fuel.Color)
	LoadTextHeader("Score: "+strconv.Itoa(g.score), domain.GameWidth/2, 20, 30, allanFont, screen, domain.ColorWhite)
	LoadTextHeader("Level: "+strconv.Itoa(g.level), domain.GameWidth-75, 20, 30, allanFont, screen, domain.ColorWhite)
	DrawRectGame(float64(g.gameOver.BoxPosition.X), float64(g.gameOver.BoxPosition.Y), float64(g.gameOver.BoxPosition.Width), float64(g.gameOver.BoxPosition.Height), screen, domain.ColorDarkGray)
	LoadTextHeader("Game Over", float64(g.gameOver.TextPosition.X), float64(g.gameOver.TextPosition.Y), 50, allanFont, screen, domain.ColorRed)
	if g.car.Image != nil {
		carImage := &ebiten.DrawImageOptions{}
		carImage.GeoM.Scale(0.25, 0.25)
		carImage.GeoM.Translate(float64(g.car.Position.X), float64(g.car.Position.Y))
		screen.DrawImage(g.car.Image, carImage)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(domain.GameWidth), int(domain.GameHeight)
}
