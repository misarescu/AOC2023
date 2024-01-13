package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var validCount int

func IsFinish(targets []*rune, springs []rune, validationNumbers []int, pos int) bool {
	if pos != len(targets)-1 {
		return false
	}

	counts := []int{}
	// isCounting := false
	currentCount := 0

	for _, s := range springs {
		if s == '#' {
			currentCount++
		} else {
			if currentCount > 0 {
				counts = append(counts, currentCount)
				currentCount = 0
			}

		}
	}

	if currentCount > 0 {
		counts = append(counts, currentCount)
		currentCount = 0
	}

	// fmt.Printf("counts: %v\n", counts)

	if len(counts) != len(validationNumbers) {
		return false
	}

	for i := 0; i < len(counts); i++ {
		if counts[i] != validationNumbers[i] {
			return false
		}
	}

	fmt.Printf("combination: %s\tcounts: %v\n", string(springs), counts)
	return true
	// return sort.IntsAreSorted(indeces)
}

func IsValid(targets []*rune, validationNumbers []int, pos int) bool {
	return true
}

func ComputeCombinations(targets []*rune, springs []rune, validationNumbers []int, pos int, validCount *int) {
	// fmt.Printf("pos:%d\n", pos)
	if pos < len(targets) {
		for _, value := range ".#" {
			*targets[pos] = value
			if IsValid(targets, validationNumbers, pos) {
				if IsFinish(targets, springs, validationNumbers, pos) {
					// fmt.Printf("combination: %s\n", string(springs))
					*validCount++
				} else {
					ComputeCombinations(targets, springs, validationNumbers, pos+1, validCount)
				}
			}
		}
	}
}

func main() {
	inputFile, err := os.Open("day12.input.official")
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

		lineTokens := strings.Split(line, " ")
		numbers := []int{}
		springs := []rune{}
		for _, char := range lineTokens[0] {
			springs = append(springs, char)
		}
		strings.Split(lineTokens[0], "")

		for _, numStr := range strings.Split(lineTokens[1], ",") {
			if numInt, err := strconv.Atoi(numStr); err == nil {
				numbers = append(numbers, numInt)
			}
		}

		targets := []*rune{}

		for idx := range springs {
			if springs[idx] == '?' {
				targets = append(targets, &springs[idx])
			}
		}

		fmt.Printf("map: %s\tnumbers:%v\n", string(springs), numbers)
		validCount := 0
		ComputeCombinations(targets, springs, numbers, 0, &validCount)
		fmt.Printf("valid count: %d\n", validCount)
		sum += validCount
	}

	fmt.Printf("total sum: %d\n", sum)
}
