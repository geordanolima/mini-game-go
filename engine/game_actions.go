package engine

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"mini-game-go/database"
	"mini-game-go/database/models"
	"mini-game-go/domain"
	"mini-game-go/helpers"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	menu      domain.Menu
	car       domain.Car
	obstacles []domain.Obstacle
	score     domain.Score
	level     int
	road      []domain.Object
	roadMove  time.Time
	gameOver  domain.GameOver
	font      *text.GoTextFaceSource
	dificulty int
	objectGas domain.Gasoline
	state     domain.GameState
	user      domain.User
}

func (game *Game) NewGame() {
	LoadImageObstacleImages()
	obstaclesGame, _ := loadObstacles(5, nil)
	carSize := domain.Size{Height: 290, Width: 120}
	carImage, _ := helpers.LoadImageResize("car.png", carSize.Width, carSize.Height)

	game.car = domain.NewCar(carImage, carSize)
	game.score = domain.Score{Score: 0, Time: time.Now()}
	game.level = 01
	game.obstacles = obstaclesGame
	game.gameOver = domain.NewGameOver()
	game.state = domain.StateGameRunning
}

func createDatabase() {
	db := database.Conn()
	database.CreateTables(db)
	db.Close()
}

func saveRecord(game *Game) {
	db := database.Conn()
	err := models.InsertRecord(db, game.user.Name, game.score.Score)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}

func listRecords() []models.Record {
	db := database.Conn()
	records, err := models.GetRecords(db)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
	return records
}

func CreateGame() Game {
	createDatabase()
	font, err := helpers.LoadFont("Outfit.ttf")
	if err != nil {
		log.Fatal(err)
	}
	return Game{
		road:     loadRoad(),
		gameOver: domain.NewGameOver(),
		state:    domain.StateMenu,
		font:     font,
		menu:     domain.NewMenu(),
	}
}

func (game *Game) Update() error {
	switch game.state {
	case domain.StateMenu:
		verifyClick(game)
	case domain.StateGameRunning:
		if !game.gameOver.Flag {
			updateFuel(game)
			updateScore(game)
			updateRoad(game)
			updateCar(game)
		} else {
			if ebiten.IsKeyPressed(ebiten.KeyM) {
				game.state = domain.StateEnterName
				game.user.InputActive = true
			}
		}
	case domain.StateEnterName:
		if game.user.InputActive {
			game.handleTextInput()
		}
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			game.user.InputActive = false
			saveRecord(game)
			game.state = domain.StateMenu
		}
	case domain.StateControls:
		if ebiten.IsKeyPressed(ebiten.KeyM) {
			game.state = domain.StateMenu
		}
	case domain.StateRecords:
		if ebiten.IsKeyPressed(ebiten.KeyM) {
			game.state = domain.StateMenu
		}
	}
	return nil
}

func (game *Game) drawEnterNameScreen(screen *ebiten.Image) {
	posX := ((domain.GameWidth - domain.ButtonWidth) / 5) - 25
	vector.DrawFilledRect(screen, float32(posX), 200, float32(domain.GameWidth-(2*posX))+5, 55, color.Gray{}, false)
	vector.DrawFilledRect(screen, float32(posX), 200, float32(domain.GameWidth-(2*posX)), 50, color.White, false)

	LoadText(
		fmt.Sprintf("Your score are %d", game.score.Score),
		domain.Position{X: posX + 10, Y: 102},
		30,
		game.font,
		screen,
		domain.Black,
		text.AlignStart,
		text.AlignStart,
	)

	LoadText(
		"Write yor name and press enter to save:",
		domain.Position{X: posX + 10, Y: 202},
		20,
		game.font,
		screen,
		domain.Black,
		text.AlignStart,
		text.AlignStart,
	)

	LoadText(
		game.user.Name,
		domain.Position{X: posX + 10, Y: 235},
		25,
		game.font,
		screen,
		domain.Black,
		text.AlignStart,
		text.AlignCenter,
	)

}

