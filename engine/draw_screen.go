package engine

import (
	"fmt"
	"mini-game-go/domain/entitie"
	"mini-game-go/helpers"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawRect64(x, y, w, h float64, screen *ebiten.Image, color string) error {
	return DrawRect32(float32(x), float32(y), float32(w), float32(h), screen, color)
}

func DrawRect32(x, y, w, h float32, screen *ebiten.Image, color string) error {
	recColor, err := helpers.HexToRGBA(color)
	if err != nil {
		return fmt.Errorf("Error to draw rect on screen: %w", err)
	}
	vector.DrawFilledRect(screen, x, y, w, h, recColor, false)
	return nil
}

func LoadText(
	textDraw string,
	position entitie.Position,
	size float64,
	font *text.GoTextFaceSource,
	screen *ebiten.Image,
	color string,
	align text.Align,
	alignV text.Align,
) error {
	// define default to texts of game
	op := &text.DrawOptions{}
	op.GeoM.Translate(position.X, position.Y)
	colorRgba, err := helpers.HexToRGBA(color)
	if err != nil {
		return fmt.Errorf("Error to load text on screen: %w", err)
	}
	op.ColorScale.Scale(
		float32(colorRgba.R)/255,
		float32(colorRgba.G)/255,
		float32(colorRgba.B)/255,
		float32(colorRgba.A)/255,
	)
	op.LineSpacing = 30
	op.PrimaryAlign = align
	op.SecondaryAlign = alignV
	text.Draw(screen, textDraw, &text.GoTextFace{
		Source: font,
		Size:   size,
	}, op)
	return nil
}
