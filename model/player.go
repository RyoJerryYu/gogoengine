package model

type Player interface {
	Passed() bool
	GetStoneType() StoneType
	GetTag() string
}

type player struct {
	stoneType StoneType
	passed    bool
	tag       string
}

var _ Player = (*player)(nil)

func (p player) Passed() bool {
	return p.passed
}

func (p player) GetStoneType() StoneType {
	return p.stoneType
}

func (p player) GetTag() string {
	return p.tag
}

func NewPlayer(stoneType StoneType, tag string) Player {
	return player{
		stoneType: stoneType,
		passed:    false,
		tag:       tag,
	}
}
