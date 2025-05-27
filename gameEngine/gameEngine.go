package gameEngine

import (
	"fmt"

	"github.com/gaurav2721/low_level_design/gameEngine/board"
	"github.com/gaurav2721/low_level_design/gameEngine/constants"
)

type GameEngine interface {
	GetBoard(game string) (board.Board, error)
}

type gameEngineChild struct {
}

func (ge gameEngineChild) GetBoard(game string) (board.Board, error) {
	if game == constants.TickTackToe {
		return &board.TickTackToe{}, nil
	} else {
		return nil, fmt.Errorf("this game is not available in game engine")
	}
}

func Initialize(game string) (board.Board, error) {
	return gameEngineChild{}.GetBoard(game)
}
