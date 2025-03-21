package entitie

type Position struct {
	X float64
	Y float64
}

type Size struct {
	Width  float64
	Height float64
}

type Object struct {
	Size     Size
	Position Position
	Margin   float64
}
