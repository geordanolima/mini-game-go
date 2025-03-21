package entitie

import "time"

type GameState int

const (
	StateMenu GameState = iota
	StateNewGame
	StateDifficulty
	StateGameRunning
	StateEnterName
	StateRecords
	StateControls
)

type Selector struct {
	Actions []Action
	Active  bool
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
	State       int
	Active      bool
}
