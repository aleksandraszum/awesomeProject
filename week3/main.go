package main

import (
	"fmt"
	"math/rand"
	"time"
)

// this program allows you to play with computer in the game: rock, paper and scissors

func main() {

	// declare the variales
	var rounds, machineMove int
	var myMove rune
	var mMove string
	var draws, myWin, computerWin int

	// decide how many rounds do you want to play
	fmt.Println("How many rounds do you want to play?")
	fmt.Scanf("%d", &rounds)

	for i := 0; i < rounds; i++ {

		// your move
		fmt.Println("\nRounds ", i+1, ": Choose either R, P or S")
		fmt.Scanf("%c", &myMove)

		if myMove == 'R' {
			fmt.Println("You move: rock")
		} else if myMove == 'P' {
			fmt.Println("You move: paper")
		} else if myMove == 'S' {
			fmt.Println("You move: scissors")
		} else {
			fmt.Println("Illegal move ")
			i--
			continue
		}

		// computer move using the random number generator
		if (myMove == 'R') || (myMove == 'P') || (myMove == 'S') {
			rand.New(rand.NewSource(time.Now().UnixNano()))
			machineMove = rand.Intn(3)

			if machineMove == 0 {
				mMove = "Rock"
			} else if machineMove == 1 {
				mMove = "Paper"
			} else {
				mMove = "Scissors"
			}
			fmt.Println("Machine plays ", mMove)

			// decision who wins current round
			concatenated := fmt.Sprintf("%c%d", myMove, machineMove)
			//fmt.Println(concatenated)

			if (concatenated == "R0") || (concatenated == "P1") || (concatenated == "S2") {
				draws = draws + 1
				fmt.Println("Draw")
				continue
			} else if (concatenated == "R2") || (concatenated == "P0") || (concatenated == "S1") {
				myWin = myWin + 1
				fmt.Println("You win")
				continue
			} else {
				computerWin = computerWin + 1
				fmt.Println("Computer wins")
				continue
			}
		}
	}

	//final grade:
	fmt.Println("You win:", myWin, "times, computer wins: ", computerWin, " times. Draws are: ", draws, "times.")
}
