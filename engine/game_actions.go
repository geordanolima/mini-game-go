package engine

import (
	_ "image/png"
	"log"
	"mini-game-go/domain"
	"mini-game-go/helpers"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	car       domain.Car
	obstacles []domain.Obstacle
	score     int
	level     int
	road      []domain.Object
	roadMove  time.Time
	gameOver  domain.GameOver
	font      text.GoTextFaceSource
}

func NewGame() Game {
	obstaclesGame, _ := loadObstacles(5, nil)
	return Game{
		car:       domain.NewCar(),
		score:     0,
		level:     01,
		obstacles: obstaclesGame,
		road:      loadRoad(),
		gameOver:  domain.NewGameOver(),
	}
}

func (game *Game) Update() error {
	updateFuel(game)
	updateRoad(game)
	updateCar(game)
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	//  create screen
	bacgoundColor, err := helpers.HexToRGBA(domain.BacgoundColor)
	if err != nil {
		log.Fatal(err)
	}
	screen.Fill(bacgoundColor)
	allanFont, err := helpers.LoadFont("AllanRegular.ttf")
	if err != nil {
		log.Fatal(err)
	}
	DrawRoad(screen, game)
	for _, obstacle := range game.obstacles {
		img := &ebiten.DrawImageOptions{}
		// img.GeoM.Scale(0.6, 0.6)
		img.GeoM.Translate(float64(obstacle.Object.Position.X), float64(obstacle.Object.Position.Y))
		screen.DrawImage(obstacle.Image, img)
		LoadText(
			"y:["+strconv.Itoa(obstacle.Object.Position.Y)+"] yh:["+strconv.Itoa(obstacle.Object.Position.Y+obstacle.Object.Size.Height)+"]",
			float64(obstacle.Object.Position.X),
			float64(obstacle.Object.Position.Y+obstacle.Object.Size.Height-obstacle.Object.Margin),
			30,
			allanFont,
			screen,
			domain.ColorDarkGray,
		)
	}
	// create header
	DrawRectGame(0, 0, domain.GameWidth, 55, screen, domain.ColorDarkGray)
	drawHeader(game, screen, allanFont)
	DrawRectGame(
		float64(game.gameOver.BoxObject.Position.X),
		float64(game.gameOver.BoxObject.Position.Y),
		float64(game.gameOver.BoxObject.Size.Width),
		float64(game.gameOver.BoxObject.Size.Height),
		screen,
		domain.ColorDarkGray,
	)
	LoadText(
		game.gameOver.Text,
		float64(game.gameOver.TextPosition.X),
		float64(game.gameOver.TextPosition.Y),
		50,
		allanFont,
		screen,
		domain.ColorRed,
	)
	if game.car.Image != nil {
		carImage := &ebiten.DrawImageOptions{}
		carImage.GeoM.Translate(float64(game.car.Object.Position.X), float64(game.car.Object.Position.Y))
		screen.DrawImage(game.car.Image, carImage)
		LoadText("x:["+strconv.Itoa(game.car.Object.Position.X)+"] y:["+strconv.Itoa(game.car.Object.Position.Y)+"]", float64(game.car.Object.Position.X), float64(game.car.Object.Position.Y), 30, allanFont, screen, domain.ColorDarkGray)
	}
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return domain.GameWidthInt(), domain.GameHeightInt()
}

func drawHeader(game *Game, screen *ebiten.Image, fontDraw *text.GoTextFaceSource) {
	LoadText("Fuel: "+strconv.Itoa(game.car.Fuel.Percent)+"%", 20, 20, 30, fontDraw, screen, game.car.Fuel.Color)
	LoadText("Score: "+strconv.Itoa(game.score), domain.GameWidth/2, 20, 30, fontDraw, screen, domain.ColorWhite)
	LoadText("Level: "+strconv.Itoa(game.level), domain.GameWidth-75, 20, 30, fontDraw, screen, domain.ColorWhite)

	// DrawRectGame(10, 30, 700, 290, screen, domain.ColorDarkGray)
	// LoadText(strings.Join(getTextPosition("car", game.car.Object), " | "), 30, 50, 30, fontDraw, screen, domain.ColorRed)
	// LoadText(strings.Join(getTextPosition(game.obstacles[0].FilePath, game.obstacles[0].Object), " | "), 30, 100, 30, fontDraw, screen, domain.ColorRed)
	// LoadText(strings.Join(getTextPosition(game.obstacles[1].FilePath, game.obstacles[1].Object), " | "), 30, 140, 30, fontDraw, screen, domain.ColorRed)
	// LoadText(strings.Join(getTextPosition(game.obstacles[2].FilePath, game.obstacles[2].Object), " | "), 30, 180, 30, fontDraw, screen, domain.ColorRed)
	// LoadText(strings.Join(getTextPosition(game.obstacles[3].FilePath, game.obstacles[3].Object), " | "), 30, 220, 30, fontDraw, screen, domain.ColorRed)
	// LoadText(strings.Join(getTextPosition(game.obstacles[4].FilePath, game.obstacles[4].Object), " | "), 30, 260, 30, fontDraw, screen, domain.ColorRed)
}

// func getTextPosition(name string, obj domain.Object) []string {
// 	return []string{
// 		name,
// 		fmt.Sprint(obj.Angule),
// 		"x", strconv.Itoa(obj.Position.X),
// 		"y", strconv.Itoa(obj.Position.Y),
// 		"x+w", strconv.Itoa(obj.Position.X + obj.Size.Width),
// 		"y+h", strconv.Itoa(obj.Position.Y + obj.Size.Height),
// 	}
// }
