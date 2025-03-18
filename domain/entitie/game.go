package entitie

import "time"

type GameState int

const (
	StateMenu GameState = iota
	StateNewGame
	StateGameRunning
	StateEnterName
	StateRecords
	StateControls
)

type Menu struct {
	Actions []Action
}

type User struct {
	Name        string
	InputActive bool
}

type Score struct {
	Score int
	Time  time.Time
}

type GameOver struct {
	Flag        bool
	BoxObject   Object
	TextOptions TextOptions
}

type Action struct {
	TextOptions TextOptions
	Object      Object
	Visible     bool
	State       GameState
}
