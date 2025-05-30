package aiengine

import (
	"github.com/gaurav2721/low_level_design/gameEngine/move"
	"github.com/gaurav2721/low_level_design/gameEngine/players"
)

type IAiEngine interface {
	SuggestMove(p players.IPlayer) (move.Move, error)
}

func InitializeAiEngine() IAiEngine {
	return AiEngine{}
}
