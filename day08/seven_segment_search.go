package day08

import (
	"bufio"
	"os"
	"strings"
)

// extract and parse .txt file
func readDataFromFile(file string) [][]string {
	// turn string into string array by splting by space
	fileData, _ := os.Open(file)
	scanner := bufio.NewScanner(fileData)
	var output [][]string
	for scanner.Scan() {
		strSlice := strings.Split(scanner.Text(), " ")
		output = append(output, strSlice)

	}
	return output
}

// read the digits
func DigitReader(file string) int {
	input := readDataFromFile(file)
	// iterate through the slice of strings
	// find the length of the string slice
	// look at last 4 elements
	// count elements if only equal to below numbers
	// 1 segs 2
	// 4 segs 4
	// 7 segs 3
	// 8 segs 7

	var count int
	var observeSlice []string
	for _, strSlice := range input {
		observeSlice = strSlice[len(strSlice)-4:]
		for _, str := range observeSlice {
			if (len(str) == 2) || (len(str) == 4) || (len(str) == 3) || (len(str) == 7) {
				count++
			}
		}
	}
	return count
}
