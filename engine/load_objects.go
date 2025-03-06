package engine

import (
	"fmt"
	"log"
	"math/rand"
	"mini-game-go/domain"
	"mini-game-go/helpers"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func loadObstacles(numObstacles int, game *Game) ([]domain.Obstacle, error) {
	obstacles := make([]domain.Obstacle, numObstacles)
	positionsX := make([]int, len(domain.PositionsX))
	copy(positionsX, domain.PositionsX)
	for i := 0; i < numObstacles; i++ {
		randomImage := domain.ObstacleImages[rand.Intn(len(domain.ObstacleImages))]
		indice := rand.Intn(len(positionsX))
		listObstacles := obstacles
		if game != nil {
			listObstacles = game.obstacles
		}
		randomImage.Object.Position.X = positionsX[indice]
		randomImage.Object.Position.Y = -(randomImage.Image.Bounds().Dy() + (int(i/4) * 1300))
		hasConflict := verifyConflictInObstacles(randomImage.Object, listObstacles)
		if !hasConflict {
			obstacles[i] = randomImage
			obstacles[i].Object = randomImage.Object
			positionsX[indice] = positionsX[len(positionsX)-1]
			positionsX = positionsX[:len(positionsX)-1]
		}
	}
	return obstacles, nil
}

func verifyConflictInObstacles(object domain.Object, listObjects []domain.Obstacle) bool {
	results := make([]bool, len(listObjects))
	for _, item := range listObjects {
		hasConflict := verifyConflict(object, item.Object, true)
		results = append(results, hasConflict)
	}
	for _, item := range results {
		if item {
			return item
		}
	}
	return false
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

func drawGas(screen *ebiten.Image, game *Game) error {
	if game.objectGas.Image == nil {
		game.objectGas = domain.Gasoline{
			Percent:  helpers.GetProportionalPercent(),
			FilePath: "gasoline.png",
			Object: domain.Object{
				Size:     domain.Size{Width: 125, Height: 125},
				Position: domain.Position{X: domain.PositionsX[rand.Intn(len(domain.PositionsX))], Y: -1000},
				Margin:   -30,
			},
		}
		for true {
			if !verifyConflictInObstacles(game.objectGas.Object, game.obstacles) {
				break
			}
			game.objectGas.Object.Position.X = domain.PositionsX[rand.Intn(len(domain.PositionsX))]
		}
		img, err := helpers.LoadImageResize(game.objectGas.FilePath, game.objectGas.Object.Size.Width, game.objectGas.Object.Size.Height)
		if err != nil {
			return fmt.Errorf("error loading obstacle image %s: %w", game.objectGas.FilePath, err)
		}
		game.objectGas.Image = ebiten.NewImageFromImage(img)
	}
	gasImage := &ebiten.DrawImageOptions{}
	gasImage.GeoM.Translate(float64(game.objectGas.Object.Position.X), float64(game.objectGas.Object.Position.Y))
	screen.DrawImage(game.objectGas.Image, gasImage)
	allanFont, err := helpers.LoadFont("AllanRegular.ttf")
	if err != nil {
		log.Fatal(err)
	}
	LoadText(
		strconv.Itoa(game.objectGas.Percent)+"%",
		float64(game.objectGas.Object.Position.X)+45,
		float64(game.objectGas.Object.Position.Y)+55,
		30,
		allanFont,
		screen,
		domain.ColorWhite,
	)
	return nil
}
