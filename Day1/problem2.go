package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
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
	var firstDigit int
	var lastDigit int
	var minPos int
	var maxPos int
	// constants, don't change
	var numWords = [...]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var numDigits = [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	lineNumber := 1

	for inputScanner.Scan() {
		// fmt.Printf("%d:\t", lineNumber)
		line := inputScanner.Text()
		firstDigit = 0
		lastDigit = 0
		minPos = math.MaxInt
		maxPos = math.MinInt

		for digit := 0; digit <= 9; digit++ {
			var posWord int
			var posNum int
			// get the positions for num and word digits by first index
			posWord = strings.Index(line, numWords[digit])

			if posWord != -1 {
				if posWord < minPos {
					// fmt.Printf("%s found at: %d\t", numWords[digit], posWord+1)
					minPos = posWord
					firstDigit = digit
				}
				if posWord > maxPos {
					// fmt.Printf("%s found at: %d\t", numWords[digit], posWord+1)
					maxPos = posWord
					lastDigit = digit
				}
			}

			posNum = strings.Index(line, numDigits[digit])

			if posNum != -1 {
				if posNum < minPos {
					// fmt.Printf("%s found at: %d\t", numDigits[digit], posNum+1)
					minPos = posNum
					firstDigit = digit
				}
				if posNum > maxPos {
					// fmt.Printf("%s found at: %d\t", numDigits[digit], posNum+1)
					maxPos = posNum
					lastDigit = digit
				}
			}

			// get the positions for num and word digits by last index
			posWord = strings.LastIndex(line, numWords[digit])

			if posWord != -1 {
				if posWord < minPos {
					// fmt.Printf("%s found at: %d\t", numWords[digit], posWord+1)
					minPos = posWord
					firstDigit = digit
				}
				if posWord > maxPos {
					// fmt.Printf("%s found at: %d\t", numWords[digit], posWord+1)
					maxPos = posWord
					lastDigit = digit
				}
			}

			posNum = strings.LastIndex(line, numDigits[digit])

			if posNum != -1 {
				if posNum < minPos {
					// fmt.Printf("%s found at: %d\t", numDigits[digit], posNum+1)
					minPos = posNum
					firstDigit = digit
				}
				if posNum > maxPos {
					// fmt.Printf("%s found at: %d\t", numDigits[digit], posNum+1)
					maxPos = posNum
					lastDigit = digit
				}
			}
		}

		number := firstDigit*10 + lastDigit
		sum += number
		// fmt.Printf(" number: %d\n", number)
		// fmt.Println(number)
		lineNumber++
	}

	fmt.Printf("sum: %d\n", sum)
}
