package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPoint(t *testing.T) {
	p := NewPoint(1, 2)
	x, y := p.GetXY()
	assert.Equal(t, uint32(1), x)
	assert.Equal(t, uint32(2), y)
}
