package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	line int
	col  int
}

func ComputeDistance(emptyLines, emptyCols []int, point1, point2 Position) int {
	leftLine := 0
	rightLine := 0
	topCol := 0
	bottomCol := 0
	if point1.line > point2.line {
		leftLine = point2.line
		rightLine = point1.line
	} else {
		leftLine = point1.line
		rightLine = point2.line
	}
	if point1.col > point2.col {
		topCol = point1.col
		bottomCol = point2.col
	} else {
		topCol = point2.col
		bottomCol = point1.col
	}

	extraSpace := 0
	for _, line := range emptyLines {
		if line >= leftLine && line <= rightLine {
			extraSpace++
		}
	}
	for _, col := range emptyCols {
		if col >= bottomCol && col <= topCol {
			extraSpace++
		}
	}

	distance := rightLine - leftLine + topCol - bottomCol

	return extraSpace + distance
}

func main() {
	inputFile, err := os.Open("day11.input.official")
	defer inputFile.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	inputScanner := bufio.NewScanner(inputFile)
	inputScanner.Split(bufio.ScanLines)

	galaxyMap := []string{}
	emptyLines := []int{}
	emptyCols := []int{}
	galaxies := []Position{}

	for inputScanner.Scan() {
		line := inputScanner.Text()
		galaxyMap = append(galaxyMap, line)

		// if no galaxy in line, duplicate
		if !strings.Contains(line, "#") {
			emptyLines = append(emptyLines, len(galaxyMap)-1)
		}

		// fmt.Println(line)
	}

	mapSize := len(galaxyMap)

	for i := 0; i < mapSize; i++ {
		isEmpty := true
		for j := 0; j < mapSize; j++ {
			if galaxyMap[j][i] == '#' {
				isEmpty = false
				galaxies = append(galaxies, Position{line: j, col: i})
			}
		}
		if isEmpty {
			emptyCols = append(emptyCols, i)
		}
	}

	for i := 0; i < len(galaxyMap); i++ {
		fmt.Println(galaxyMap[i])
	}

	fmt.Printf("empty lines: %v\nempty cols: %v\ngalaxies:%v\n", emptyLines, emptyCols, galaxies)

	totalDistance := 0
	for p1 := 0; p1 < len(galaxies)-1; p1++ {
		for p2 := p1 + 1; p2 < len(galaxies); p2++ {
			currentDistance := ComputeDistance(emptyLines, emptyCols, galaxies[p1], galaxies[p2])
			fmt.Printf("distance from (%d,%d) to (%d,%d):%d\n", galaxies[p1].line, galaxies[p1].col, galaxies[p2].line, galaxies[p2].col, currentDistance)
			totalDistance += currentDistance
		}
	}

	fmt.Printf("\ntotal distance: %d\n", totalDistance)
}
