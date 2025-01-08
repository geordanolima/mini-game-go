package engine

import (
	"fmt"
	_ "image/png"
	"log"
	"math/rand"
	"mini-game-go/domain"
	"mini-game-go/helpers"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var gameWidth float64 = 800
var gameHeight float64 = 1000

type Game struct {
	car       domain.Car
	obstacles []domain.Obstacle
	score     int
	level     int
}

func loadObstacules(numObstacles int) ([]domain.Obstacle, error) {
	imagePaths := []string{"cone.png", "hole.png", "truck.png", "bus.png"}
	obstacleImages := make([]*ebiten.Image, len(imagePaths))
	for i, path := range imagePaths {
		img, err := helpers.LoadImage(path)
		if err != nil {
			return nil, fmt.Errorf("error loading obstacle image %s: %w", path, err)
		}
		obstacleImages[i] = ebiten.NewImageFromImage(img)
	}

	obstacles := make([]domain.Obstacle, numObstacles)
	for i := 0; i < numObstacles; i++ {
		// get a randon image
		randomIndex := rand.Intn(len(obstacleImages))
		randomImage := obstacleImages[randomIndex]

		obstacles[i] = domain.Obstacle{
			Position: domain.Position{
				X: int(rand.Float64() * (gameWidth - 100)),
				Y: 50,
			},
			Image: randomImage,
		}
	}
	return obstacles, nil

}

func NewGame() (Game, error) {
	carImage, err := helpers.LoadImage("car.png")
	if err != nil {
		return Game{}, err
	}
	obstaculesGame, err := loadObstacules(3)
	if err != nil {
		return Game{}, err
	}
	return Game{
		car: domain.Car{
			Image: ebiten.NewImageFromImage(carImage),
			Position: domain.Position{
				X: 350,
				Y: 700,
			},
			Speed:  3,
			Fuel:   100,
			Angule: 0,
		},
		score:     0,
		level:     1,
		obstacles: obstaculesGame,
	}, nil
}

func (g *Game) Update() error {
	//  update position main car
	if (ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA)) && g.car.Position.X >= 5 {
		g.car.Position.X -= int(g.car.Speed)
	}
	if (ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD)) && g.car.Position.X <= 690 {
		g.car.Position.X += int(g.car.Speed)
	}
	if (ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW)) && g.car.Position.Y >= 55 {
		g.car.Position.Y -= int(g.car.Speed)
	}
	if (ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS)) && g.car.Position.Y <= 757 {
		g.car.Position.Y += int(g.car.Speed)
	}
	return nil
}

func DrawTextHeader(textDraw string, position float64, fontDraw *text.GoTextFaceSource, screen *ebiten.Image) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(position, 20)
	op.LineSpacing = 30

	text.Draw(screen, textDraw, &text.GoTextFace{
		Source: fontDraw,
		Size:   20,
	}, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	//  create screen
	colorRgba, err := helpers.HexToRGBA("#3C3C3C")
	if err != nil {
		log.Fatal(err)
	}
	screen.Fill(colorRgba)

	allanFont, err := helpers.LoadFont("AllanRegular.ttf")
	DrawTextHeader("Fuel: "+strconv.Itoa(g.car.Fuel)+"%", 20, allanFont, screen)
	DrawTextHeader("Score: "+strconv.Itoa(g.score), 350, allanFont, screen)
	DrawTextHeader("Level: "+strconv.Itoa(g.level), 700, allanFont, screen)

	if g.car.Image != nil {
		carImage := &ebiten.DrawImageOptions{}
		carImage.GeoM.Scale(0.25, 0.25)
		carImage.GeoM.Translate(float64(g.car.Position.X), float64(g.car.Position.Y))
		screen.DrawImage(g.car.Image, carImage)

		// x := g.car.Position.X
		// y := g.car.Position.Y
		// log.Printf("Position of car: X: %d, Y: %d", x, y)
	}
	for _, obstacle := range g.obstacles {
		img := &ebiten.DrawImageOptions{}
		img.GeoM.Scale(0.6, 0.6)
		img.GeoM.Translate(float64(obstacle.Position.X), float64(obstacle.Position.Y))
		screen.DrawImage(obstacle.Image, img)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(gameWidth), int(gameHeight)
}
