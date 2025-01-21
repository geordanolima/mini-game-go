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

func verifyCanAddNewObstacle(obstacles []domain.Obstacle, positionsX []int, addedNewObstacle bool, indice, obsY int) (bool, int) {
	for _, item := range obstacles {
		h := item.Position.Height
		y := item.Position.Y
		posY := y > (obsY + h - 100)
		addedNewObstacle = item.Position.X != positionsX[indice] || posY
		if addedNewObstacle {
			break
		}
		indice = rand.Intn(len(positionsX))
	}
	return addedNewObstacle, indice
}

func loadObstacles(numObstacles int, g *Game) ([]domain.Obstacle, error) {
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
		addedNewObstacle := false
		if g == nil {
			addedNewObstacle = true
		}
		for !addedNewObstacle {
			obsH := randomImage.Bounds().Dy()
			obsY := -obsH + 50
			// check the obstacles that are already added to the game
			addedNewObstacle, indice = verifyCanAddNewObstacle(g.obstacles, positionsX, addedNewObstacle, indice, obsY)
			// check the obstacles that will be added to the game
			addedNewObstacle, indice = verifyCanAddNewObstacle(obstacles, positionsX, addedNewObstacle, indice, obsY)
		}
		if addedNewObstacle {
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

func LoadTextHeader(textDraw string, positionX, positionY, size float64, fontDraw *text.GoTextFaceSource, screen *ebiten.Image, textColor string) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(positionX, positionY)
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
		Size:   size,
	}, op)
}

func DrawRoad(screen *ebiten.Image, g *Game) {
	// create lines of road
	DrawRectGame(0, 0, domain.LineWidth, domain.GameHeight, screen, domain.ColorWhite)
	DrawRectGame(domain.GameWidth-domain.LineWidth, 0, domain.LineWidth, domain.GameHeight, screen, domain.ColorWhite)

	for i := 0; i < len(g.road); i++ {
		DrawRectGame(
			float64(g.road[i].X),
			float64(g.road[i].Y)+float64(g.car.Speed),
			domain.LineWidth,
			domain.LineHeight,
			screen,
			domain.ColorWhite,
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
