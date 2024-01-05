package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	L string
	R string
}

func GetNodes(camelMap *map[string]Node, lastLetter byte) []string {
	result := []string{}

	for k := range *camelMap {
		if k[len(k)-1] == lastLetter {
			result = append(result, k)
		}
	}

	return result
}

func MoveNodes(camelMap *map[string]Node, nodes []string, move rune) []string {
	newNodes := []string{}
	for _, node := range nodes {
		if move == 'L' {
			newNodes = append(newNodes, (*camelMap)[node].L)
		} else {
			newNodes = append(newNodes, (*camelMap)[node].R)
		}
	}
	return newNodes
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func lcmArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = lcm(result, arr[i])
	}
	return result
}

func CheckIsDone(nodes []string) bool {
	for _, node := range nodes {
		if node[len(node)-1] != 'Z' {
			return false
		}
	}
	return true
}

func main() {
	inputFile, err := os.Open("day8.input.official")
	defer inputFile.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	inputScanner := bufio.NewScanner(inputFile)
	inputScanner.Split(bufio.ScanLines)

	camelMap := map[string]Node{}

	inputScanner.Scan()
	directions := inputScanner.Text()
	inputScanner.Scan()

	// populate the map
	for inputScanner.Scan() {
		line := inputScanner.Text()
		tokens := strings.Split(line, " = ")
		node := tokens[0]
		trimmedToken := strings.TrimLeft(strings.TrimRight(tokens[1], ")"), "(")
		children := strings.Split(trimmedToken, ", ")
		camelMap[node] = Node{L: children[0], R: children[1]}
	}
	fmt.Printf("camelMap:\n%v\n\ndirections: %v\n", camelMap, directions)

	currentNodes := GetNodes(&camelMap, 'A')
	finishNodes := GetNodes(&camelMap, 'Z')

	fmt.Printf("starting nodes: %v\n", currentNodes)
	fmt.Printf("finish nodes: %v\n", finishNodes)

	cycles := []int{}

	for i := range currentNodes {
		currentNode := currentNodes[i : i+1]
		steps := 0
		isFound := false
		for !isFound {
			for _, d := range directions {
				currentNode = MoveNodes(&camelMap, currentNode, d)
				// fmt.Printf("current nodes: %v after: %c\n", currentNode, d)
				steps++
				if isFound = CheckIsDone(currentNode); isFound {
					break
				}
			}
		}
		cycles = append(cycles, steps)
	}

	fmt.Printf("cycles: %d\n", cycles)
	fmt.Printf("greatest cycle: %d\n", lcmArray(cycles))
}
