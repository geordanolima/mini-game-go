package helpers

import (
	"image"

	"golang.org/x/image/draw"
)

func resize(img image.Image, newWidth, newHeight int) image.Image {
	rect := image.Rect(0, 0, newWidth, newHeight)
	res := image.NewRGBA(rect)
	draw.NearestNeighbor.Scale(res, rect, img, img.Bounds(), draw.Over, nil)
	return res
}
