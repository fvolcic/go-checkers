package main

import (
	"fmt"
	"gocheckers"
)

func main() {
	board := gocheckers.NewCheckersBoard()

	for {

		if board.GetTurn() == 1 {
			fmt.Println("Black's turn")
		} else {
			fmt.Println("White's turn")
		}

		moves := board.GenerateMoves()

		for i := 0; i < len(moves); i++ {
			fmt.Printf("%d.) ", i)
			fmt.Println(moves[i])
		}

		fmt.Println(board.ToString())

		fmt.Println("\nEnter your move: ")

		var move int

		fmt.Scan(&move)

		board.MakeMove(moves[move])

	}

}
