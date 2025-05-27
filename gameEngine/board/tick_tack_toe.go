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

func (t *TickTackToe) RegisterPlayers(players []players.Player) (bool, error) {
	return true, nil
}
func (t *TickTackToe) Start() (bool, error) {
	return true, nil
}
func (t *TickTackToe) IsStarted() (bool, error) {
	return true, nil
}
func (t *TickTackToe) IsRegistered(player []players.Player) ([]bool, []error) {
	return []bool{true}, []error{nil}
}
func (t *TickTackToe) IsFinished() (bool, error) {
	return true, nil
}
func (t *TickTackToe) Move(player players.Player, move move.Move) (bool, error) {
	return true, nil
}
