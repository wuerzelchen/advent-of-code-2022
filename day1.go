package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Stack []int

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(value int) {
	*s = append(*s, value) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func day1() {
	var stack Stack
	var maxValue int
	// load file from disk
	data, err := ioutil.ReadFile("day1input.txt")
	if err != nil {
		panic(err)
	}
	var sum int
	sum = 0
	// read data line by line
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		_ = i
		// check if line is empty
		if line == "" {

			for len(stack) > 0 {
				x, y := stack.Pop()
				if y {
					sum += x
				}
			}
			if condition := sum > maxValue; condition {
				maxValue = sum
			}
			sum = 0
			continue
		}
		// convert string to int
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		stack = append(stack, num)
	}
	fmt.Println(maxValue)
}
