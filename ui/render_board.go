package ui

import (
	"strconv"

	"github.com/RyoJerryYu/gogoengine/model"
)

type boardRenderer func(uint32, uint32, model.Board) ([][]string, error)

func (ui *userInterface) renderBoard(
	sizeX, sizeY uint32,
	board model.Board,
) ([][]string, error) {

	rendered := make([][]string, 0, sizeX)
	for x := uint32(0); x < sizeX; x++ {

		line := make([]string, 0, sizeY)
		for y := uint32(0); y < sizeY; y++ {

			stone, err := board.Get(model.NewPoint(x, y))
			if err != nil {
				return nil, err
			}

			renderedCell := ui.renderCell(x, y, sizeX, sizeY, stone)
			line = append(line, renderedCell)
		}

		rendered = append(rendered, line)
	}

	return rendered, nil
}

func (ui *userInterface) renderCell(
	x, y uint32,
	sizeX, sizeY uint32,
	stone model.StoneType,
) string {
	// render stone
	if stone != model.StoneType_Empty {
		if cell, ok := ui.stoneCellMap[stone]; ok {
			return cell
		}
	}

	// render grid
	var gridX int
	var gridY int

	switch x {
	case 0:
		gridX = 0
	case sizeX - 1:
		gridX = 2
	default:
		gridX = 1
	}

	switch y {
	case 0:
		gridY = 0
	case sizeY - 1:
		gridY = 2
	default:
		gridY = 1
	}

	return ui.grid[gridX][gridY]
}

func renderWithLabel(renderer boardRenderer) boardRenderer {
	return func(sizeX, sizeY uint32, b model.Board) ([][]string, error) {
		rendered, err := renderer(sizeX, sizeY, b)
		if err != nil {
			return rendered, err
		}
		renderedSizeX := len(rendered)
		if renderedSizeX <= 0 {
			return rendered, nil
		}
		renderedSizeY := len(rendered[0])

		rendered = renderWithXLabel(renderedSizeX, rendered)
		rendered = renderWithYLabel(renderedSizeY, rendered)
		return rendered, nil
	}
}

func renderWithXLabel(renderedSizeX int, rendered [][]string) [][]string {
	digitsX := len(strconv.Itoa(renderedSizeX - 1))
	for xi := 0; xi < renderedSizeX; xi++ {
		xLabel := intToDigitsLengthChars(xi, digitsX)
		rendered[xi] = append(rendered[xi], xLabel...)
	}
	return rendered
}

func renderWithYLabel(renderedSizeY int, rendered [][]string) [][]string {
	digitsY := len(strconv.Itoa(renderedSizeY - 1))
	yLabels := make([][]string, 0, renderedSizeY)
	for i := 0; i < renderedSizeY; i++ {
		yLabels = append(yLabels, intToDigitsLengthChars(i, digitsY))
	}

	for i := 0; i < digitsY; i++ {
		line := make([]string, 0, renderedSizeY)
		for _, labels := range yLabels {
			line = append(line, labels[i])
		}
		rendered = append(rendered, line)
	}
	return rendered
}

func intToDigitsLengthChars(v int, digits int) []string {
	res := make([]string, digits) // digits of 0
	for i := 0; i < digits; i++ {
		res[len(res)-i-1] = strconv.Itoa(v % 10)
		v = v / 10
	}
	return res
}
