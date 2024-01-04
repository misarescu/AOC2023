package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// set the cards enum based on the values from 0 to 6
const (
	HIGH_CARD       = iota // 0
	ONE_PAIR               // 1
	TWO_PAIR               // 2
	THREE_OF_A_KIND        // 3
	FULL_HOUSE             // 4
	FOUR_OF_A_KIND         // 5
	FIVE_OF_A_KIND         // 6
)

type Hand struct {
	cards    string
	bid      int
	handType int
}

type ByType []Hand

func (a ByType) Len() int {
	return len(a)
}

func (a ByType) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByType) Less(i, j int) bool {
	scoreMap := map[rune]int{
		'A': 12,
		'K': 11,
		'Q': 10,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
		'J': 0,
	}

	if a[i].handType < a[j].handType {
		return true
	} else if a[i].handType == a[j].handType {
		for ii := 0; ii < 5; ii++ {
			if scoreMap[rune(a[i].cards[ii])] != scoreMap[rune(a[j].cards[ii])] {
				return scoreMap[rune(a[i].cards[ii])] < scoreMap[rune(a[j].cards[ii])]
			}
		}
	}
	return false
}

func DetermineType(cards string) int {
	frequencyMap := map[rune]int{}

	// set the frequencies
	for _, card := range cards {
		frequencyMap[card]++
	}

	var frequencies []int

	jPower := frequencyMap['J']
	delete(frequencyMap, 'J')

	// get the frequencies
	for _, v := range frequencyMap {
		frequencies = append(frequencies, v)
		// fmt.Printf("card:%c appears: %d\n", k, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(frequencies)))
	if len(frequencies) > 0 {
		frequencies[0] += jPower
	} else {
		frequencies = append(frequencies, jPower)
	}

	// check types
	if len(frequencies) == 1 { // check five of a kind
		return FIVE_OF_A_KIND
	} else if len(frequencies) == 2 { // check four of a kind or full house
		if frequencies[0] == 4 {
			return FOUR_OF_A_KIND
		} else {
			return FULL_HOUSE
		}
	} else if len(frequencies) == 3 {
		if frequencies[0] == 3 {
			return THREE_OF_A_KIND
		} else if frequencies[0] == 2 && frequencies[1] == 2 {
			return TWO_PAIR
		}
	} else if len(frequencies) == 4 {
		return ONE_PAIR
	} else {
		return HIGH_CARD
	}
	return 0
}

func NewHand(cardLine string) Hand {
	tokens := strings.Split(cardLine, " ")
	cards := tokens[0]
	bid, _ := strconv.Atoi(tokens[1])
	handType := DetermineType(cards)
	return Hand{cards: cards, bid: bid, handType: handType}
}

func main() {
	inputFile, err := os.Open("day7.input.official")
	defer inputFile.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	inputScanner := bufio.NewScanner(inputFile)
	inputScanner.Split(bufio.ScanLines)

	var hands []Hand

	for inputScanner.Scan() {
		line := inputScanner.Text()

		hand := NewHand(line)
		hands = append(hands, hand)
	}

	sort.Sort(ByType(hands))

	totalWinnings := 0

	for rank, hand := range hands {
		totalWinnings += ((rank + 1) * hand.bid)
	}

	fmt.Printf("hands: %v, total winnings: %d\n", hands, totalWinnings)

}
