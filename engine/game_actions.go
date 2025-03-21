package engine

import (
	_ "image/png"
	"mini-game-go/domain"
	"mini-game-go/domain/entitie"

	"github.com/hajimehoshi/ebiten/v2"
)

func (game *Game) Update() error {
	switch game.State {
	case entitie.StateMenu:
		verifyClickMenu(game)
	case entitie.StateDifficulty:
		verifyClickDifficulty(game)
		verifyClickBack(game)
		verifyClickConfirm(game)
	case entitie.StateGameRunning:
		if !game.GameOver.Flag {
			updateFuel(game)
			updateScore(game)
			updateRoad(game)
			updateCar(game)
		} else {
			if ebiten.IsKeyPressed(ebiten.KeyM) {
				game.State = entitie.StateEnterName
				game.User.InputActive = true
			}
		}
	case entitie.StateEnterName:
		if game.User.InputActive {
			game.handleTextInput()
		}
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			game.User.InputActive = false
			saveRecord(game)
			game.State = entitie.StateMenu
		}
	case entitie.StateControls, entitie.StateRecords:
		verifyClickBack(game)
	}
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.initBackgound(screen)
	switch game.State {
	case entitie.StateMenu:
		game.drawMenu(screen)
	case entitie.StateDifficulty:
		game.drawDifficulty(screen)
	case entitie.StateNewGame:
		game.NewGame()
	case entitie.StateEnterName:
		game.drawEnterNameScreen(screen)
	case entitie.StateGameRunning:
		game.drawGame(screen)
	case entitie.StateControls:
		game.drawControls(screen)
	case entitie.StateRecords:
		game.drawRecords(screen)
	}
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(domain.GameWidth), int(domain.GameHeight)
}
