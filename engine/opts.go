package engine

import (
	"github.com/RyoJerryYu/gogoengine/game"
	"github.com/RyoJerryYu/gogoengine/ui"
)

func WithGame(game game.Game) engineOption {
	return func(e *engine) {
		e.game = game
	}
}

func WithUI(ui ui.UserInterface) engineOption {
	return func(e *engine) {
		e.ui = ui
	}
}
