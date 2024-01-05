package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var maxSteps int

type Location struct {
	line int
	col  int
}

func CheckNeighbour(stepsMap [][]int, restrictionsMap []string, currLine, currCol, nextLine, nextCol int) bool {
	if nextLine < 0 || nextCol < 0 || nextLine >= len(stepsMap) || nextCol >= len(stepsMap[0]) {
		return false
	}

	currLocation := restrictionsMap[currLine][currCol]
	nextLocation := restrictionsMap[nextLine][nextCol]
	nextStep := stepsMap[nextLine][nextCol]

	if nextStep != 0 {
		return false
	}

	// fmt.Printf("current location %c, next location %c\n", currLocation, nextLocation)

	if currLocation == 'S' {
		return (nextLocation == '|' || nextLocation == '-' ||
			nextLocation == 'L' || nextLocation == 'J' ||
			nextLocation == '7' || nextLocation == 'F')
	}

	if nextLocation == 'S' {
		fmt.Println("evrika")
		// return true
	}

	//move up
	if nextLine < currLine {
		return (currLocation == '|' || currLocation == 'L' || currLocation == 'J') && (nextLocation == '|' || nextLocation == '7' || nextLocation == 'F')
	}

	//move down
	if nextLine > currLine {
		return (currLocation == '|' || currLocation == '7' || currLocation == 'F') && (nextLocation == '|' || nextLocation == 'L' || nextLocation == 'J')
	}

	//move left
	if nextCol < currCol {
		return (currLocation == '-' || currLocation == '7' || currLocation == 'J') && (nextLocation == '-' || nextLocation == 'L' || nextLocation == 'F')
	}

	//move right
	if nextCol > currCol {
		return (currLocation == '-' || currLocation == 'L' || currLocation == 'F') && (nextLocation == '-' || nextLocation == '7' || nextLocation == 'J')
	}

	return false
}

func WalkContour(stepsMap [][]int, restrictionsMap []string, line, col, step int) {
	if step > maxSteps {
		maxSteps = step
	}
	internalMap := make([][]int, len(stepsMap))
	for i := range stepsMap {
		internalMap[i] = make([]int, len(stepsMap[i]))
		copy(internalMap[i], stepsMap[i])
	}
	internalMap[line][col] = step

	// fmt.Printf("internal map: %v\n", internalMap)
	if CheckNeighbour(internalMap, restrictionsMap, line, col, line-1, col) {
		WalkContour(internalMap, restrictionsMap, line-1, col, step+1)
	}
	if CheckNeighbour(internalMap, restrictionsMap, line, col, line+1, col) {
		WalkContour(internalMap, restrictionsMap, line+1, col, step+1)
	}
	if CheckNeighbour(internalMap, restrictionsMap, line, col, line, col-1) {
		WalkContour(internalMap, restrictionsMap, line, col-1, step+1)
	}
	if CheckNeighbour(internalMap, restrictionsMap, line, col, line, col+1) {
		WalkContour(internalMap, restrictionsMap, line, col+1, step+1)
	}
}

func GenerateStepsMap(restrictionMap []string, startLine, startCol int) [][]int {
	height := len(restrictionMap)
	width := len(restrictionMap[0])
	internalMap := [][]int{}

	for i := 0; i < height; i++ {
		internalMap = append(internalMap, make([]int, width))
	}

	WalkContour(internalMap, restrictionMap, startLine, startCol, 1)

	// fmt.Printf("walked path: %v\n", path)

	return internalMap
}

func main() {
	inputFile, err := os.Open("day10.input.official")
	defer inputFile.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	inputScanner := bufio.NewScanner(inputFile)
	inputScanner.Split(bufio.ScanLines)
	var startLine, startCol int

	metalMap := []string{}

	for inputScanner.Scan() {
		line := inputScanner.Text()

		if strings.Contains(line, "S") {
			startLine = len(metalMap)
			startCol = strings.Index(line, "S")
		}

		metalMap = append(metalMap, line)

		// fmt.Println(line)
	}

	maxSteps = 0
	GenerateStepsMap(metalMap, startLine, startCol)
	fmt.Printf("starting coords: (%d, %d)\n", startLine, startCol)

	fmt.Printf("\nwalked path: %v\n", maxSteps/2)
	// fmt.Printf("\nMaxSteps: %d\n", len(path)/2)
}
