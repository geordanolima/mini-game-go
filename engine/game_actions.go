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
}

func NewGame() (Game, error) {
	carImage, err := helpers.LoadImage("car.png")
	if err != nil {
		return Game{}, err
	}
	obstaculesGame, err := loadObstacules(5, nil)
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
		obstacles: obstaculesGame,
		road:      loadRoad(),
	}, nil
}

func (g *Game) Update() error {
	colorFuel := getColorFuel(g.car.Fuel.Percent)
	if colorFuel != "" {
		g.car.Fuel.Color = colorFuel
	} else {
		return nil
	}
	if time.Since(g.car.Fuel.Time) >= 1000*time.Millisecond {
		g.car.Fuel.Time = time.Now()
		fuelConsume := int(g.car.Speed / 5)
		if fuelConsume <= 0 {
			fuelConsume = 1
		}
		g.car.Fuel.Percent -= fuelConsume
	}
	if time.Since(g.roadMove) >= 50*time.Millisecond {
		g.roadMove = time.Now()
		// Move road
		for i := 0; i < len(g.road); i++ {
			g.road[i].Y += g.car.Speed * 2
			if g.road[i].Y > int(domain.GameHeight) {
				g.road[i].Y = -g.road[i].Height - 50
			}
		}
		for i := 0; i < len(g.obstacles); i++ {
			g.obstacles[i].Position.Y += g.car.Speed * 2
			if g.obstacles[i].Position.Y > int(domain.GameHeight) {
				g.obstacles = append(g.obstacles[:i], g.obstacles[i+1:]...)
				obstaculesGame, _ := loadObstacules(1, g)
				g.obstacles = append(g.obstacles, obstaculesGame[0])
			}
		}
	}
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

func (g *Game) Draw(screen *ebiten.Image) {
	//  create screen
	bacgoundColor, err := helpers.HexToRGBA("#3C3C3C")
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
	DrawRectGame(0, 0, domain.GameWidth, 55, screen, "#0F0F0F")
	allanFont, err := helpers.LoadFont("AllanRegular.ttf")
	LoadTextHeader("Fuel: "+strconv.Itoa(g.car.Fuel.Percent)+"%", 20, allanFont, screen, g.car.Fuel.Color)
	LoadTextHeader("Score: "+strconv.Itoa(g.score), domain.GameWidth/2, allanFont, screen, "#FFFFFF")
	LoadTextHeader("Level: "+strconv.Itoa(g.level), domain.GameWidth-75, allanFont, screen, "#FFFFFF")
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
