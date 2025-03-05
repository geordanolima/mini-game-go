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

func verifyCanAddNewObstacle(obstacles []domain.Obstacle, posX []int, newObstacle bool, indice, obsY int) (bool, int) {
	for _, item := range obstacles {
		h := item.Object.Size.Height
		y := item.Object.Position.Y
		posY := y > (obsY + h - 100)
		newObstacle = item.Object.Position.X != posX[indice] || posY
		if newObstacle {
			break
		}
		indice = rand.Intn(len(posX))
	}
	return newObstacle, indice
}

func loadObstacles(numObstacles int, game *Game) ([]domain.Obstacle, error) {
	obstacleImages := []domain.Obstacle{
		{FilePath: "cone.png", Object: domain.Object{
			Size:     domain.Size{Width: 100, Height: 100},
			Position: domain.Position{X: 20},
			Margin:   30,
		}},
		{FilePath: "cone2.png", Object: domain.Object{
			Size:     domain.Size{Width: 125, Height: 125},
			Position: domain.Position{X: 10},
			Margin:   45,
		}},
		{FilePath: "hole.png", Object: domain.Object{
			Size:     domain.Size{Width: 100, Height: 100},
			Position: domain.Position{X: 25},
			Margin:   55,
		}},
		{FilePath: "truck.png", Object: domain.Object{
			Size:     domain.Size{Width: 160, Height: 330},
			Position: domain.Position{X: 5},
			Margin:   8,
		}},
		{FilePath: "bus.png", Object: domain.Object{
			Size:     domain.Size{Width: 150, Height: 430},
			Position: domain.Position{X: 0},
			Margin:   25,
		}},
	}
	for i, obs := range obstacleImages {
		img, err := helpers.LoadImageResize(obs.FilePath, obs.Object.Size.Width, obs.Object.Size.Height)
		if err != nil {
			return nil, fmt.Errorf("error loading obstacle image %s: %w", obs.FilePath, err)
		}
		obstacleImages[i].Image = ebiten.NewImageFromImage(img)
	}
	obstacles := make([]domain.Obstacle, numObstacles)
	positionsX := []int{15, 155, 305, 455, 605}
	for i := 0; i < numObstacles; i++ {
		// get a randon image
		randomImage := obstacleImages[rand.Intn(len(obstacleImages))]
		// get randon position
		indice := rand.Intn(len(positionsX))
		// listObstacles := obstacles
		// if game != nil {
		// 	listObstacles = game.obstacles
		// }
		// for _, item := range listObstacles {
		// 	hasConflict := verifyConflict(randomImage.Object, item.Object)
		// 	if !hasConflict {
		// 		obstacles[i] = randomImage
		// 		obstacles[i].Object.Position.X += positionsX[indice]
		// 		obstacles[i].Object.Position.Y = -(randomImage.Image.Bounds().Dy() + (int(i/4) * 1300))
		// 		positionsX[indice] = positionsX[len(positionsX)-1]
		// 		positionsX = positionsX[:len(positionsX)-1]
		// 		break
		// 	}
		// }
		addedNewObstacle := false
		if game == nil {
			addedNewObstacle = true
		}

		for !addedNewObstacle {
			obsH := randomImage.Image.Bounds().Dy()
			obsY := -obsH + 50
			// TODO check the obstacles that are already added to the game
			addedNewObstacle, indice = verifyCanAddNewObstacle(game.obstacles, positionsX, addedNewObstacle, indice, obsY)
			// TODO check the obstacles that will be added to the game
			addedNewObstacle, indice = verifyCanAddNewObstacle(obstacles, positionsX, addedNewObstacle, indice, obsY)
		}
		if addedNewObstacle {
			obstacles[i] = randomImage
			obstacles[i].Object.Position.X += positionsX[indice]
			obstacles[i].Object.Position.Y = -(randomImage.Image.Bounds().Dy() + 50 + (int(i/4) * 1300))
			positionsX[indice] = positionsX[len(positionsX)-1]
			positionsX = positionsX[:len(positionsX)-1]
		}
	}
	return obstacles, nil
}

func loadRoad() []domain.Object {
	lines := make([]domain.Object, 21)
	positionsX := []int{150, 300, 450, 600}
	index := 0
	for line := 0; line <= 4; line++ {
		for column := 0; column <= 3; column++ {
			lines[index] = domain.Object{
				Position: domain.Position{
					X: positionsX[column],
					Y: int(int(domain.LineHeight+50) * line),
				},
				Size: domain.Size{
					Height: int(domain.LineHeight),
					Width:  0,
				},
			}
			index++
		}
	}
	return lines
}

func LoadText(textDraw string, posX, posY, size float64, font *text.GoTextFaceSource, screen *ebiten.Image, color string) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(posX, posY)
	colorRgba, _ := helpers.HexToRGBA(color)
	op.ColorM.Scale(
		float64(colorRgba.R)/255,
		float64(colorRgba.G)/255,
		float64(colorRgba.B)/255,
		float64(colorRgba.A)/255,
	)
	op.LineSpacing = 30

	text.Draw(screen, textDraw, &text.GoTextFace{
		Source: font,
		Size:   size,
	}, op)
}

func DrawRoad(screen *ebiten.Image, game *Game) {
	// create lines of road
	DrawRectGame(0, 0, domain.LineWidth, domain.GameHeight, screen, domain.ColorWhite)
	DrawRectGame(domain.GameWidth-domain.LineWidth, 0, domain.LineWidth, domain.GameHeight, screen, domain.ColorWhite)

	for i := 0; i < len(game.road); i++ {
		DrawRectGame(
			float64(game.road[i].Position.X),
			float64(game.road[i].Position.Y)+float64(game.car.Speed),
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
