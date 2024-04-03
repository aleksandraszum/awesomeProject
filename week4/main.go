package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// this program assesses a company's salary ranger for an employer
// it finds the smallest, average, and largest salary in a company
// data are from txt file

// function that calculates min salary
func minSalary(m map[string]int) string {
	min := 1000000000000
	var employerName string
	for name, salary := range m {
		if salary < min {
			min = salary
			employerName = name
		}
	}
	return employerName
}

// function that calculates max salary
func maxSalary(m map[string]int) string {
	var maxVal int = 0
	var employerName string
	for name, salary := range m {
		if salary > maxVal {
			maxVal = salary
			employerName = name
		}
	}
	return employerName
}

// function that calculates the average
func calculateAverage(m map[string]int) float64 {
	var sum int = 0
	var i int
	for _, salary := range m {
		sum = sum + salary
		i = i + 1
	}
	var avg float64
	avg = float64(sum) / float64(i)
	return avg
}

func main() {

	// initialise the map
	dataMap := make(map[string]int)

	// read the data from txt file and save as a map
	var fileName string
	fileName = "C:\\Users\\Ola\\awesomeProject\\week4\\names.txt"
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	for i := 0; i < len(fileLines); i++ {
		res := strings.Split(fileLines[i], " ")
		lastAndFirstName := res[0] + " " + res[1]
		intVar, err := strconv.Atoi(res[2])
		if err != nil {
			//	fmt.Println("Invalid salary:")
			continue
		}
		dataMap[lastAndFirstName] = intVar
	}

	//inicialise the variable
	var smallest, largest string
	var averageSalary float64

	// find the smallest salary
	smallest = minSalary(dataMap)

	// find the largest salary
	largest = maxSalary(dataMap)

	// calculate an average
	averageSalary = calculateAverage(dataMap)

	// print results
	fmt.Println("The smallest salary has: ", smallest, "the average of salary is: ",
		averageSalary, ", and the largest salary has: ", largest)

}