func (game *Game) handleTextInput() {
	game.user.Name += string(ebiten.AppendInputChars([]rune{}))
	if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
		if len(game.user.Name) > 0 {
			game.user.Name = game.user.Name[:len(game.user.Name)-1]
		}
	}
}

func (game *Game) initBackgound(screen *ebiten.Image) {
	bacgoundColor, err := helpers.HexToRGBA(domain.Gray)
	if err != nil {
		log.Fatal(err)
	}
	screen.Fill(bacgoundColor)
}

func (game *Game) drawMenu(screen *ebiten.Image) {
	for _, action := range game.menu.Actions {
		if action.Visible {
			vector.DrawFilledRect(
				screen,
				float32(action.Object.Position.X),
				float32(action.Object.Position.Y),
				float32(action.Object.Size.Width+1),
				float32(action.Object.Size.Height+1),
				color.White,
				false,
			)
			vector.DrawFilledRect(
				screen,
				float32(action.Object.Position.X),
				float32(action.Object.Position.Y),
				float32(action.Object.Size.Width),
				float32(action.Object.Size.Height),
				color.Gray{},
				false,
			)
			LoadText(
				action.TextOptions.Text,
				action.TextOptions.Position,
				action.TextOptions.TextSize,
				game.font,
				screen,
				domain.White,
				text.AlignCenter,
				text.AlignCenter,
			)
		}
	}
}

func (game *Game) drawRecords(screen *ebiten.Image) {
	records := listRecords()
	posX := ((domain.GameWidth - domain.ButtonWidth) / 6) - 25
	i := 0
	LoadText(
		"Records",
		domain.Position{X: (domain.GameWidth / 2), Y: 50},
		30,
		game.font,
		screen,
		domain.White,
		text.AlignCenter,
		text.AlignCenter,
	)
	LoadText(
		"When          | Score           |  Name",
		domain.Position{X: posX + 20, Y: 100},
		25,
		game.font,
		screen,
		domain.White,
		text.AlignStart,
		text.AlignCenter,
	)
	for _, record := range records {
		posY := (domain.GameHeight+domain.ButtonHeight)/8 + float64(i)*(domain.ButtonHeight+20)
		vector.DrawFilledRect(
			screen,
			float32(posX),
			float32(posY),
			float32(domain.GameWidth-(2*posX)),
			float32(domain.ButtonHeight),
			color.Gray{},
			false,
		)
		LoadText(
			fmt.Sprintf("%s | %d |  %s", record.CreatedAt.Format("06-01-02 15:04:05"), record.Score, record.Name),
			domain.Position{X: posX + 20, Y: posY + 25},
			25,
			game.font,
			screen,
			domain.White,
			text.AlignStart,
			text.AlignCenter,
		)
		i++
	}
	i++
	vector.DrawFilledRect(
		screen,
		float32(20),
		float32(domain.GameHeight-domain.ButtonHeight-20),
		float32(domain.ButtonWidth),
		float32(domain.ButtonHeight),
		color.Gray{},
		false,
	)
	LoadText(
		"Voltar",
		domain.Position{X: 120, Y: domain.GameHeight - domain.ButtonHeight + 5},
		30,
		game.font,
		screen,
		domain.White,
		text.AlignCenter,
		text.AlignCenter,
	)
}

