package ui

import (
	"fmt"
	"strings"

	"github.com/RyoJerryYu/gogoengine/model"
)

type UserInterface interface {
	PrintBoard(model.Board) error // TODO: handle error while return
	PrintPlayerMessages(turnPlayer model.Player, passedPlayers []model.Player)
	GetInput(model.Player) (model.Point, error)
}

type userInterface struct {
	stoneCellMap map[model.StoneType]string
	grid         displayGrid
}

type displayGrid [3][3]string

func defaultGrid() displayGrid {
	return [3][3]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}
}

func (ui *userInterface) PrintBoard(board model.Board) error {
	sizeX, sizeY := board.Size()
	renderer := ui.renderBoard
	renderer = renderWithLabel(renderer)
	renderedBoard, err := renderer(sizeX, sizeY, board)
	if err != nil {
		return err
	}

	printRendered(renderedBoard)
	return nil
}

func printRendered(rendered [][]string) {
	for _, line := range rendered {
		fmt.Println(strings.Join(line, ""))
	}
}
