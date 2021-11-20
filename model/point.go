package model

type Point interface {
	GetXY() (x, y uint32)
}

type point struct {
	x uint32
	y uint32
}

var _ Point = (*point)(nil)

func (p point) GetXY() (x, y uint32) {
	return p.x, p.y
}
