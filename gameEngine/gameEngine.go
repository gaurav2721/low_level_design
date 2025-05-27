package gameEngine

import "github.com/gaurav2721/low_level_design/gameEngine/board"

type GameEngine interface {
	GetGameObj(game string) board.Board
}

type GameEngineChild struct {
}

func (ge GameEngineChild) GetGameObj(game string) board.Board {

}
