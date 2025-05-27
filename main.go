package main

import (
	"fmt"

	gameEngine "github.com/gaurav2721/low_level_design/GameEngine"
	"github.com/gaurav2721/low_level_design/gameEngine/constants"
	"github.com/gaurav2721/low_level_design/gameEngine/players"
)

func main() {
	fmt.Println("main started")

	// initialize the board
	board, err := gameEngine.Initialize(constants.TickTackToe)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// initialize the players
	computer := players.Player{
		Name:   "computer",
		Marker: 'X',
	}
	human := players.Player{
		Name:   "human",
		Marker: '0',
	}
	playersArr := []players.Player{human, computer}

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
		isFinished, err := board.IsFinished()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if isFinished {
			break
		}

		m, err := utils.askHumanForMove()

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
