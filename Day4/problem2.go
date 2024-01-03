package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("day4.input.official")
	defer inputFile.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	inputScanner := bufio.NewScanner(inputFile)
	inputScanner.Split(bufio.ScanLines)
	cardScores := []int{}

	for inputScanner.Scan() {
		line := inputScanner.Text()

		splitLine := strings.Split(line, ":")
		numbers := strings.Split(splitLine[1], "|")
		winningKeys := strings.Split(strings.Trim(numbers[0], " "), " ")
		winningNumbers := []int{}

		winningNumbersMap := make(map[int]bool)

		// create the winning numbers map
		for _, key := range winningKeys {
			if intKey, err := strconv.Atoi(strings.Trim(key, " ")); err == nil {

				winningNumbersMap[intKey] = true
			}
		}

		// loop through the scratch numbers
		for _, num := range strings.Split(strings.Trim(numbers[1], " "), " ") {
			if intNum, err := strconv.Atoi(num); err == nil {

				if winningNumbersMap[intNum] {
					winningNumbers = append(winningNumbers, intNum)
				}
			}
		}

		score := len(winningNumbers)
		cardScores = append(cardScores, score)
	}

	fmt.Printf("initial card scores array: %v\n", cardScores)

	// count original cards
	sum := len(cardScores)
	// cardScores[len(cardScores)-1] = 1

	// loop through in reverse to avoid unneccessary recursion
	for i := len(cardScores) - 2; i >= 0; i-- {

		fmt.Printf("Card %d with %d -> ", i+1, cardScores[i])

		duplicates := cardScores[i]
		// count the winnings duplicates
		for j := 1; j <= cardScores[i] && i+j < len(cardScores); j++ {
			duplicates += cardScores[i+j]
			// fmt.Printf(" + %d", cardScores[i+j])
		}
		cardScores[i] = duplicates
		fmt.Printf("%d\n", cardScores[i])

		sum += cardScores[i]
	}

	// fmt.Printf("final card scores array: %v\n", cardScores)
	fmt.Printf("total sum: %d\n", sum)
}
