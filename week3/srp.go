package main

import (
	"fmt"
	"strings"
)

// this program allows you to play with computer in the game: rock, paper and scissors
// there are 100 round.
// computer "remembers" your goals and try to choose the correct move

func main() {

	// declare the variales
	var myMove string
	var machineMove string
	var draws, myWin, computerWin int
	var papers, rocks, scissors int
	papers = 0
	rocks = 0
	scissors = 0
	moveTable := [3]string{"R", "S", "P"}

	for i := 0; i < 100; i++ {

		// your move
		fmt.Println("\nRounds ", i+1, ": Choose either R, P or S")
		fmt.Scan(&myMove)

		// your move:
		if strings.ToUpper(myMove) == "R" {
			fmt.Println("You move: rock")
			rocks = rocks + 1
		} else if strings.ToUpper(myMove) == "P" {
			fmt.Println("You move: paper")
			papers = papers + 1
		} else if strings.ToUpper(myMove) == "S" {
			fmt.Println("You move: scissors")
			scissors = scissors + 1
		} else {
			fmt.Println("Illegal move ")
			i--
			continue
		}

		// computer moves
		if i < 10 {
			j := i % 3
			machineMove = moveTable[j]
			fmt.Println("Machine move: ", machineMove)
		} else {
			if (scissors > rocks) && (scissors > papers) {
				machineMove = "R"
				fmt.Println("Machine move: rock")
			} else if (papers > scissors) && (papers > rocks) {
				machineMove = "S"
				fmt.Println("Machine moves: scissor ")
			} else {
				machineMove = "P"
				fmt.Println("Machine movea: paper")

			}
		}

		// check who wins
		if strings.ToUpper(myMove) == machineMove {
			fmt.Println("Draw")
			draws = draws + 1
		}
		if (strings.ToUpper(myMove) == "S" && machineMove == "P") || (strings.ToUpper(myMove) == "P" && machineMove == "R") || (strings.ToUpper(myMove) == "R" && machineMove == "S") {
			fmt.Println("You win")
			myWin = myWin + 1
		}
		if (machineMove == "S" && strings.ToUpper(myMove) == "P") || (machineMove == "P" && strings.ToUpper(myMove) == "R") || (machineMove == "R" && strings.ToUpper(myMove) == "S") {
			fmt.Println("Computer wins")
			computerWin = computerWin + 1
		}

	}

	//final grade:
	fmt.Println("You win:", myWin, "times, computer wins: ", computerWin, " times. Draws are: ", draws, "times.")
}
