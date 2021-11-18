package engine

import (
	"github.com/RyoJerryYu/gogoengine/game"
	"github.com/RyoJerryYu/gogoengine/ui"
)

func SetGame(game game.Game) engineOption {
	return func(e *engine) {
		e.game = game
	}
}

func SetUI(ui ui.UserInterface) engineOption {
	return func(e *engine) {
		e.ui = ui
	}
}
