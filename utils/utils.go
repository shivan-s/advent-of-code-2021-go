package utils

import (
	"bufio"
	"log"
	"os"
)

// Read data from .txt file line by line
func ReadDataFromFile(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var values []string
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}
	return values, scanner.Err()
}
