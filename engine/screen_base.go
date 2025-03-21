package engine

import (
	"mini-game-go/domain"
	"mini-game-go/domain/entitie"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func titleSize(title, subtitle string, font *text.GoTextFaceSource, screen *ebiten.Image, titleSize, subtitleSize float64) {
	posX := (domain.GameWidth / 2)
	LoadText(
		title, entitie.Position{X: posX, Y: 50}, titleSize, font, screen, domain.White, text.AlignCenter, text.AlignCenter,
	)
	LoadText(
		subtitle, entitie.Position{X: posX, Y: 100}, subtitleSize, font, screen, domain.White, text.AlignCenter, text.AlignCenter,
	)
}

func title(title, subtitle string, font *text.GoTextFaceSource, screen *ebiten.Image) {
	titleSize(title, subtitle, font, screen, 30, 25)
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
	drawButtonMarginColor(
		label,
		buttonPosition,
		buttonSize,
		textSize,
		font,
		screen,
		domain.White,
		domain.White,
		domain.DarkGray,
		text.AlignStart,
		text.AlignCenter,
	)
}

func drawButtonMarginColor(
	label string,
	buttonPosition entitie.Position,
	buttonSize entitie.Size,
	textSize float64,
	font *text.GoTextFaceSource,
	screen *ebiten.Image,
	fontColor, buttonMarginColor, buttonColor string,
	align, alignV text.Align,
) {
	DrawRect64(buttonPosition.X, buttonPosition.Y, buttonSize.Width+1, buttonSize.Height+1, screen, buttonMarginColor)
	DrawRect64(buttonPosition.X, buttonPosition.Y, buttonSize.Width, buttonSize.Height, screen, buttonColor)
	LoadText(
		label,
		entitie.Position{X: buttonPosition.X + 20, Y: buttonPosition.Y + 25},
		textSize,
		font,
		screen,
		fontColor,
		align,
		alignV,
	)
}

func backButon(game *Game, screen *ebiten.Image) {
	drawButtonMargin(
		"Back",
		entitie.Position{X: 20, Y: domain.GameHeight - domain.ButtonHeight - 20},
		entitie.Size{Width: domain.ButtonWidth, Height: domain.ButtonHeight},
		30,
		game.Font,
		screen,
	)
}

func confirmButon(game *Game, screen *ebiten.Image) {
	drawButtonMargin(
		"Confirm",
		entitie.Position{X: domain.GameWidth - domain.ButtonWidth - 20, Y: domain.GameHeight - domain.ButtonHeight - 20},
		entitie.Size{Width: domain.ButtonWidth, Height: domain.ButtonHeight},
		30,
		game.Font,
		screen,
	)
}
