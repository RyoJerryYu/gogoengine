package engine

import (
	"github.com/RyoJerryYu/gogoengine/model"
)

type Engine interface {
	Run() error
}

type Initialize func(Engine) error
type Terminate func(Engine) error
type MainPhase func(Engine, model.Player, model.Point) error
