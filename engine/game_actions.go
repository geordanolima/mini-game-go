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
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var gameWidth float64 = 750
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
		positions := []int{15, 155, 305, 455, 605}
		obstacles[i] = domain.Obstacle{
			Position: domain.Position{
				X: positions[rand.Intn(len(positions))],
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
				X: 325,
				Y: 700,
			},
			Speed:  3,
			Fuel:   100,
			Angule: 0,
		},
		score:     0,
		level:     01,
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

func DrawRectGame(x, y, w, h float64, screen *ebiten.Image, color string) {
	recColor, err := helpers.HexToRGBA(color)
	if err != nil {
		log.Fatal(err)
	}
	ebitenutil.DrawRect(screen, x, y, w, h, recColor)
}

func DrawRoad(screen *ebiten.Image, speed float64) {
	// create lines of road
	var lineWidth float64 = 10
	var lineHeight float64 = 200

	DrawRectGame(0, 0, lineWidth, gameHeight, screen, "#FFFFFF")
	DrawRectGame(gameWidth-lineWidth, 0, lineWidth, gameHeight, screen, "#FFFFFF")

	for i := 0; i <= 3; i++ {
		DrawRectGame(150, (lineHeight+50)*float64(i)+speed, lineWidth, lineHeight, screen, "#FFFFFF")
		DrawRectGame(300, (lineHeight+50)*float64(i)+speed, lineWidth, lineHeight, screen, "#FFFFFF")
		DrawRectGame(450, (lineHeight+50)*float64(i)+speed, lineWidth, lineHeight, screen, "#FFFFFF")
		DrawRectGame(600, (lineHeight+50)*float64(i)+speed, lineWidth, lineHeight, screen, "#FFFFFF")
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	//  create screen
	bacgoundColor, err := helpers.HexToRGBA("#3C3C3C")
	if err != nil {
		log.Fatal(err)
	}
	screen.Fill(bacgoundColor)

	DrawRoad(screen, float64(g.car.Speed))

	// create header
	DrawRectGame(0, 0, gameWidth, 55, screen, "#0F0F0F")
	allanFont, err := helpers.LoadFont("AllanRegular.ttf")
	DrawTextHeader("Fuel: "+strconv.Itoa(g.car.Fuel)+"%", 20, allanFont, screen)
	DrawTextHeader("Score: "+strconv.Itoa(g.score), gameWidth/2, allanFont, screen)
	DrawTextHeader("Level: "+strconv.Itoa(g.level), gameWidth-75, allanFont, screen)

	for _, obstacle := range g.obstacles {
		img := &ebiten.DrawImageOptions{}
		img.GeoM.Scale(0.6, 0.6)
		img.GeoM.Translate(float64(obstacle.Position.X), float64(obstacle.Position.Y))
		screen.DrawImage(obstacle.Image, img)
	}

	if g.car.Image != nil {
		carImage := &ebiten.DrawImageOptions{}
		carImage.GeoM.Scale(0.25, 0.25)
		carImage.GeoM.Translate(float64(g.car.Position.X), float64(g.car.Position.Y))
		screen.DrawImage(g.car.Image, carImage)

		// x := g.car.Position.X
		// y := g.car.Position.Y
		// log.Printf("Position of car: X: %d, Y: %d", x, y)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(gameWidth), int(gameHeight)
}
