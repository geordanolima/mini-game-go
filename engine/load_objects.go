package engine

import (
	"fmt"
	"log"
	"math/rand"
	"mini-game-go/domain"
	"mini-game-go/helpers"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func loadObstacules(numObstacles int, g *Game) ([]domain.Obstacle, error) {
	imagePaths := []string{"cone.png", "cone2.png", "hole.png", "truck.png", "bus.png"}
	obstacleImages := make([]*ebiten.Image, len(imagePaths))
	for i, path := range imagePaths {
		img, err := helpers.LoadImage(path)
		if err != nil {
			return nil, fmt.Errorf("error loading obstacle image %s: %w", path, err)
		}
		obstacleImages[i] = ebiten.NewImageFromImage(img)
	}
	obstacles := make([]domain.Obstacle, numObstacles)
	positionsX := []int{15, 155, 305, 455, 605}
	for i := 0; i < numObstacles; i++ {
		// get a randon image
		randomImage := obstacleImages[rand.Intn(len(obstacleImages))]
		// get randon position
		indice := rand.Intn(len(positionsX))
		addedNewObstacule := false
		if g == nil {
			addedNewObstacule = true
		}
		obsH := randomImage.Bounds().Dy()
		obsY := -obsH + 50
		for !addedNewObstacule {
			for _, item := range g.obstacles {
				h := item.Position.Height
				y := item.Position.Y
				posY := y > (obsY + h - 100)
				addedNewObstacule = item.Position.X != positionsX[indice] || posY
				if addedNewObstacule {
					break
				}
				indice = rand.Intn(len(positionsX))
			}
			for _, item := range obstacles {
				h := item.Position.Height
				y := item.Position.Y
				posY := y < (obsY + h - 100)
				addedNewObstacule = item.Position.X != positionsX[indice] || posY
				if addedNewObstacule {
					break
				}
				indice = rand.Intn(len(positionsX))
			}
		}
		if addedNewObstacule {
			obstacles[i] = domain.Obstacle{
				Position: domain.Position{
					X:      positionsX[indice],
					Y:      -randomImage.Bounds().Dy() + 50 + (int(i/4) * 1300),
					Height: randomImage.Bounds().Dy(),
				},
				Image: randomImage,
			}
			positionsX[indice] = positionsX[len(positionsX)-1]
			positionsX = positionsX[:len(positionsX)-1]
		}
	}
	return obstacles, nil
}

func loadRoad() []domain.Position {
	lines := make([]domain.Position, 21)
	positionsX := []int{150, 300, 450, 600}
	index := 0
	for line := 0; line <= 4; line++ {
		for column := 0; column <= 3; column++ {
			lines[index] = domain.Position{
				X:      positionsX[column],
				Y:      int(int(domain.LineHeight+50) * line),
				Height: int(domain.LineHeight),
			}
			index++
		}
	}
	return lines
}

func LoadTextHeader(textDraw string, position float64, fontDraw *text.GoTextFaceSource, screen *ebiten.Image, textColor string) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(position, 20)
	textColorRgba, _ := helpers.HexToRGBA(textColor)
	op.ColorM.Scale(
		float64(textColorRgba.R)/255,
		float64(textColorRgba.G)/255,
		float64(textColorRgba.B)/255,
		float64(textColorRgba.A)/255,
	)
	op.LineSpacing = 30

	text.Draw(screen, textDraw, &text.GoTextFace{
		Source: fontDraw,
		Size:   20,
	}, op)
}

func DrawRoad(screen *ebiten.Image, g *Game) {
	// create lines of road
	DrawRectGame(0, 0, domain.LineWidth, domain.GameHeight, screen, "#FFFFFF")
	DrawRectGame(domain.GameWidth-domain.LineWidth, 0, domain.LineWidth, domain.GameHeight, screen, "#FFFFFF")

	for i := 0; i < len(g.road); i++ {
		DrawRectGame(
			float64(g.road[i].X),
			float64(g.road[i].Y)+float64(g.car.Speed),
			domain.LineWidth,
			domain.LineHeight,
			screen,
			"#FFFFFF",
		)
	}
}

func getColorFuel(percent int) string {
	// stop game if fuel is 0
	switch {
	case percent > 75:
		return domain.ColorGreen
	case percent > 25:
		return domain.ColorYellow
	case percent > 0:
		return domain.ColorRed
	default:
		return ""
	}
}

func DrawRectGame(x, y, w, h float64, screen *ebiten.Image, color string) {
	recColor, err := helpers.HexToRGBA(color)
	if err != nil {
		log.Fatal(err)
	}
	ebitenutil.DrawRect(screen, x, y, w, h, recColor)
}
