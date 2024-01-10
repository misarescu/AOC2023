package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func IsFinish(springs []rune, validationNumbers []int, pos int) bool {
	indeces := []int{}
	findPos := 0
	for _, num := range validationNumbers {
		validationStr := ""
		for i := 0; i < num; i++ {
			validationStr += "#"
		}
		idx := strings.Index(string(springs)[findPos:], validationStr)
		if idx < 0 {
			return false
		}
		findPos = idx + len(validationStr)
		indeces = append(indeces, idx)
	}

	return sort.IntsAreSorted(indeces)
}

func IsValid(targets []*rune, validationNumbers []int, pos int) bool {
	return true
}

func ComputeCombinations(targets []*rune, springs []rune, validationNumbers []int) int {
	pos := 0
	validCount := 0

}

func main() {
	inputFile, err := os.Open("day12.input.demo1")
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
		arrCount := 1

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

		sum += arrCount
		fmt.Printf("map: %s\tnumbers:%v\n", string(springs), numbers)

	}

	fmt.Printf("total sum: %d\n", sum)
}
