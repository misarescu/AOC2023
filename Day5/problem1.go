package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Element struct {
	value   int
	visited bool
}

type GardenMap struct {
	mapList []struct {
		destinationStart  int
		destinationFinish int
		sourceStart       int
		sourceFinish      int
	}
}

func (targetMap *GardenMap) populateMap(inputScanner *bufio.Scanner, line *string) {
	for inputScanner.Scan() {
		*line = inputScanner.Text()
		// fmt.Printf("line: %s", line)
		numbers := []int{}
		for _, numberStr := range strings.Split(*line, " ") {
			if number, err := strconv.Atoi(numberStr); err == nil {
				numbers = append(numbers, number)
			}
		}

		if len(numbers) == 0 {
			break
		}

		// fmt.Printf("numbers: %v\n", numbers)

		destination := numbers[0]
		source := numbers[1]
		length := numbers[2]

		targetMap.mapList = append(targetMap.mapList, struct {
			destinationStart  int
			destinationFinish int
			sourceStart       int
			sourceFinish      int
		}{destinationStart: destination, destinationFinish: destination + length - 1, sourceStart: source, sourceFinish: source + length - 1})
	}
	fmt.Printf("map: %v\n", targetMap)
}

func (targetMap *GardenMap) getMapValue(input int) int {
	result := input

	for _, el := range targetMap.mapList {
		if input >= el.sourceStart && input <= el.sourceFinish {
			return el.destinationStart + (input - el.sourceStart)
		}
	}

	return result
}

func main() {
	inputFile, err := os.Open("day5.input.official")
	defer inputFile.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	inputScanner := bufio.NewScanner(inputFile)
	inputScanner.Split(bufio.ScanLines)

	seedsList := []int{}
	seedToSoilMap := GardenMap{}
	soilToFertilizerMap := GardenMap{}
	fertilizerToWaterMap := GardenMap{}
	waterToLightMap := GardenMap{}
	lightToTemperatureMap := GardenMap{}
	temperatureToHumidityMap := GardenMap{}
	humdityToLocaionMap := GardenMap{}

	for inputScanner.Scan() {
		line := inputScanner.Text()

		if strings.Contains(line, "seeds:") {
			for _, seedStr := range strings.Split(strings.Split(line, ": ")[1], " ") {
				if seed, err := strconv.Atoi(seedStr); err == nil {
					seedsList = append(seedsList, seed)
				}
			}
			// fmt.Printf("Seeds: %v\n", seedsList)
		}

		if strings.Contains(line, "seed-to-soil") {
			seedToSoilMap.populateMap(inputScanner, &line)
			// fmt.Printf("seed-to-soil map: %v\n", seedToSoilMap)
		}

		if strings.Contains(line, "soil-to-fertilizer") {
			soilToFertilizerMap.populateMap(inputScanner, &line)
			// fmt.Printf("soil-to-fertilizer map: %v\n", soilToFertilizerMap)
		}

		if strings.Contains(line, "fertilizer-to-water") {
			fertilizerToWaterMap.populateMap(inputScanner, &line)
			// fmt.Printf("fertilizer-to-water map: %v\n", fertilizerToWaterMap)
		}

		if strings.Contains(line, "water-to-light") {
			waterToLightMap.populateMap(inputScanner, &line)
			// fmt.Printf("fertilizer-to-water map: %v\n", waterToLightMap)
		}

		if strings.Contains(line, "light-to-temperature") {
			lightToTemperatureMap.populateMap(inputScanner, &line)
			// fmt.Printf("light-to-temperature map: %v\n", lightToTemperatureMap)
		}

		if strings.Contains(line, "temperature-to-humidity") {
			temperatureToHumidityMap.populateMap(inputScanner, &line)
			// fmt.Printf("temperature-to-humidity map: %v\n", temperatureToHumidityMap)
		}

		if strings.Contains(line, "humidity-to-location") {
			humdityToLocaionMap.populateMap(inputScanner, &line)
			// fmt.Printf("humidity-to-location map: %v\n", humdityToLocaionMap)
		}
	}

	fmt.Print("locations: ")
	minLocation := math.MaxInt
	for _, seed := range seedsList {
		// var soil, fertilizer, water, light, temperature, humidity, location int

		soil := seedToSoilMap.getMapValue(seed)
		fertilizer := soilToFertilizerMap.getMapValue(soil)
		water := fertilizerToWaterMap.getMapValue(fertilizer)
		light := waterToLightMap.getMapValue(water)
		temperature := lightToTemperatureMap.getMapValue(light)
		humidity := temperatureToHumidityMap.getMapValue(temperature)
		location := humdityToLocaionMap.getMapValue(humidity)

		if location < minLocation {
			minLocation = location
		}

		fmt.Printf("seed: %d, soil: %d, fertilizer: %d, water: %d, light: %d, temperature: %d, humidity: %d, location: %d,\n", seed, soil, fertilizer, water, light, temperature, humidity, location)
	}
	fmt.Printf("\nMin locaion: %d\n", minLocation)
}
