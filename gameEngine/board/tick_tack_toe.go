package board

import (
	"github.com/gaurav2721/low_level_design/gameEngine/move"
	"github.com/gaurav2721/low_level_design/gameEngine/players"
)

type TickTackToe struct {
	players    []players.IPlayer
	isStarted  bool
	nextPlayer players.IPlayer
}

func (t *TickTackToe) RegisterPlayers(players []players.IPlayer) (bool, error) {
	return true, nil
}
func (t *TickTackToe) Start() (bool, error) {
	return true, nil
}
func (t *TickTackToe) IsStarted() (bool, error) {
	return true, nil
}
func (t *TickTackToe) IsRegistered(player players.IPlayer) (bool, error) {
	return true, nil
}
func (t *TickTackToe) IsFinished() (bool, error) {
	return true, nil
}
func (t *TickTackToe) Move(player players.IPlayer, m move.Move) (bool, error) {
	return true, nil
}
