package engine

import (
	"mini-game-go/domain"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func updateRoad(game *Game) {
	if game.gameOver.Flag {
		return
	}
	if time.Since(game.roadMove) >= 50*time.Millisecond {
		game.roadMove = time.Now()
		moveLinesRoad(game)
		moveObstaclesRoad(game)
	}
}

func moveLinesRoad(game *Game) {
	for i := 0; i < len(game.road); i++ {
		game.road[i].Position.Y += game.car.Speed * 2
		if game.road[i].Position.Y > domain.GameHeightInt() {
			game.road[i].Position.Y = -game.road[i].Size.Height - 50
		}
	}
}
func moveObstaclesRoad(game *Game) {
	for i := 0; i < len(game.obstacles); i++ {
		game.obstacles[i].Object.Position.Y += game.car.Speed * 2
		if game.obstacles[i].Object.Position.Y > domain.GameHeightInt() {
			game.obstacles = append(game.obstacles[:i], game.obstacles[i+1:]...)
			obstaclesGame, _ := loadObstacles(1, game)
			game.obstacles = append(game.obstacles, obstaclesGame[0])
		}
	}
}

func updateFuel(game *Game) {
	if game.gameOver.Flag {
		return
	}
	colorFuel := getColorFuel(game.car.Fuel.Percent)
	game.car.Fuel.Color = colorFuel

	if time.Since(game.car.Fuel.Time) >= 1000*time.Millisecond {
		game.car.Fuel.Time = time.Now()
		taxConsume := int(game.car.Speed / 5)
		if taxConsume <= 0 {
			taxConsume = 1
		}
		game.car.Fuel.Percent -= taxConsume
		if game.car.Fuel.Percent <= 0 {
			game.car.Fuel.Percent = 0
			drawGameOver(game)
		}
	}
}

func verifyConflict(mainObject, conflictObject domain.Object) bool {
	mainWidth := mainObject.Size.Width
	mainHeight := mainObject.Size.Height
	mainX := mainObject.Position.X
	mainY := mainObject.Position.Y

	confWidth := conflictObject.Size.Width
	confHeight := conflictObject.Size.Height
	confX := conflictObject.Position.X
	confY := conflictObject.Position.Y
	margin := conflictObject.Margin

	conflitoDir := (mainX+mainWidth+margin <= confX+confWidth || mainX+margin <= conflictObject.Position.X+confWidth)
	conflitoEsc := (mainX+margin >= conflictObject.Position.X || mainX+mainWidth+margin >= conflictObject.Position.X)
	confSup := mainY-margin <= confY+confHeight
	confInf := mainY+mainHeight+margin >= confY

	return conflitoDir && conflitoEsc && confSup && confInf
}

func updateCar(game *Game) {
	if game.gameOver.Flag {
		return
	}
	for i := 0; i < len(game.obstacles); i++ {
		if verifyConflict(game.car.Object, game.obstacles[i].Object) {
			drawGameOver(game)
		}
	}
	//  update position main car
	if (ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA)) &&
		game.car.Object.Position.X >= 20 {
		game.car.Object.Position.X -= int(game.car.Speed)
	}
	if (ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD)) &&
		game.car.Object.Position.X <= domain.GameWidthInt()-game.car.Object.Size.Width-20 {
		game.car.Object.Position.X += int(game.car.Speed)
	}
	if (ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW)) &&
		game.car.Object.Position.Y >= 60 {
		game.car.Object.Position.Y -= int(game.car.Speed)
	}
	if (ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS)) &&
		game.car.Object.Position.Y <= domain.GameWidthInt()-20 {
		game.car.Object.Position.Y += int(game.car.Speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if game.car.Speed == 0 {
			game.car.Speed = 5
		} else {
			game.car.Speed = 0
		}
	}
}

func drawGameOver(game *Game) {
	game.gameOver.TextPosition.Y = 300
	game.gameOver.BoxObject.Position.Y = 300
	game.gameOver.Flag = true
}
