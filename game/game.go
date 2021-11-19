package game

import (
	"fmt"

	"github.com/RyoJerryYu/gogoengine/model"
)

type Game interface {
	Initialize() error
	CleanUp() error
	ProcessPlaceOn(model.Point) error

	IsFinished() bool
	GetTurnPlayer() model.Player
	GetBoard() model.Board
	GetNewBoardWithSameSize() model.Board
	GetPassedPlayers() []model.Player
	GetPlayerMapByStoneType() map[model.StoneType]model.Player
}

type Initialize func(*game) error
type CleanUp func(*game) error
type MainPhase func(*game, model.Player, model.Point) error

type game struct {
	board      model.Board
	playerSet  []model.Player
	turnOf     uint
	initialize Initialize
	cleanUp    CleanUp
	mainPhase  MainPhase
}

var _ Game = (*game)(nil)

func (g *game) Initialize() error {
	return g.initialize(g)
}

func (g *game) CleanUp() error {
	return g.cleanUp(g)
}

func (g *game) ProcessPlaceOn(p model.Point) error {
	if err := g.mainPhase(g, g.GetTurnPlayer(), p); err != nil {
		return err
	}
	if err := g.toNextTurn(); err != nil {
		return err
	}
	return nil
}

func (g *game) toNextTurn() error {
	n := len(g.playerSet)
	if n <= 0 {
		return fmt.Errorf("no players in game")
	}
	g.turnOf = (g.turnOf + 1) % uint(n)
	return nil
}

func (g game) IsFinished() bool {
	for _, player := range g.playerSet {
		if !player.Passed() {
			return false
		}
	}
	return true
}

func (g *game) GetTurnPlayer() model.Player {
	return g.playerSet[g.turnOf]
}

func (g game) GetBoard() model.Board {
	return g.board
}

func (g game) GetNewBoardWithSameSize() model.Board {
	x, y := g.board.Size()
	return model.NewBoard(
		model.WithSize(x, y),
	)
}

func (g game) GetPassedPlayers() []model.Player {
	passedPlayers := make([]model.Player, 0, len(g.playerSet))
	for _, pl := range g.playerSet {
		if pl.Passed() {
			passedPlayers = append(passedPlayers, pl)
		}
	}
	return passedPlayers
}

func (g game) GetPlayerMapByStoneType() map[model.StoneType]model.Player {
	playerMap := make(map[model.StoneType]model.Player, len(g.playerSet))
	for _, pl := range g.playerSet {
		playerMap[pl.GetStoneType()] = pl
	}
	return playerMap
}
