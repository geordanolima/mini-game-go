package engine

import (
	"math/rand"
	"mini-game-go/domain"
	"mini-game-go/domain/entitie"
	"mini-game-go/helpers"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func updateRoad(game *Game) {
	if time.Since(game.RoadMove) >= 50*time.Millisecond {
		game.RoadMove = time.Now()
		moveLinesRoad(game)
		moveObstaclesRoad(game)
	}
}

func moveLinesRoad(game *Game) {
	for i := 0; i < len(game.Road); i++ {
		game.Road[i].Position.Y += game.Car.Speed * 2
		if game.Road[i].Position.Y > domain.GameHeight {
			game.Road[i].Position.Y = -game.Road[i].Size.Height - 50
		}
	}
}

func moveObstaclesRoad(game *Game) {
	for i := 0; i < len(game.Obstacles); i++ {
		game.Obstacles[i].Object.Position.Y += game.Car.Speed * 2
		if game.Obstacles[i].Object.Position.Y > domain.GameHeight {
			game.Obstacles = append(game.Obstacles[:i], game.Obstacles[i+1:]...)
			obstaclesGame := loadObstacles(1, game)
			game.Obstacles = append(game.Obstacles, obstaclesGame[0])
		}
	}
	game.ObjectGas.Object.Position.Y += game.Car.Speed * 2
	if game.ObjectGas.Object.Position.Y > domain.GameHeight {
		game.ObjectGas.Object.Position.Y = -1000
		value := helpers.GetProportionalPercent()
		game.ObjectGas.TextValue = strconv.Itoa(value) + "%"
		game.ObjectGas.Value = value
		game.ObjectGas.Object.Position.X = domain.PositionsX[rand.Intn(len(domain.PositionsX))]
	}
}

func updateFuel(game *Game) {
	colorFuel := getColorFuel(game.Car.Fuel.Percent)
	game.Car.Fuel.Color = colorFuel

	if time.Since(game.Car.Fuel.Time) >= 1000*time.Millisecond {
		game.Car.Fuel.Time = time.Now()
		taxConsume := int(game.Car.Speed / 5)
		if taxConsume <= 0 {
			taxConsume = 1
		}
		game.Car.Fuel.Percent -= taxConsume
		if game.Car.Fuel.Percent%5 == 0 {
			for i := 9; i > 0; i-- {
				if game.Car.Fuel.Percent > i*10 {
					game.Car.Speed += float64(10 - i + int(game.Dificulty))
					game.Car.SpeedView += 5 * (10 - i)
					break
				}
			}
		}
		if game.Car.Fuel.Percent <= 0 {
			game.Car.Fuel.Percent = 0
			drawGameOver(game)
		}
	}
}

func verifyConflict(mainObject, conflictObject entitie.Object, validateHeight bool) bool {
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
	if time.Since(game.Score.Time) >= 1000*time.Millisecond {
		game.Score.Score += int(game.Car.Speed / 5)
		game.Score.Time = time.Now()
	}
}

func updateCar(game *Game) {
	// verify game over
	for _, obstacle := range game.Obstacles {
		if verifyConflict(game.Car.Object, obstacle.Object, false) {
			drawGameOver(game)
		}
	}
	// verify if get fuel
	if verifyConflict(game.Car.Object, game.ObjectGas.Object, true) {
		game.ObjectGas.Image = nil
		game.Car.Fuel.Percent += game.ObjectGas.Value
		if game.Car.Fuel.Percent > 100 {
			game.Car.Fuel.Percent = 100
		}
	}
	// move car
	if (ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA)) &&
		game.Car.Object.Position.X >= 20 {
		game.Car.Object.Position.X -= game.Car.Speed
	}
	if (ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD)) &&
		game.Car.Object.Position.X <= domain.GameWidth-game.Car.Object.Size.Width-20 {
		game.Car.Object.Position.X += game.Car.Speed
	}
	if (ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW)) &&
		game.Car.Object.Position.Y >= 60 {
		game.Car.Object.Position.Y -= game.Car.Speed
	}
	if (ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS)) &&
		game.Car.Object.Position.Y <= domain.GameHeight-300 {
		game.Car.Object.Position.Y += game.Car.Speed
	}
}

func verifyClick(game *Game) {
	mouseX, mouseY := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if game.State == entitie.StateMenu {
			for _, action := range game.Menu.Actions {
				if mouseX >= int(action.Object.Position.X) && mouseX <= int(action.Object.Position.X+action.Object.Size.Width) &&
					mouseY >= int(action.Object.Position.Y) && mouseY <= int(action.Object.Position.Y+action.Object.Size.Height) {
					game.State = action.State
				}
			}
		}
		if game.State == entitie.StateControls {
			if mouseX >= 20 && mouseX <= int(domain.ButtonWidth+20) &&
				mouseY >= int(domain.GameHeight-domain.ButtonHeight-20) && mouseY <= int(domain.GameHeight-20) {
				game.State = entitie.StateMenu
			}
		}
	}
}

func drawGameOver(game *Game) {
	game.GameOver.TextOptions.Position.Y = 300
	game.GameOver.BoxObject.Position.Y = 300
	game.GameOver.Flag = true
}
