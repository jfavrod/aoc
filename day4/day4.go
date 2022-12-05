package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type assignment struct {
	min uint
	max uint
}

type team struct {
	one assignment
	two assignment
}

func main() {
	var pairs = []team{}
	processPairs(&pairs)
}

func processPairs(pairs *[]team) {
	scanner := getInput()
	count := 0

	for scanner.Scan() {
		pair := team{}

		input := scanner.Text()
		breakOut := strings.Split(input, ",")

		for x := range [2]int{0, 1} {
			rawAssignment := strings.Split(breakOut[x], "-")
			min, minErr := strconv.Atoi(rawAssignment[0])
			max, maxErr := strconv.Atoi(rawAssignment[1])

			if minErr == nil && maxErr == nil {
				var ass *assignment

				if x == 0 {
					ass = &pair.one
				} else {
					ass = &pair.two
				}

				*ass = assignment{
					min: uint(min),
					max: uint(max),
				}
			} else {
				os.Exit(2)
			}
		}

		if isOverlap(&pair) {
			count++
		}

		*pairs = append(*pairs, pair)
	}

	fmt.Printf("count %d\n", count)
}

func isOverlap(pair *team) bool {
	return pair.one.min <= pair.two.min && pair.one.max >= pair.two.max ||
		pair.two.min <= pair.one.min && pair.two.max >= pair.one.max
}

func getInput() *bufio.Scanner {
	dir, _ := os.Getwd()
	file, err := os.Open(dir + "/day4/input.txt")

	if err != nil {
		os.Exit(1)
	}

	return bufio.NewScanner(file)
}
