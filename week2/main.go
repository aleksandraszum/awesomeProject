package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

// this program calculate the probability than two people in group of [10, 20, ... 100] persons have a birtday at the same day

func main() {

	fmt.Println("People in the room       Probability of 2 (or more) having the same birthday")
	// iteration between the group of people
	for i := 10; i <= 100; i += 10 { // the group of 10, 20, 30, ... 100 people

		var howMany int = 0 // how many pairs are there
		// 10000 simulation for each group of 10, 20 , ... people
		// there are contain howMany one pair have birthay at the same day
		for j := 0; j < 10000; j++ {
			// create an empty slice containing the date of birth
			var birthdaySlice []int
			// condition for breaking for loop

			// random birthdates - count of dates depends on count of people in one room (the value of "i" variable)
			for k := 0; k < i; k++ {
				birthdaySlice = append(birthdaySlice, rand.Intn(365))
			}
			// check if two people have a birthday at the same day
			for j := 0; j < i; j++ {
				removeIndex := RemoveIndex(birthdaySlice, j)
				containsTen := slices.Contains(removeIndex, birthdaySlice[j])
				if containsTen {
					howMany++
					break
				}
			}
		}

		// calculate the probability
		prob := float32(howMany) / 10000
		fmt.Println(i, "                     ", prob)

	}
}
