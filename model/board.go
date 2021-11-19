package model

import (
	"fmt"
)

type StoneType int16

const StoneType_Empty StoneType = -1

type Board interface {
	Get(Point) (StoneType, error)
	Set(Point, StoneType) error
	Size() (x, y uint32)
}

type board struct {
	pool  [][]StoneType
	sizeX uint32
	sizeY uint32
}

var _ Board = (*board)(nil)

func (b board) isOutOfRange(x, y uint32) bool {
	return x >= b.sizeX || y >= b.sizeY
}

func (b board) Get(p Point) (StoneType, error) {
	x, y := p.GetXY()
	if b.isOutOfRange(x, y) {
		return StoneType_Empty, fmt.Errorf("point out of range with x(%d), y(%d)", x, y)
	}
	return b.pool[x][y], nil
}

func (b *board) Set(p Point, stone StoneType) error {
	x, y := p.GetXY()
	if b.isOutOfRange(x, y) {
		return fmt.Errorf("point out of range with x(%d), y(%d)", x, y)
	}
	b.pool[x][y] = stone
	return nil
}

func (b board) Size() (x, y uint32) {
	return b.sizeX, b.sizeY
}

func newBoard(x, y uint32, value StoneType) *board {
	const uint0 = uint32(0)

	pool := make([][]StoneType, 0, x)
	for xi := uint0; xi < x; xi++ {
		line := make([]StoneType, 0, y)
		for yi := uint0; yi < y; yi++ {
			line = append(line, value)
		}
		pool = append(pool, line)
	}
	return &board{
		pool:  pool,
		sizeX: x,
		sizeY: y,
	}
}

type boardBuilder struct {
	sizeX         uint32
	sizeY         uint32
	initiateValue StoneType
}

type boardOption func(*boardBuilder)

func NewBoard(opts ...boardOption) *board {
	builder := &boardBuilder{
		sizeX:         9,
		sizeY:         9,
		initiateValue: StoneType_Empty,
	}

	for _, opt := range opts {
		opt(builder)
	}

	return newBoard(
		builder.sizeX,
		builder.sizeY,
		builder.initiateValue,
	)
}

func WithSize(x, y uint32) boardOption {
	return func(bb *boardBuilder) {
		bb.sizeX = x
		bb.sizeY = y
	}
}

func WithInitValue(v StoneType) boardOption {
	return func(bb *boardBuilder) {
		bb.initiateValue = v
	}
}
