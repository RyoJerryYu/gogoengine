package engine

import "github.com/RyoJerryYu/gogoengine/model"

type engine struct {
	players []model.Player

	inits      []Initialize
	termins    []Terminate
	mainphases []MainPhase
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

func (e engine) Run() error {
	for _, initialize := range e.inits {
		err := initialize(e)
		if err != nil {
			return err
		}
	}

	for !e.allPlayersPassed() {
		for _, player := range e.players {
			var activeP model.Point
			if !player.Passed() {
				e.display()
				activeP = e.input(player)
			}
			if player.Passed() {
				e.showPassInfo()
				continue
			}
			for _, mainphase := range e.mainphases {
				err := mainphase(e, player, activeP)
				if err != nil {
					return err
				}
			}
		}

	}

	for _, terminate := range e.termins {
		err := terminate(e)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e engine) allPlayersPassed() bool {
	return true
}

func (e engine) display() {

}

func (e engine) input(player model.Player) model.Point {
	return nil
}

func (e engine) showPassInfo() {}
