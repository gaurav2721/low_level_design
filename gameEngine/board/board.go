package board

import (
	"github.com/gaurav2721/low_level_design/gameEngine/move"
	"github.com/gaurav2721/low_level_design/gameEngine/players"
)

type Board interface {
	RegisterPlayer(player []players.Player) ([]bool, []error)
	Start() (bool, error)
	IsStarted() (bool, error)
	IsRegistered(player []players.Player) ([]bool, []error)
	IsFinished() (bool, error)
	Move(player players.Player, move move.Move) (bool, error)
}
