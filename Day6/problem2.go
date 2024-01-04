package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("day6.input.official")
	defer inputFile.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	inputScanner := bufio.NewScanner(inputFile)
	inputScanner.Split(bufio.ScanLines)

	var times []int
	var distances []int

	inputScanner.Scan()
	timeStrLine := inputScanner.Text()

	timeStr := strings.ReplaceAll(strings.Split(timeStrLine, ":")[1], " ", "")
	if timeInt, err := strconv.Atoi(timeStr); err == nil {
		times = append(times, timeInt)
	}

	inputScanner.Scan()
	distanceStr := inputScanner.Text()

	distStr := strings.ReplaceAll(strings.Split(distanceStr, ":")[1], " ", "")
	if distInt, err := strconv.Atoi(distStr); err == nil {
		distances = append(distances, distInt)
	}

	totalMargin := 1

	for race := 0; race < len(times); race++ {
		var hold int
		for hold = 1; hold <= times[race]/2; hold++ {
			// found the first hold time to win the race
			if hold*(times[race]-hold) > distances[race] {
				break
			}
		}
		margin := times[race] + 1 - 2*hold
		fmt.Printf("race: %d, first hold: %d, margin: %d\n", race, hold, margin)
		totalMargin *= margin
	}

	fmt.Printf("times: %v\n", times)
	fmt.Printf("distances: %v\n", distances)
	fmt.Printf("total margin: %d\n", totalMargin)

}
