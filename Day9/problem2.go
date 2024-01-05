package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IsConstant(vec []int) bool {
	if len(vec) <= 1 {
		return true
	}
	value := vec[0]

	for i := 1; i < len(vec); i++ {
		if vec[i] != value {
			return false
		}
	}

	return true
}

func Derivate(vec []int) []int {
	if len(vec) <= 1 {
		return vec
	}
	dVec := []int{}

	for i := 1; i < len(vec); i++ {
		dVec = append(dVec, vec[i]-vec[i-1])
	}

	return dVec
}

func Extrapolate(vec []int) int {
	if IsConstant(vec) {
		return vec[0] // return the constant
	}

	return vec[0] - Extrapolate(Derivate(vec))
}

func main() {
	inputFile, err := os.Open("day9.input.official")
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

		vec := []int{}

		for _, numStr := range strings.Split(line, " ") {
			if numInt, err := strconv.Atoi(numStr); err == nil {
				vec = append(vec, numInt)
			}
		}

		extrapolated := Extrapolate(vec)
		sum += extrapolated
		fmt.Printf("for vec: %v extrapolate value: %d\n", vec, extrapolated)
	}

	// vec := []int{0, 3, 6, 9, 12, 15}
	// fmt.Printf("vec: %v\ndVec: %v\n", vec, Derivate(vec))
	// fmt.Printf("dVec is const: %t\n", IsConstant(Derivate(vec)))
	// fmt.Printf("extrapolated value: %d\n", Extrapolate(vec))

	fmt.Printf("\nFinal sum: %d\n", sum)
}
