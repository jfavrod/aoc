package input

import (
	"bufio"
	"fmt"
	"os"
)

func GetInputForDay(day int) *bufio.Scanner {
	dir, _ := os.Getwd()
	file, err := os.Open(fmt.Sprintf("%s/day%d/input.txt", dir, day))

	if err != nil {
		os.Exit(1)
	}

	return bufio.NewScanner(file)
}
