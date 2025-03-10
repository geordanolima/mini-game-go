package helpers

import (
	"image"
	"os"
	"path/filepath"

	"golang.org/x/image/draw"
)

func LoadImage(fileName string) (image.Image, error) {
	// get file image
	file, err := os.Open(filepath.Join("assets", fileName))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// decode image
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func LoadImageResize(fileName string, width, height float64) (image.Image, error) {
	img, err := LoadImage(fileName)
	if err != nil {
		return nil, err
	}
	rect := image.Rect(0, 0, int(width), int(height))
	res := image.NewRGBA(rect)
	draw.NearestNeighbor.Scale(res, rect, img, img.Bounds(), draw.Over, nil)
	return res, nil
}
