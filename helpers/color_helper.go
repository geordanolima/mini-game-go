package helpers

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

func HexToRGBA(hex string, opacity ...uint8) (color.RGBA, error) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return color.RGBA{}, fmt.Errorf("invalid hexadecimal color format, must be #RRGGBB")
	}
	red, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}
	green, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}
	blue, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}
	var alpha uint8 = 255
	if len(opacity) > 0 {
		alpha = opacity[0]
	}
	return color.RGBA{uint8(red), uint8(green), uint8(blue), alpha}, nil
}
