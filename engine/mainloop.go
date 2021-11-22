package engine

import "github.com/RyoJerryYu/gogoengine/model"

func (e engine) Run() error {
	if err := e.initialize(); err != nil {
		return err
	}

	for player, isFinished := e.nextTurn(); !isFinished; player, isFinished = e.nextTurn() {
		e.displayPreTurn(player)
		if err := e.activeMainPhase(player); err != nil {
			return err
		}
	}

	if err := e.cleanUp(); err != nil {
		return err
	}

	return nil
}

func (e engine) initialize() error {
	return e.game.Initialize()
}

func (e engine) nextTurn() (model.Player, bool) {
	player := e.game.GetTurnPlayer()
	isFinished := e.game.IsFinished()
	return player, isFinished
}

func (e engine) displayPreTurn(turnPlayer model.Player) {
	board := e.game.GetBoard()
	passedPlayers := e.game.GetPassedPlayers()
	for {
		err := e.ui.PrintBoard(board)
		if err != nil {
			continue
		}
		break
	}
	e.ui.PrintPlayerMessages(turnPlayer, passedPlayers)
}

func (e engine) activeMainPhase(turnPlayer model.Player) error {
	point, err := e.ui.GetInput(turnPlayer)
	if err != nil {
		return err
	}
	err = e.game.ProcessPlaceOn(point)
	if err != nil {
		return err
	}
	return nil
}

func (e engine) cleanUp() error {
	return e.game.CleanUp()
}
