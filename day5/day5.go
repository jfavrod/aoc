package main

import (
	"advent/day5/stack"
	"advent/input"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type stackFunc func(id int)

const STACK_COUNT = 9
const OFFSET = 4

func main() {
	in := input.GetInputForDay(5)
	stacks := make(map[int]*stack.Stack)

	iterateStacks(&stacks, func(id int) {
		stacks[id] = stack.NewStack()
	})

	loadInitialState(in, &stacks)

	iterateStacks(&stacks, func(id int) {
		fmt.Printf("%d: %v\n", id, *stacks[id])
	})

	in.Scan()
	fmt.Printf("line: %s\n", in.Text())
	exec(in, &stacks)

	iterateStacks(&stacks, func(id int) {
		fmt.Printf("%d: %v\n", id, *stacks[id])
	})
}

func exec(scanner *bufio.Scanner, stacks *map[int]*stack.Stack) {
	for scanner.Scan() {
		rx, err := regexp.Compile(`move (\d+) from (\d+) to (\d+)`)
		text := scanner.Text()
		matches := (*rx).FindStringSubmatch(text)

		count, err1 := strconv.Atoi(matches[1])
		from, err2 := strconv.Atoi(matches[2])
		to, err3 := strconv.Atoi(matches[3])

		if err != nil || err1 != nil || err2 != nil || err3 != nil {
			os.Exit(1)
		}

		stack.Move(count, (*stacks)[toStackId(from)], (*stacks)[toStackId(to)])
	}

}

func iterateStacks(stacks *map[int]*stack.Stack, fun stackFunc) {
	for i := 1; i < (OFFSET * STACK_COUNT); i += OFFSET {
		fun(i)
	}
}

func loadInitialState(scanner *bufio.Scanner, stacks *map[int]*stack.Stack) {
	for scanner.Scan() {
		text := []byte(scanner.Text())

		if text[0] != '[' {
			break
		}

		for loc, stk := range *stacks {
			if text[loc] != ' ' {
				stack.Add(stk, string(text[loc]))
			}
		}
	}

	iterateStacks(stacks, func(id int) {
		stack.Reverse((*stacks)[id])
	})
}

func toStackId(stackNum int) int {
	return ((stackNum * OFFSET) - OFFSET) + 1
}
