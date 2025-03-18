package engine

import "github.com/hajimehoshi/ebiten/v2"

func (game *Game) handleTextInput() {
	game.User.Name += string(ebiten.AppendInputChars([]rune{}))
	if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
		if len(game.User.Name) > 0 {
			game.User.Name = game.User.Name[:len(game.User.Name)-1]
		}
	}
}
