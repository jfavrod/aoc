package main

import (
	"fmt"
	"os"
)

const BUFFER_SIZE = 14

func main() {
	var start = 0
	workDir, _ := os.Getwd()
	file, err := os.Open(fmt.Sprintf("%s/day6/input.txt", workDir))

	if err == nil {
		getStartIndex(file, &start)
	}

	defer file.Close()

	fmt.Printf("Start at: %d\n", start)
}

func getStartIndex(file *os.File, start *int) {
	var char []byte = make([]byte, 1)
	buffer := [BUFFER_SIZE]byte{}

	*start = 0

	for ; char[0] != '\n'; file.Read(char) {
		bufferLen := len(buffer)
		next := char[0]

		if *start < (BUFFER_SIZE - 1) {
			buffer[bufferLen-1] = next
		} else {
			for i := 0; i < BUFFER_SIZE-1; i++ {
				buffer[i] = buffer[i+1]
			}

			buffer[BUFFER_SIZE-1] = next

			if !hasDuplicateBytes(&buffer) {
				break
			}
		}

		(*start)++
	}
}

func hasDuplicateBytes(buffer *[BUFFER_SIZE]byte) bool {
	for i := 0; i < BUFFER_SIZE; i++ {
		for j := 0; j < BUFFER_SIZE; j++ {
			if j != i {
				if (*buffer)[i] == (*buffer)[j] {
					return true
				}
			}
		}
	}

	return false
}
