package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const rock = "rock"
const paper = "paper"
const scissors = "scissors"

type choice struct {
	name string
	val  uint
}

func main() {
	var inputScanner bufio.Scanner

	path, _ := os.Getwd()
	file, err := os.Open(path + "/two/input.txt")

	if err == nil {
		inputScanner = *bufio.NewScanner(file)
	} else {
		fmt.Printf("%v\n", err.Error())
	}

	fmt.Printf("Total: %d\n", processStrategyGuide(inputScanner))
}

func processStrategyGuide(inputScanner bufio.Scanner) uint {
	var gameTotal uint = 0

	for inputScanner.Scan() {
		text := inputScanner.Text()

		// split the text into an array with two strings.
		splitText := strings.Split(text, " ")

		// theirCode is the first (and only) byte of the first string
		// in the split text.
		theirCode := byte(splitText[0][0])

		// myCode is the first (and only) byte of the second string
		// in the split text.
		myCode := byte(splitText[1][0])

		theirChoice := getChoice(theirCode)
		myChoice := getChoice(theirCode, myCode)

		if didWin(&theirChoice, &myChoice) {
			gameTotal += 6
		} else if wasDraw(&theirChoice, &myChoice) {
			gameTotal += 3
		}

		gameTotal += myChoice.val
	}

	return gameTotal
}

func didWin(them *choice, us *choice) bool {
	return them.name == rock && us.name == paper ||
		them.name == paper && us.name == scissors ||
		them.name == scissors && us.name == rock
}

func wasDraw(them *choice, us *choice) bool {
	return them.name == us.name
}

func getChoice(codes ...byte) choice {
	codesLen := len(codes)

	if codesLen == 1 {
		switch codes[0] {
		case 'A':
			return newRock()
		case 'B':
			return newPaper()
		default:
			return newScissors()
		}
	}

	// Otherwise, the codes len should be 2.

	// They picked rock.
	if codes[0] == 'A' {
		switch codes[1] {
		case 'X':
			return newScissors()
		case 'Y':
			return newRock()
		default:
			return newPaper()
		}
	}

	// They picked paper.
	if codes[0] == 'B' {
		switch codes[1] {
		case 'X':
			return newRock()
		case 'Y':
			return newPaper()
		default:
			return newScissors()
		}
	}

	// They must have chosen 'C' - scissors.
	switch codes[1] {
	case 'X':
		return newPaper()
	case 'Y':
		return newScissors()
	default:
		return newRock()
	}
}

func newRock() choice {
	return choice{
		name: rock,
		val:  1,
	}
}

func newPaper() choice {
	return choice{
		name: paper,
		val:  2,
	}
}

func newScissors() choice {
	return choice{
		name: scissors,
		val:  3,
	}
}
