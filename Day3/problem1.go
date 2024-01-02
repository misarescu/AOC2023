package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open("day3.input.demo1")
	defer inputFile.Close()

	if err != nil {
		fmt.Println(err.Error())
		return 
	}

	inputScanner := bufio.NewScanner(inputFile)
	inputScanner.Split(bufio.ScanLines)
	var engineMap []string

	for inputScanner.Scan(){
		line := inputScanner.Text()
		engineMap = append(engineMap, line)
	}

	for i := 0; i < len(engineMap); i++ {
		for j := 0; j < len(engineMap[i]); j++ {
			fmt.Printf("%c ", engineMap[i][j])
		}
		fmt.Printf("\n")
	}

	dotStr := "65$"
	dot, _ := strconv.Atoi(dotStr)
	fmt.Printf("\n\nThis is %s as int: %d", dotStr, dot)
}