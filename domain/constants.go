package domain

var GameWidth float64 = 750
var GameHeight float64 = 1000
var LineWidth float64 = 10
var LineHeight float64 = 200
var ColorGreen string = "#07B51E"
var ColorYellow string = "B57E07"
var ColorRed string = "#B50707"
var ColorWhite string = "#FFFFFF"
var ColorDarkGray string = "#0F0F0F"
var BacgoundColor string = "#3C3C3C"

func GameWidthInt() int {
	return int(GameWidth)
}
func GameHeightInt() int {
	return int(GameHeight)
}
