package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func randomChoice(array []string) string {
	return array[rand.Intn(len(array))]
}

func main() {
	rand.Seed(time.Now().UnixNano())
	activities, _ := readLines("lists/activity")
	adjectives, _ := readLines("lists/adjective")
	locations, _ := readLines("lists/location")
	persons, _ := readLines("lists/person")
	inputs := [][]string{adjectives, persons, locations, activities}
	connectors := []string{" ", " in ", " ", ""}
	for i, input := range inputs {
		fmt.Print(randomChoice(input))
		fmt.Print(connectors[i])
	}
}
