package game

import (
	"github.com/RyoJerryYu/gogoengine/model"
)

type Initialize func(Game) error
type Terminate func(Game) error
type MainPhase func(Game, model.Player, model.Point) error

type Game interface {
	Initialize() error
	CleanUp() error
	ProcessPlaceOn(model.Point) error

	IsFinished() bool
	GetTurnPlayer() model.Player
	GetBoard() model.Board
	GetPassedPlayers() []model.Player
}
