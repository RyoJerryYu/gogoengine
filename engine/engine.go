package engine

import (
	"github.com/RyoJerryYu/gogoengine/game"
	"github.com/RyoJerryYu/gogoengine/ui"
)

type Engine interface {
	Run() error
}

type engine struct {
	game game.Game
	ui   ui.UserInterface
}

var _ Engine = (*engine)(nil)

type engineOption func(*engine)

func NewEngine(opts ...engineOption) Engine {
	e := &engine{}

	for _, opt := range opts {
		opt(e)
	}
	return e
}
