package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	L string
	R string
}

func main() {
	inputFile, err := os.Open("day8.input.official")
	defer inputFile.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	inputScanner := bufio.NewScanner(inputFile)
	inputScanner.Split(bufio.ScanLines)

	camelMap := map[string]Node{}

	inputScanner.Scan()
	directions := inputScanner.Text()
	inputScanner.Scan()

	// populate the map
	for inputScanner.Scan() {
		line := inputScanner.Text()
		tokens := strings.Split(line, " = ")
		node := tokens[0]
		trimmedToken := strings.TrimLeft(strings.TrimRight(tokens[1], ")"), "(")
		children := strings.Split(trimmedToken, ", ")
		camelMap[node] = Node{L: children[0], R: children[1]}
	}
	fmt.Printf("camelMap: %v\ndirections: %v\n", camelMap, directions)

	currentNode := "AAA"
	steps := 0
	for currentNode != "ZZZ" {
		for _, d := range directions {
			if d == 'L' {
				currentNode = camelMap[currentNode].L
			} else {
				currentNode = camelMap[currentNode].R
			}
			steps++
			if currentNode == "ZZZ" {
				break
			}
		}
	}

	fmt.Printf("Steps: %d\n", steps)
}
