package game

import (
	"fmt"
	"testing"

	"github.com/RyoJerryYu/gogoengine/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func emptyInitialize(g *game) error {
	return nil
}

func emptyCleanUp(g *game) error {
	return nil
}

func emptyMainPhase(g *game, pl model.Player, p model.Point) error {
	return nil
}

func TestGame(t *testing.T) {
	var (
		g   Game
		err error
	)

	_, err = NewGame()
	require.Error(t, err)

	_, err = NewGame(
		WithPlayers([]model.Player{
			model.NewPlayer(model.StoneType_Empty, "player1"),
			model.NewPlayer(model.StoneType_Empty, "player2"),
		}),
	)
	require.Error(t, err)

	_, err = NewGame()
	require.Error(t, err)

	g, err = NewGame(
		WithPlayers([]model.Player{
			model.NewPlayer(model.StoneType_Empty, "player1"),
			model.NewPlayer(model.StoneType_Empty, "player2"),
		}),
		WithSize(1, 1),
		WithCleanUp(emptyCleanUp),
		WithInitialize(emptyInitialize),
		WithMainPhase(emptyMainPhase),
	)
	require.NoError(t, err)
	assert.NotNil(t, g)
}

func TestGame_Methods(t *testing.T) {
	var (
		g        Game
		err      error
		errInit  = fmt.Errorf("initialize error")
		errClean = fmt.Errorf("cleanup error")
	)

	g, err = NewGame(
		WithPlayers([]model.Player{
			model.NewPlayer(model.StoneType_Empty, "player1"),
			model.NewPlayer(model.StoneType_Empty, "player2"),
		}),
		WithSize(1, 1),
		WithCleanUp(func(g *game) error { return errClean }),
		WithInitialize(func(g *game) error { return errInit }),
		WithMainPhase(emptyMainPhase),
	)
	require.NoError(t, err)
	assert.NotNil(t, g)

	err = g.Initialize()
	require.ErrorIs(t, err, errInit)
	err = g.CleanUp()
	require.ErrorIs(t, err, errClean)
}
