package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	// internalMap := make([][]int, len(stepsMap))
	// for i := range stepsMap {
	// 	internalMap[i] = make([]int, len(stepsMap[i]))
	// 	copy(internalMap[i], stepsMap[i])
	// }
	internalMap := stepsMap
	internalMap[line][col] = step

	// fmt.Printf("internal map: %v\n", internalMap)
	if CheckNeighbour(internalMap, restrictionsMap, line, col, line-1, col) {
		WalkContour(internalMap, restrictionsMap, line-1, col, step)
	}
	if CheckNeighbour(internalMap, restrictionsMap, line, col, line+1, col) {
		WalkContour(internalMap, restrictionsMap, line+1, col, step)
	}
	if CheckNeighbour(internalMap, restrictionsMap, line, col, line, col-1) {
		WalkContour(internalMap, restrictionsMap, line, col-1, step)
	}
	if CheckNeighbour(internalMap, restrictionsMap, line, col, line, col+1) {
		WalkContour(internalMap, restrictionsMap, line, col+1, step)
	}
}

func CheckCanFill(stepsMap [][]int, restrictionsMap []string, nextLine, nextCol int) bool {
	if nextLine < 0 || nextCol < 0 || nextLine >= len(stepsMap) || nextCol >= len(stepsMap[0]) {
		return false
	}

	// nextLocation := restrictionsMap[nextLine][nextCol]
	nextStep := stepsMap[nextLine][nextCol]

	return nextStep == 0

	// if nextStep != 0 {
	// 	return false
	// }

	// return stepsMap[nextLine][nextCol] == 0
}

func CheckIsInside(stepsMap [][]int, restrictionsMap []string, nextLine, nextCol int) bool {
	if nextLine < 0 || nextCol < 0 || nextLine >= len(stepsMap) || nextCol >= len(stepsMap[0]) {
		return false
	}

	countLeft := 0

	// count left pipes
	for i := 0; i < nextCol; i++ {
		if (restrictionsMap[nextLine][i] == '|' || restrictionsMap[nextLine][i] == '7' || restrictionsMap[nextLine][i] == 'J' || restrictionsMap[nextLine][i] == 'F' || restrictionsMap[nextLine][i] == 'L') && stepsMap[nextLine][i] == 1 {
			countLeft++
		}
	}

	countRight := 0

	// count right pipess
	for i := nextCol + 1; i < len(restrictionsMap[nextLine]); i++ {
		if (restrictionsMap[nextLine][i] == '|' || restrictionsMap[nextLine][i] == '7' || restrictionsMap[nextLine][i] == 'J' || restrictionsMap[nextLine][i] == 'F' || restrictionsMap[nextLine][i] == 'L') && stepsMap[nextLine][i] == 1 {
			countRight++
		}
	}

	countUp := 0

	// count up pipes
	for i := 0; i < nextLine; i++ {
		if (restrictionsMap[i][nextCol] == '-' || restrictionsMap[i][nextCol] == '7' || restrictionsMap[i][nextCol] == 'J' || restrictionsMap[i][nextCol] == 'F' || restrictionsMap[i][nextCol] == 'L') && stepsMap[i][nextCol] == 1 {
			countUp++
		}
	}

	countDown := 0

	// count down pipess
	for i := nextLine + 1; i < len(restrictionsMap); i++ {
		if (restrictionsMap[i][nextCol] == '-' || restrictionsMap[i][nextCol] == '7' || restrictionsMap[i][nextCol] == 'J' || restrictionsMap[i][nextCol] == 'F' || restrictionsMap[i][nextCol] == 'L') && stepsMap[i][nextCol] == 1 {
			countDown++
		}
	}

	// if nextLine == 2 && nextCol == 2 {
	// 	fmt.Printf("pipes: L%d,R%d,U%d,D%d\n", countLeft, countRight, countUp, countDown)
	// }

	// it is inside if at least one ray casted line intersects the loop an odd amount
	// return countLeft%2 == 1 || countRight%2 == 1 || countUp%2 == 1 || countDown%2 == 1
	return countLeft%2 == 1 || countRight%2 == 1 || countUp%2 == 1 || countDown%2 == 1

}

func FloodFill(stepsMap [][]int, restrictionsMap []string, line, col, value int) {
	stepsMap[line][col] = value

	if CheckCanFill(stepsMap, restrictionsMap, line-1, col) {
		FloodFill(stepsMap, restrictionsMap, line-1, col, value)
	}
	if CheckCanFill(stepsMap, restrictionsMap, line+1, col) {
		FloodFill(stepsMap, restrictionsMap, line+1, col, value)
	}
	if CheckCanFill(stepsMap, restrictionsMap, line, col-1) {
		FloodFill(stepsMap, restrictionsMap, line, col-1, value)
	}
	if CheckCanFill(stepsMap, restrictionsMap, line, col+1) {
		FloodFill(stepsMap, restrictionsMap, line, col+1, value)
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

	// filter out the fills that touch the edges
	for i := 0; i < height; i++ {
		if internalMap[i][0] != -1 && CheckCanFill(internalMap, restrictionMap, i, 0) {
			FloodFill(internalMap, restrictionMap, i, 0, -1)
		}
		if internalMap[i][width-1] != -1 && CheckCanFill(internalMap, restrictionMap, i, width-1) {
			FloodFill(internalMap, restrictionMap, i, width-1, -1)
		}
	}
	for j := 0; j < width; j++ {
		if internalMap[0][j] != -1 && CheckCanFill(internalMap, restrictionMap, 0, j) {
			FloodFill(internalMap, restrictionMap, 0, j, -1)
		}
		if internalMap[height-1][j] != -1 && CheckCanFill(internalMap, restrictionMap, height-1, j) {
			FloodFill(internalMap, restrictionMap, height-1, j, -1)
		}
	}

	region := 2
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			// fmt.Printf("%c(%d,%d):%t\n", restrictionMap[i][j], i, j, CheckIsInside(internalMap, restrictionMap, i, j))
			if internalMap[i][j] == 0 && CheckIsInside(internalMap, restrictionMap, i, j) {
				FloodFill(internalMap, restrictionMap, i, j, region)
				// internalMap[i][j] = region
				// region
			}
		}
	}

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

	stepsMap := GenerateStepsMap(metalMap, startLine, startCol)
	fmt.Printf("starting coords: (%d, %d)\n", startLine, startCol)

	fmt.Printf("map colored:\n")

	enclosed := 0
	for i := range stepsMap {
		for j := range stepsMap[i] {
			fmt.Printf("%d\t", stepsMap[i][j])
			if stepsMap[i][j] == 2 {
				enclosed++
			}
		}
		fmt.Println()
	}
	fmt.Printf("\nenclosed: %d\n", enclosed)
}
