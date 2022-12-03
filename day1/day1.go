package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type elf struct {
	id       uint
	calories uint
}

func main() {
	var inputScanner bufio.Scanner
	var elfCalories map[uint]uint = map[uint]uint{}

	path, _ := os.Getwd()
	file, err := os.Open(path + "/day1/input.txt")

	if err == nil {
		inputScanner = *bufio.NewScanner(file)
	} else {
		fmt.Printf("%v\n", err.Error())
		os.Exit(1)
	}

	aggregateCalories(inputScanner, elfCalories)
	fmt.Printf("%v\n", sumTopThree(elfCalories))
}

// Use the given bufio.Scanner to read the input file and aggregate
// the calories carried by each elf into the give map.
func aggregateCalories(inputScanner bufio.Scanner, elfCalories map[uint]uint) {
	var elfCount uint = 0

	for inputScanner.Scan() {
		text := inputScanner.Text()

		if len(text) == 0 {
			elfCount++
		} else {
			converted, _ := strconv.Atoi(text)
			elfCalories[elfCount] = elfCalories[elfCount] + uint(converted)
		}
	}
}

// Iterate the given map and find the elf with the most calories.
func findMax(elfCalories map[uint]uint) elf {
	max := elf{id: 0, calories: elfCalories[0]}

	for id, cal := range elfCalories {
		if cal > max.calories {
			max = elf{id: id, calories: cal}
		}
	}

	return max
}

// Find the three elves with the most calories in the elfCalories map
// and add their calories together.
func sumTopThree(elfCalories map[uint]uint) uint {
	var sum uint = 0
	elfCount := 0

	for elfCount < 3 {
		elfCal := findMax(elfCalories)
		sum += elfCal.calories
		delete(elfCalories, elfCal.id)
		elfCount++
	}

	return sum
}
