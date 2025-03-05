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
var PositionsX = []int{15, 155, 305, 455, 605}
var PercentsGas = []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

func GameWidthInt() int {
	return int(GameWidth)
}
func GameHeightInt() int {
	return int(GameHeight)
}

var ObstacleImages []Obstacle = []Obstacle{
	{FilePath: "cone.png", Object: Object{Size: Size{Width: 100, Height: 100}, Position: Position{X: 20}, Margin: -20}},
	{FilePath: "cone2.png", Object: Object{Size: Size{Width: 125, Height: 125}, Position: Position{X: 10}, Margin: -35}},
	{FilePath: "hole.png", Object: Object{Size: Size{Width: 100, Height: 100}, Position: Position{X: 25}, Margin: -30}},
	{FilePath: "truck.png", Object: Object{Size: Size{Width: 160, Height: 330}, Position: Position{X: 5}, Margin: -10}},
	{FilePath: "bus.png", Object: Object{Size: Size{Width: 150, Height: 430}, Position: Position{X: 0}, Margin: -10}},
}
