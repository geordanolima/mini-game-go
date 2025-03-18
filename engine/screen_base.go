package engine

import (
	"mini-game-go/domain"
	"mini-game-go/domain/entitie"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func title(title, subtitle string, font *text.GoTextFaceSource, screen *ebiten.Image) {
	posX := (domain.GameWidth / 2)
	LoadText(
		title, entitie.Position{X: posX, Y: 50}, 30, font, screen, domain.White, text.AlignCenter, text.AlignCenter,
	)
	LoadText(
		subtitle, entitie.Position{X: posX, Y: 100}, 25, font, screen, domain.White, text.AlignCenter, text.AlignCenter,
	)
}

func drawButton(label string, position entitie.Position, font *text.GoTextFaceSource, screen *ebiten.Image, size float64) {
	DrawRect64(position.X, position.Y, domain.GameWidth-(2*position.X), domain.ButtonHeight, screen, domain.DarkGray)
	LoadText(
		label,
		entitie.Position{X: position.X + 20, Y: position.Y + 25},
		size,
		font,
		screen,
		domain.White,
		text.AlignStart,
		text.AlignCenter,
	)
}

func drawButtonMargin(
	label string,
	buttonPosition entitie.Position,
	buttonSize entitie.Size,
	textSize float64,
	font *text.GoTextFaceSource,
	screen *ebiten.Image,
) {
	DrawRect64(buttonPosition.X, buttonPosition.Y, buttonSize.Width+1, buttonSize.Height+1, screen, domain.White)
	DrawRect64(buttonPosition.X, buttonPosition.Y, buttonSize.Width, buttonSize.Height, screen, domain.DarkGray)
	LoadText(
		label,
		entitie.Position{X: buttonPosition.X + 20, Y: buttonPosition.Y + 25},
		textSize,
		font,
		screen,
		domain.White,
		text.AlignStart,
		text.AlignCenter,
	)
}
