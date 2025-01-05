package helpers

import (
	"bytes"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func LoadFont(fileName string) (*text.GoTextFaceSource, error) {
	fontBytes, err := ioutil.ReadFile(filepath.Join("assets", "fonts", fileName))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	myFont, err := text.NewGoTextFaceSource(bytes.NewReader(fontBytes))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return myFont, nil
}
