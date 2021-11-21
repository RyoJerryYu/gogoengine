package game

import "github.com/RyoJerryYu/gogoengine/model"

func WithSize(x, y uint32) gameOption {
	return func(g *game) {
		g.board = model.NewBoard(
			model.WithSize(x, y),
		)
	}
}

func WithPlayers(pls []model.Player) gameOption {
	return func(g *game) {
		g.playerSet = pls
	}
}

func WithInitialize(i Initialize) gameOption {
	return func(g *game) {
		g.initialize = i
	}
}

func WithMainPhase(m MainPhase) gameOption {
	return func(g *game) {
		g.mainPhase = m
	}
}

func WithCleanUp(c CleanUp) gameOption {
	return func(g *game) {
		g.cleanUp = c
	}
}