func (game *Game) drawControls(screen *ebiten.Image) {
	posX := ((domain.GameWidth - domain.ButtonWidth) / 2) - 25
	LoadText(
		"Controls",
		domain.Position{X: (domain.GameWidth / 2), Y: 50},
		30,
		game.font,
		screen,
		domain.White,
		text.AlignCenter,
		text.AlignCenter,
	)
	i := 0
	for _, action := range domain.Controls {
		posY := (domain.GameHeight+domain.ButtonHeight)/5 + float64(i)*(domain.ButtonHeight+20)
		vector.DrawFilledRect(
			screen,
			float32(posX),
			float32(posY),
			float32(domain.ButtonWidth+50),
			float32(domain.ButtonHeight),
			color.Gray{},
			false,
		)
		LoadText(
			action,
			domain.Position{X: posX + 20, Y: posY + 25},
			30,
			game.font,
			screen,
			domain.White,
			text.AlignStart,
			text.AlignCenter,
		)
		i++
	}
	i++
	vector.DrawFilledRect(
		screen,
		float32(20),
		float32(domain.GameHeight-domain.ButtonHeight-20),
		float32(domain.ButtonWidth),
		float32(domain.ButtonHeight),
		color.Gray{},
		false,
	)
	LoadText(
		"Voltar",
		domain.Position{X: 120, Y: domain.GameHeight - domain.ButtonHeight + 5},
		30,
		game.font,
		screen,
		domain.White,
		text.AlignCenter,
		text.AlignCenter,
	)
}

func (game *Game) drawGame(screen *ebiten.Image) {
	DrawRoad(screen, game)
	drawGas(screen, game)
	for _, obstacle := range game.obstacles {
		if obstacle.Image != nil {
			img := &ebiten.DrawImageOptions{}
			img.GeoM.Translate(float64(obstacle.Object.Position.X), float64(obstacle.Object.Position.Y))
			screen.DrawImage(obstacle.Image, img)
		}
	}
	DrawRectGame(0, 0, domain.GameWidth, 55, screen, domain.DarkGray)
	drawHeader(game, screen, game.font)
	if game.car.Image != nil {
		carImage := &ebiten.DrawImageOptions{}
		carImage.GeoM.Translate(float64(game.car.Object.Position.X), float64(game.car.Object.Position.Y))
		screen.DrawImage(game.car.Image, carImage)
	}
	//load gameover box
	DrawRectGame(
		float64(game.gameOver.BoxObject.Position.X),
		float64(game.gameOver.BoxObject.Position.Y),
		float64(game.gameOver.BoxObject.Size.Width),
		float64(game.gameOver.BoxObject.Size.Height),
		screen,
		domain.DarkGray,
	)
	//load gameover text
	LoadText(
		game.gameOver.TextOptions.Text,
		game.gameOver.TextOptions.Position,
		game.gameOver.TextOptions.TextSize,
		game.font,
		screen,
		domain.Red,
		text.AlignCenter,
	)
	LoadText(
		game.gameOver.TextOptions.SubText,
		domain.Position{X: game.gameOver.TextOptions.Position.X, Y: game.gameOver.TextOptions.Position.Y + 70},
		game.gameOver.TextOptions.SubTextSize,
		game.font,
		screen,
		domain.Red,
		text.AlignCenter,
	)
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.initBackgound(screen)
	switch game.state {
	case domain.StateMenu:
		game.drawMenu(screen)
	case domain.StateNewGame:
		game.NewGame()
	case domain.StateEnterName:
		game.drawEnterNameScreen(screen)
	case domain.StateGameRunning:
		game.drawGame(screen)
	case domain.StateControls:
		game.drawControls(screen)
	case domain.StateRecords:
		game.drawRecords(screen)
	}
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(domain.GameWidth), int(domain.GameHeight)
}

func drawHeader(game *Game, screen *ebiten.Image, fontDraw *text.GoTextFaceSource) {
	textFuel := "Fuel: " + strconv.Itoa(game.car.Fuel.Percent) + "%"
	textScore := "Score: " + strconv.Itoa(game.score.Score)
	textSpeed := "Speed: " + strconv.Itoa(game.car.SpeedView)
	LoadText(textFuel, domain.Position{X: 20, Y: 20}, 30, fontDraw, screen, game.car.Fuel.Color, text.AlignStart)
	LoadText(textScore, domain.Position{X: (domain.GameWidth / 2), Y: 20}, 30, fontDraw, screen, domain.White, text.AlignCenter)
	LoadText(textSpeed, domain.Position{X: domain.GameWidth - 20, Y: 20}, 30, fontDraw, screen, domain.White, text.AlignEnd)
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
