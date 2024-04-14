package main

import (
	"fmt"
	"math/rand"
)

// this program sorts the array using merge sort method

// divide the array into two group
func mergeSort(tableToSort []int) []int {
	if len(tableToSort) < 2 {
		return tableToSort
	}
	first := mergeSort(tableToSort[:len(tableToSort)/2])
	second := mergeSort(tableToSort[len(tableToSort)/2:])
	return merge(first, second)
}

// merge the data
func merge(first []int, second []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			final = append(final, first[i])
			i++
		} else {
			final = append(final, second[j])
			j++
		}
	}
	for ; i < len(first); i++ {
		final = append(final, first[i])
	}
	for ; j < len(second); j++ {
		final = append(final, second[j])
	}
	return final
}

// main function
func main() {

	// create unsorted array using rand method
	var tableToSort []int
	for i := 0; i < 10000; i++ {
		tableToSort = append(tableToSort, rand.Intn(10000))
	}

	// print the data
	fmt.Println("Original array:", tableToSort)

	// sort the data using the mergesort method
	sorted := mergeSort(tableToSort)

	// print the sorted daa
	fmt.Println("Sorted array:", sorted)

}
