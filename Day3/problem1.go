package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isNumber(inputCharacter byte) bool {
	return inputCharacter >= '0' && inputCharacter <= '9'
}

func getNumber(inputLine string, startPos int) int {
	numberSubStr := string(inputLine[startPos])

	// start searching left
	for i := startPos - 1; i >= 0 && isNumber(inputLine[i]); i-- {
		numberSubStr = fmt.Sprintf("%c%s", inputLine[i], numberSubStr)
	}

	for i := startPos + 1; i < len(inputLine) && isNumber(inputLine[i]); i++ {
		numberSubStr = fmt.Sprintf("%s%c", numberSubStr, inputLine[i])
	}

	// fmt.Printf("number as str: %s", numberSubStr)

	result, _ := strconv.Atoi(numberSubStr)
	return result
}

func main() {
	inputFile, err := os.Open("day3.input.official")
	defer inputFile.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	inputScanner := bufio.NewScanner(inputFile)
	inputScanner.Split(bufio.ScanLines)
	var engineMap []string
	sum := 0

	// fill in the engine map
	for inputScanner.Scan() {
		line := inputScanner.Text()
		engineMap = append(engineMap, line)
	}

	for i := 0; i < len(engineMap); i++ {
		for j := 0; j < len(engineMap[i]); j++ {
			// find the special characters
			if !(engineMap[i][j] == '.' || isNumber(engineMap[i][j])) {
				fmt.Printf("found %c at (%d,%d): \t", engineMap[i][j], i+1, j+1)

				// find the near numbers
				if i-1 >= 0 && isNumber(engineMap[i-1][j]) {
					num := getNumber(engineMap[i-1], j)
					fmt.Printf("U number(%d) ", num)
					sum += num
				} else {
					if i-1 >= 0 && isNumber(engineMap[i-1][j-1]) {
						num := getNumber(engineMap[i-1], j-1)
						fmt.Printf("UL number(%d) ", num)
						sum += num
					}
					if i-1 >= 0 && isNumber(engineMap[i-1][j+1]) {
						num := getNumber(engineMap[i-1], j+1)
						fmt.Printf("UR number(%d) ", num)
						sum += num
					}
				}
				if i+1 < len(engineMap) && isNumber(engineMap[i+1][j]) {
					num := getNumber(engineMap[i+1], j)
					fmt.Printf("D number(%d) ", num)
					sum += num
				} else {
					if i+1 >= 0 && isNumber(engineMap[i+1][j-1]) {
						num := getNumber(engineMap[i+1], j-1)
						fmt.Printf("DL number(%d) ", num)
						sum += num
					}
					if i+1 >= 0 && isNumber(engineMap[i+1][j+1]) {
						num := getNumber(engineMap[i+1], j+1)
						fmt.Printf("DR number(%d) ", num)
						sum += num
					}
				}
				if j-1 >= 0 && isNumber(engineMap[i][j-1]) {
					num := getNumber(engineMap[i], j-1)
					fmt.Printf("L number(%d) ", num)
					sum += num
				}
				if j+1 < len(engineMap) && isNumber(engineMap[i][j+1]) {
					num := getNumber(engineMap[i], j+1)
					fmt.Printf("R number(%d) ", num)
					sum += num
				}
				fmt.Print("\n")
			}
		}
	}

	fmt.Printf("final sum: %d\n", sum)
}
