package engine

import (
	"math/rand"
	"mini-game-go/domain"
	"mini-game-go/helpers"
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
		if game.road[i].Position.Y > domain.GameHeight {
			game.road[i].Position.Y = -game.road[i].Size.Height - 50
		}
	}
}

func moveObstaclesRoad(game *Game) {
	for i := 0; i < len(game.obstacles); i++ {
		game.obstacles[i].Object.Position.Y += game.car.Speed * 2
		if game.obstacles[i].Object.Position.Y > domain.GameHeight {
			game.obstacles = append(game.obstacles[:i], game.obstacles[i+1:]...)
			obstaclesGame, _ := loadObstacles(1, game)
			game.obstacles = append(game.obstacles, obstaclesGame[0])
		}
	}
	game.objectGas.Object.Position.Y += game.car.Speed * 2
	if game.objectGas.Object.Position.Y > domain.GameHeight {
		game.objectGas.Object.Position.Y = -1000
		game.objectGas.Percent = helpers.GetProportionalPercent()
		game.objectGas.Object.Position.X = domain.PositionsX[rand.Intn(len(domain.PositionsX))]
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
		if game.car.Fuel.Percent%5 == 0 {
			for i := 9; i > 0; i-- {
				if game.car.Fuel.Percent > i*10 {
					game.car.Speed += float64(10 - i + game.dificulty)
					game.car.SpeedView += 5 * (10 - i)
					break
				}
			}
		}
		if game.car.Fuel.Percent <= 0 {
			game.car.Fuel.Percent = 0
			drawGameOver(game)
		}
	}
}

func verifyConflict(mainObject, conflictObject domain.Object, validateHeight bool) bool {
	mainWidth := mainObject.Size.Width
	mainHeight := mainObject.Size.Height
	mainX := mainObject.Position.X
	mainY := mainObject.Position.Y

	confWidth := conflictObject.Size.Width
	confHeight := conflictObject.Size.Height
	confX := conflictObject.Position.X
	confY := conflictObject.Position.Y
	margin := conflictObject.Margin

	overlapX := mainX < confX+confWidth+margin && mainX+mainWidth+margin > confX
	overlapY := mainY < confY+confHeight+margin && mainY+mainHeight+margin > confY

	if validateHeight {
		conflitoDir := (mainX+mainWidth+margin <= confX+confWidth || mainX+margin <= confX+confWidth)
		conflitoEsc := (mainX >= confX-margin || mainX+mainWidth >= confX+margin)

		confSup := mainY <= confY+confHeight-margin
		confInf := mainY+mainHeight-margin >= confY

		confH := confY < mainY+mainHeight && mainY < confY && mainY < confY+confHeight
		return overlapX && overlapY && conflitoDir && conflitoEsc && confSup && confInf && confH
	}
	return overlapX && overlapY
}

func updateScore(game *Game) {
	if game.gameOver.Flag {
		return
	}
	if time.Since(game.score.Time) >= 1000*time.Millisecond {
		game.score.Score += int(game.car.Speed / 5)
		game.score.Time = time.Now()
	}
}

func updateCar(game *Game) {
	if ebiten.IsKeyPressed(ebiten.KeyN) {
		game.state = domain.StateMenu
	}
	if game.gameOver.Flag {
		return
	}
	// verify game over
	for i := 0; i < len(game.obstacles); i++ {
		if verifyConflict(game.car.Object, game.obstacles[i].Object, false) {
			drawGameOver(game)
		}
	}
	// verify if get fuel
	if verifyConflict(game.car.Object, game.objectGas.Object, true) {
		game.objectGas.Image = nil
		game.car.Fuel.Percent += game.objectGas.Percent
		if game.car.Fuel.Percent > 100 {
			game.car.Fuel.Percent = 100
		}
	}
	// move car
	if (ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA)) &&
		game.car.Object.Position.X >= 20 {
		game.car.Object.Position.X -= game.car.Speed
	}
	if (ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD)) &&
		game.car.Object.Position.X <= domain.GameWidth-game.car.Object.Size.Width-20 {
		game.car.Object.Position.X += game.car.Speed
	}
	if (ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW)) &&
		game.car.Object.Position.Y >= 60 {
		game.car.Object.Position.Y -= game.car.Speed
	}
	if (ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS)) &&
		game.car.Object.Position.Y <= domain.GameWidth-20 {
		game.car.Object.Position.Y += game.car.Speed
	}
	mouseX, mouseY := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		for _, action := range game.menu.Actions {
			if mouseX >= int(action.Object.Position.X) && mouseX <= int(action.Object.Position.X+action.Object.Size.Width) &&
				mouseY >= int(action.Object.Position.Y) && mouseY <= int(action.Object.Position.Y+action.Object.Size.Height) {
				game.state = action.State
			}
		}
	}
}

func drawGameOver(game *Game) {
	game.gameOver.TextOptions.Position.Y = 300
	game.gameOver.BoxObject.Position.Y = 300
	game.gameOver.Flag = true
}
