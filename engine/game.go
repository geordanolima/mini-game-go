package engine

import (
	"log"
	"mini-game-go/domain"
	"mini-game-go/domain/entitie"
	"mini-game-go/helpers"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	Car                     entitie.Car
	Dificulty               domain.Difficulty
	DifficultySelector      entitie.Selector
	Font                    *text.GoTextFaceSource
	GameOver                entitie.GameOver
	Level                   int
	Menu                    entitie.Selector
	ObjectGas               entitie.Obstacle
	Obstacles               []entitie.Obstacle
	Road                    []entitie.Object
	RoadMove                time.Time
	Score                   entitie.Score
	State                   entitie.GameState
	User                    entitie.User
	WaitingSelectDifficulty bool
}

func (game *Game) NewGame() {
	LoadImageObstacleImages()
	obstaclesGame := loadObstaclesByDifficulty(game.Dificulty, nil)
	carSize := entitie.Size{Height: 290, Width: 120}
	carImage, _ := helpers.LoadImageResize("car.png", carSize.Width, carSize.Height)

	game.Car = domain.NewCar(carImage, carSize, game.Dificulty)
	game.GameOver = domain.NewGameOver()
	game.Obstacles = obstaclesGame
	game.Level = 01
	game.Score = entitie.Score{Score: 0, Time: time.Now()}
	game.State = entitie.StateGameRunning
}

func CreateGame() Game {
	createDatabase()
	font, err := helpers.LoadFont("Outfit.ttf")
	if err != nil {
		log.Fatal(err)
	}
	return Game{
		Font:               font,
		GameOver:           domain.NewGameOver(),
		Menu:               domain.NewMenu(),
		DifficultySelector: domain.NewDifficultys(),
		Road:               loadRoad(),
		State:              entitie.StateMenu,
	}
}

func (game *Game) initBackgound(screen *ebiten.Image) {
	bacgoundColor, err := helpers.HexToRGBA(domain.Gray)
	if err != nil {
		log.Fatal(err)
	}
	screen.Fill(bacgoundColor)
}
