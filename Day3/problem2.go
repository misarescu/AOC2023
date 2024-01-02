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
			gearRatio := 1
			adjacentCount := 0
			adjacentList := ""

			// find the gear characters
			if engineMap[i][j] == '*' {

				// find the near numbers
				if i-1 >= 0 && isNumber(engineMap[i-1][j]) {
					num := getNumber(engineMap[i-1], j)
					adjacentCount++
					if adjacentCount <= 2 {
						adjacentList = fmt.Sprintf("%sU number(%d) ", adjacentList, num)
						gearRatio *= num
					}
				} else {
					if i-1 >= 0 && isNumber(engineMap[i-1][j-1]) {
						num := getNumber(engineMap[i-1], j-1)
						adjacentCount++
						if adjacentCount <= 2 {
							adjacentList = fmt.Sprintf("%sUL number(%d) ", adjacentList, num)
							gearRatio *= num
						}
					}
					if i-1 >= 0 && isNumber(engineMap[i-1][j+1]) {
						num := getNumber(engineMap[i-1], j+1)
						adjacentCount++
						if adjacentCount <= 2 {
							adjacentList = fmt.Sprintf("%sUR number(%d) ", adjacentList, num)
							gearRatio *= num
						}
					}
				}
				if i+1 < len(engineMap) && isNumber(engineMap[i+1][j]) {
					num := getNumber(engineMap[i+1], j)
					adjacentCount++
					if adjacentCount <= 2 {
						adjacentList = fmt.Sprintf("%sD number(%d) ", adjacentList, num)
						gearRatio *= num
					}
				} else {
					if i+1 >= 0 && isNumber(engineMap[i+1][j-1]) {
						num := getNumber(engineMap[i+1], j-1)
						adjacentCount++
						if adjacentCount <= 2 {
							adjacentList = fmt.Sprintf("%sDL number(%d) ", adjacentList, num)
							gearRatio *= num
						}
					}
					if i+1 >= 0 && isNumber(engineMap[i+1][j+1]) {
						num := getNumber(engineMap[i+1], j+1)
						adjacentCount++
						if adjacentCount <= 2 {
							adjacentList = fmt.Sprintf("%sDR number(%d) ", adjacentList, num)
							gearRatio *= num
						}
					}
				}
				if j-1 >= 0 && isNumber(engineMap[i][j-1]) {
					num := getNumber(engineMap[i], j-1)
					adjacentCount++
					if adjacentCount <= 2 {
						adjacentList = fmt.Sprintf("%sL number(%d) ", adjacentList, num)
						gearRatio *= num
					}
				}
				if j+1 < len(engineMap) && isNumber(engineMap[i][j+1]) {
					num := getNumber(engineMap[i], j+1)
					adjacentCount++
					if adjacentCount <= 2 {
						adjacentList = fmt.Sprintf("%sR number(%d) ", adjacentList, num)
						gearRatio *= num
					}
				}

				if adjacentCount == 2 {
					fmt.Printf("gear found at (%d,%d): %s\t", i+1, j+1, adjacentList)
					fmt.Print("\n")
					sum += gearRatio
				}
			}
		}
	}

	fmt.Printf("final sum: %d\n", sum)
}
