package day10

import (
	"bufio"
	"os"
	"strings"
	"sync"
)

func readDataFromFile(file string) [][]string {
	fileData, _ := os.Open(file)
	scanner := bufio.NewScanner(fileData)
	var output [][]string
	for scanner.Scan() {
		strSlice := strings.SplitAfter(scanner.Text(), "")
		output = append(output, strSlice)
	}
	return output
}

type customStack struct {
	stack []string
	lock  sync.RWMutex
}

func (c *customStack) Push(name string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.stack = append(c.stack, name)
}

func (c *customStack) Pop() string {
	len := len(c.stack)
	var finalItem string
	if len > 0 {
		finalItem = c.stack[len-1]
		c.lock.Lock()
		defer c.lock.Unlock()
		c.stack = c.stack[:len-1]
	}
	return finalItem
}

// iterate through stringSlice
// keep tract of open (, open [, open {, open <
func FindCorrupt(file string) int {
	input := readDataFromFile(file)
	customStack := &customStack{
		stack: make([]string, 0),
	}
	bracketMap := map[string]string{
		"(": ")",
		"{": "}",
		"<": ">",
		"[": "]",
	}
	var score int
	var totalScore int

	for _, row := range input {
	outside:
		for _, c := range row { // < ( > )
			// add to stack if open
			if (c == "(") || (c == "[") || (c == "<") || (c == "{") {
				customStack.Push(c)
			} else {
				finalItem := customStack.Pop()
				if bracketMap[finalItem] != c {
					score = 0
					switch c {
					case ")":
						score = 3
					case "]":
						score = 57
					case "}":
						score = 1197
					case ">":
						score = 25137
					}
					totalScore += score
					break outside
				}
			}
		}
	}
	return totalScore
}

func RemoveCorruptAndScore(file string) int {
}
