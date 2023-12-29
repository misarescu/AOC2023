package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open("day1.input.official")
	defer inputFile.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	inputScanner := bufio.NewScanner(inputFile)

	inputScanner.Split(bufio.ScanLines)

	var sum int
	var isFound bool
	var intElem int
	var firstElem int
	var lastElem int

	for inputScanner.Scan() {

		line := inputScanner.Text()
		isFound = false
		firstElem = 0
		lastElem = 0
		intElem = 0

		for _, c := range line {
			if intElem, err = strconv.Atoi(string(c)); err == nil {
				if !isFound {

					isFound = true
					firstElem = intElem

				}

				lastElem = intElem

			}
		}

		number := firstElem*10 + lastElem
		sum += number

	}

	fmt.Printf("sum: %d\n", sum)
}
