package board

import (
	"github.com/gaurav2721/low_level_design/gameEngine/move"
	"github.com/gaurav2721/low_level_design/gameEngine/players"
)

type IBoard interface {
	RegisterPlayers(playerArr []players.IPlayer) (bool, error)
	Start() (bool, error)
	IsStarted() (bool, error)
	IsRegistered(player players.IPlayer) (bool, error)
	IsFinished() (bool, error)
	Move(player players.IPlayer, m move.Move) (bool, error)
}
