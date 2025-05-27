package utils

import (
	"fmt"

	"github.com/gaurav2721/low_level_design/gameEngine/move"
)

func askHumanForMove() (move.Move, error) {
	r, c := 0, 0
	fmt.Scan(&r, &c)
	return move.Move{
		Row: r,
		Col: c,
	}, nil
}
