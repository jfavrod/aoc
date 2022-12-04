package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	var inputScanner bufio.Scanner
	var sum uint = 0

	path, _ := os.Getwd()
	file, _ := os.Open(path + "/day3/input.txt")

	inputScanner = *bufio.NewScanner(file)

	for inputScanner.Scan() {
		line := inputScanner.Text()
		lineLen := len(line)
		mid := int(math.Floor(float64(lineLen) / 2.0))

		compartment1 := line[:mid]
		compartment2 := line[mid:]

		items := intersect(compartment1, compartment2)

		for idx := range items {
			sum += getPriority(items[idx])
		}
	}

	fmt.Printf("sum: %d\n", sum)
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
