package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var inputScanner bufio.Scanner
	var sum uint = 0

	path, _ := os.Getwd()
	file, _ := os.Open(path + "/day3/input.txt")

	inputScanner = *bufio.NewScanner(file)

	for {
		var lines = getLines(&inputScanner, 3)

		if len(lines) == 3 {
			items := intersect(lines[0], lines[1])
			items = intersect(string(items), lines[2])
			fmt.Printf("items: %v\n", string(items))

			for idx := range items {
				sum += getPriority(items[idx])
			}
		}

		// The last thing read was no bytes.
		// So we know we're done.
		if len(inputScanner.Bytes()) == 0 {
			break
		}
	}

	fmt.Printf("sum: %d\n", sum)
}

func getLines(inputScanner *bufio.Scanner, count uint) []string {
	var lines = []string{}

	for i := 0; i < int(count) && inputScanner.Scan(); i++ {
		lines = append(lines, inputScanner.Text())
	}

	return lines
}

func getPriority(char byte) uint {
	var pri uint = 0

	if char > 90 {
		pri = uint(char) - 96
	} else {
		pri = uint(char) - 38
	}

	return pri
}

func intersect(str1 string, str2 string) []byte {
	res := ""

	for idx := range str1 {
		subs := string(str1[idx])

		if strings.Contains(str2, subs) && !strings.Contains(res, subs) {
			res += subs
		}
	}

	return []byte(res)
}
