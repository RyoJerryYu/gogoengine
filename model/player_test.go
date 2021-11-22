package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlayer(t *testing.T) {
	pl := NewPlayer(StoneType(1), "test pl")
	assert.Equal(t, "test pl", pl.GetTag())
	assert.Equal(t, StoneType(1), pl.GetStoneType())
}
