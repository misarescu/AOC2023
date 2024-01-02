package main

import (
	"bufio"
	"fmt"
	"math"
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

	sum := 0

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

				// fmt.Println(intNum)
				if winningNumbersMap[intNum] {
					winningNumbers = append(winningNumbers, intNum)
				}
			}
		}

		if score := 0; len(winningNumbers) > 0 {
			exp := float64(len(winningNumbers) - 1)
			score = int(math.Pow(2.0, exp))
			sum += score
			fmt.Printf("%s:\twinning numbers: %v\t| score: %d \t| matched numbers: %v\n", splitLine[0], winningNumbersMap, score, winningNumbers)
		}
	}

	fmt.Printf("total sum: %d\n", sum)
}
