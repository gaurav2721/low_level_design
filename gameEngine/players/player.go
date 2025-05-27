package players

import (
	"github.com/gaurav2721/low_level_design/gameEngine/board"
	"github.com/gaurav2721/low_level_design/gameEngine/move"
)

type Player struct {
	name string
	id   string
}

func (p Player) Move(board board.Board, move move.Move) (bool, error) {
	board.Move(p, move)
}
