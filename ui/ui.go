package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RyoJerryYu/gogoengine/model"
)

type UserInterface interface {
	PrintBoard(model.Board) error // TODO: handle error while return
	PrintPlayerMessages(turnPlayer model.Player, passedPlayers []model.Player)
	GetInput(model.Player) (model.Point, error)
}

type userInterface struct {
	renderMap StoneRenderMap
	grid      displayGrid
}

var _ UserInterface = (*userInterface)(nil)

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

func (ui *userInterface) PrintPlayerMessages(
	turnPlayer model.Player,
	passedPlayers []model.Player,
) {
	passedPlayerNames := make([]string, 0, len(passedPlayers))
	for _, paspassedPlayer := range passedPlayers {
		passedPlayerNames = append(passedPlayerNames, paspassedPlayer.GetTag())
	}
	fmt.Printf("Passed players: %s", strings.Join(passedPlayerNames, ", "))

	if turnPlayer.Passed() {
		fmt.Printf("Player: %s was passed!", turnPlayer.GetTag())
	} else {
		fmt.Printf("Turn of: %s", turnPlayer.GetTag())
	}
}

func (ui *userInterface) GetInput(turnPlayer model.Player) (model.Point, error) {
	var point model.Point
	fmt.Printf("---")
	fmt.Printf("Please input a position as <row col>")
	fmt.Printf("e.g. 3 4")

	var input string
	fmt.Scanln(&input)

	inputs := strings.Split(strings.TrimSpace(input), " ")
	if len(inputs) <= 0 {
		return nil, fmt.Errorf("<%s> not a position", input)
	}

	if len(inputs) <= 1 {
		return nil, fmt.Errorf("value less than one")
	}

	x, err := strconv.Atoi(inputs[0])
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(inputs[1])
	if err != nil {
		return nil, err
	}
	// TODO: handle check position

	point = model.NewPoint(uint32(x), uint32(y))

	return point, nil
}

type uiOption func(*userInterface)

type displayGrid [3][3]string
type StoneRenderMap map[model.StoneType]string

func NewUI(opts ...uiOption) UserInterface {
	u := &userInterface{}

	for _, opt := range opts {
		opt(u)
	}

	return u
}

func WithDefaultGrid(ui *userInterface) {
	ui.grid = [3][3]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}
}

func WithRenderMap(renderMap StoneRenderMap) uiOption {
	return func(ui *userInterface) {
		ui.renderMap = renderMap
	}
}
