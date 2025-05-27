package board

import (
	"github.com/gaurav2721/low_level_design/gameEngine/move"
	"github.com/gaurav2721/low_level_design/gameEngine/players"
)

type TickTackToe struct {
	players    []players.Player
	isStarted  bool
	nextPlayer players.Player
}

func (t *TickTackToe) RegisterPlayer(player []players.Player) ([]bool, []error) {}
func (t *TickTackToe) Start() (bool, error)                                     {}
func (t *TickTackToe) IsStarted() (bool, error)                                 {}
func (t *TickTackToe) IsRegistered(player []players.Player) ([]bool, []error)   {}
func (t *TickTackToe) IsFinished() (bool, error)                                {}
func (t *TickTackToe) Move(player players.Player, move move.Move) (bool, error) {}
