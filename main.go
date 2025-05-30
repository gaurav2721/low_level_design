package main

import (
	"fmt"

	gameEngine "github.com/gaurav2721/low_level_design/GameEngine"
	"github.com/gaurav2721/low_level_design/gameEngine/constants"
	"github.com/gaurav2721/low_level_design/gameEngine/players"
)

func main() {
	fmt.Println("main started")

	// initialize the ai engine

	// initialize the board
	board, err := gameEngine.Initialize(constants.TickTackToe)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	playersArr := []players.IPlayer{players.InitializePlayer("computer", 'X'), players.InitializePlayer("human", '0')}

	// registering the players
	_, err = board.RegisterPlayers(playersArr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//start the game
	_, err = board.Start()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for {

		m, err := utils.askHumanForMove()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		_, err = board.Move(human, m)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		isFinished, err = board.IsFinished()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if isFinished {
			break
		}

		m, err = board.GetNextAutoMove()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		_, err = board.Move(human, m)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

	}

	winner, err := board.GetWinner()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("winner", winner)

}
