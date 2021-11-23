package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBoard(t *testing.T) {
	b := NewBoard(
		WithSize(3, 3),
		WithInitValue(StoneType(1)),
	)

	sizeX, sizeY := b.Size()
	assert.Equal(t, uint32(3), sizeX)
	assert.Equal(t, uint32(3), sizeY)

	for xi := 0; xi < 3; xi++ {
		for yi := 0; yi < 3; yi++ {
			stone, err := b.Get(NewPoint(uint32(xi), uint32(yi)))
			require.NoError(t, err)
			assert.Equal(t, StoneType(1), stone)
		}
	}

	err := b.Set(NewPoint(uint32(1), uint32(2)), StoneType(2))
	require.NoError(t, err)
	stone, err := b.Get(NewPoint(uint32(1), uint32(2)))
	require.NoError(t, err)
	assert.Equal(t, StoneType(2), stone)

	err = b.Set(NewPoint(uint32(2), uint32(3)), StoneType(1))
	require.Error(t, err)
	_, err = b.Get(NewPoint(uint32(3), uint32(1)))
	require.Error(t, err)
}
