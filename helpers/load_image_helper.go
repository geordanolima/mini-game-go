package helpers

import (
	"image"
	"os"
	"path/filepath"
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
