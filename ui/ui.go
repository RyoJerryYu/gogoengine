package ui

import "github.com/RyoJerryYu/gogoengine/model"

type UserInterface interface {
	PrintBoard(model.Board)
	PrintPlayerMessages(turnPlayer model.Player, passedPlayers []model.Player)
	GetInput(model.Player) (model.Point, error)
}
