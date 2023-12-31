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

		fmt.Printf("Game %d: ", gameId)
		// loop through all the bags in the game
		for _, bag := range strings.Split(game[1], "; ") {
			// fmt.Printf("game %d: %s\n", gameId, bag)

			// loop through all the cubes in a bag
			for _, cube := range strings.Split(bag, ", ") {
				props := strings.Split(cube, " ")
				count, _ := strconv.Atoi(props[0])
				color := props[1]

				// fmt.Printf("%d %s | ", count, color)
				// fmt.Printf("%v", props)

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
			// fmt.Print("|| ")
		} // end loop through bags
		fmt.Printf("reds: %d, greens: %d, blues: %d", redCount, greenCount, blueCount)
		fmt.Println("")

		// fmt.Printf("Game %d: reds: %d, greens: %d, blues: %d\n", gameId, redCount, greenCount, blueCount)
		if redCount <= redMax && greenCount <= greenMax && blueCount <= blueMax {
			// fmt.Printf("game %d is plausible with %d red, %d green, %d blue\n", gameId, redCount, greenCount, blueCount)
			idSum += gameId
		}
	}

	fmt.Printf("Sum ID: %d\n", idSum)
}
