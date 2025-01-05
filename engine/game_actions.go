package engine

import (
	_ "image/png"
	"log"
	"mini-game-go/domain"
	"mini-game-go/helpers"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	car   domain.Car
	score int
	level int
}

func NewGame() (Game, error) {
	carImage, err := helpers.LoadImage("car.png")
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
		score: 0,
		level: 1,
	}, nil
}

func (g *Game) Update() error {

	//  update position main car
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.car.Position.X >= 5 {
		g.car.Position.X -= int(g.car.Speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.car.Position.X <= 690 {
		g.car.Position.X += int(g.car.Speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.car.Position.Y >= 5 {
		g.car.Position.Y -= int(g.car.Speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.car.Position.Y <= 757 {
		g.car.Position.Y += int(g.car.Speed)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//  create screen
	colorRgba, err := helpers.HexToRGBA("#3C3C3C")
	if err != nil {
		log.Fatal(err)
	}
	screen.Fill(colorRgba)
	allanFont, err := helpers.LoadFont("AllanRegular.ttf")

	fuelText := "Fuel: " + strconv.Itoa(g.car.Fuel) + "%"
	if err != nil {
		log.Fatal(err)
	}
	op := &text.DrawOptions{}
	op.GeoM.Translate(20, 20)
	op.LineSpacing = 30

	text.Draw(screen, fuelText, &text.GoTextFace{
		Source: allanFont,
		Size:   20,
	}, op)

	scoreText := "Score: " + strconv.Itoa(g.score)
	if err != nil {
		log.Fatal(err)
	}
	op = &text.DrawOptions{}
	op.GeoM.Translate(350, 20)
	op.LineSpacing = 30

	text.Draw(screen, scoreText, &text.GoTextFace{
		Source: allanFont,
		Size:   20,
	}, op)

	levelText := "Level: " + strconv.Itoa(g.level)
	if err != nil {
		log.Fatal(err)
	}
	op = &text.DrawOptions{}
	op.GeoM.Translate(700, 20)
	op.LineSpacing = 30

	text.Draw(screen, levelText, &text.GoTextFace{
		Source: allanFont,
		Size:   20,
	}, op)

	if g.car.Image != nil {
		carImg := &ebiten.DrawImageOptions{}
		carImg.GeoM.Scale(0.25, 0.25)
		carImg.GeoM.Translate(float64(g.car.Position.X), float64(g.car.Position.Y))
		screen.DrawImage(g.car.Image, carImg)

		// x := g.car.Position.X
		// y := g.car.Position.Y
		// log.Printf("Position of car: X: %d, Y: %d", x, y)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 1000
}
