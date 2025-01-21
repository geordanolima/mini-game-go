package engine

import (
	"mini-game-go/domain"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func updateRoad(g *Game) {
	if g.gameOver.Flag {
		return
	}
	if g.car.Fuel.Percent == 0 {
		g.gameOver.TextPosition.Y = 300
		g.gameOver.BoxPosition.Y = 300

		g.gameOver.Flag = true
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
				obstaclesGame, _ := loadObstacles(1, g)
				g.obstacles = append(g.obstacles, obstaclesGame[0])
			}
		}
	}
}

func updateFuel(g *Game) {
	if g.gameOver.Flag {
		return
	}
	colorFuel := getColorFuel(g.car.Fuel.Percent)
	if colorFuel == "" {
		g.gameOver.Flag = false
	}
	g.car.Fuel.Color = colorFuel

	if time.Since(g.car.Fuel.Time) >= 1000*time.Millisecond {
		g.car.Fuel.Time = time.Now()
		fuelConsume := int(g.car.Speed / 5)
		if fuelConsume <= 0 {
			fuelConsume = 1
		}
		g.car.Fuel.Percent -= fuelConsume
		if g.car.Fuel.Percent < 0 {
			g.car.Fuel.Percent = 0
			g.gameOver.Flag = true
		}
	}
}

func updateCar(g *Game) {
	if g.gameOver.Flag {
		return
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
}
