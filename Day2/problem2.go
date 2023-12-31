package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("day2.input.official")
	defer inputFile.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	inputScanner := bufio.NewScanner(inputFile)

	inputScanner.Split(bufio.ScanLines)

	const redMax = 12
	const greenMax = 13
	const blueMax = 14

	var idSum int
	for inputScanner.Scan() {
		line := inputScanner.Text()
		game := strings.Split(line, ": ") // separate Game: and the list of dice
		gameId, _ := strconv.Atoi(strings.Split(game[0], " ")[1])
		var redCount int
		var greenCount int
		var blueCount int

		var redMaxCount int
		var greenMaxCount int
		var blueMaxCount int

		fmt.Printf("Game %d: ", gameId)
		// loop through all the bags in the game
		for _, bag := range strings.Split(game[1], "; ") {

			// loop through all the cubes in a bag
			for _, cube := range strings.Split(bag, ", ") {
				props := strings.Split(cube, " ")
				count, _ := strconv.Atoi(props[0])
				color := props[1]

				switch color {
				case "red":
					if count > redCount {
						redCount = count
					}
					break
				case "green":
					if count > greenCount {
						greenCount = count
					}
					break
				case "blue":
					if count > blueCount {
						blueCount = count
					}
					break
				}
			} // end looop through cubes
		} // end loop through bags
		if redCount >= redMaxCount {
			redMaxCount = redCount
		}
		if greenCount >= greenMaxCount {
			greenMaxCount = greenCount
		}
		if blueCount >= blueMaxCount {
			blueMaxCount = blueCount
		}

		idSum += (redMaxCount * greenMaxCount * blueMaxCount)
		fmt.Printf("reds: %d, greens: %d, blues: %d, power: %d", redMaxCount, greenMaxCount, blueMaxCount, redMaxCount*greenMaxCount*blueMaxCount)
		fmt.Println("")
	}

	fmt.Printf("Sum ID: %d\n", idSum)
}
