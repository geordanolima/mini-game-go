package engine

import (
	"fmt"
	"math/rand"
	"mini-game-go/domain"
	"mini-game-go/domain/entitie"
	"mini-game-go/helpers"

	"github.com/hajimehoshi/ebiten/v2"
)

func loadObstaclesByDifficulty(difficulty domain.Difficulty, game *Game) []entitie.Obstacle {
	switch difficulty {
	case domain.Hard:
		return loadObstacles(4, game)
	case domain.Medium:
		return loadObstacles(3, game)
	default: // easy
		return loadObstacles(2, game)
	}
}

func loadObstacles(numObstacles int, game *Game) []entitie.Obstacle {
	obstacles := make([]entitie.Obstacle, numObstacles)
	positionsX := make([]float64, len(domain.PositionsX))
	copy(positionsX, domain.PositionsX)
	for i := range numObstacles {
		randomImage := domain.ObstacleImages[rand.Intn(len(domain.ObstacleImages))]
		indice := rand.Intn(len(positionsX))
		listObstacles := obstacles
		if game != nil {
			listObstacles = game.Obstacles
		}
		randomImage.Object.Position.X = positionsX[indice]
		randomImage.Object.Position.Y = float64(-(randomImage.Image.Bounds().Dy() + (int(i/4) * 1300)))
		hasConflict := verifyConflictInObstacles(randomImage.Object, listObstacles)
		if !hasConflict {
			obstacles[i] = randomImage
			obstacles[i].Object = randomImage.Object
			positionsX[indice] = positionsX[len(positionsX)-1]
			positionsX = positionsX[:len(positionsX)-1]
		}
	}
	return obstacles
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
