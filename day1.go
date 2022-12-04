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

// remove element from slice
func remove(slice []int, s int) []int {
	var p int
	// find index of element
	for i := 0; i < len(slice); i++ {
		if slice[i] == s {
			// Remove the element at index i from a.
			slice[i] = slice[len(slice)-1] // Copy last element to index i.
			slice[len(slice)-1] = 0        // Erase last element (write zero value).
			slice = slice[:len(slice)-1]   // Truncate slice.
			p = i
			break
		}
	}
	// return slice without element
	return slice[:p+copy(slice[p:], slice[p+1:])]
}

func day1() {
	var stack Stack
	var maxValue int
	var top3 []int
	var sumTop3 int
	// load file from disk
	data, err := ioutil.ReadFile("day1input.txt")
	if err != nil {
		panic(err)
	}
	var sum int
	sumList := make([]int, 0)
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
					sumList = append(sumList, sum)
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
	//get top3 elements from sumList
	for i := 0; i < 3; i++ {
		max := 0
		for _, v := range sumList {
			if v > max {
				max = v
			}
		}
		top3 = append(top3, max)
		sumList = remove(sumList, max)
	}

	fmt.Println("Top 3 sum values are: ", top3)
	//summarize all values in top3
	for _, v := range top3 {
		sumTop3 += v
	}

	fmt.Println("maxiumum calories: ", maxValue)
	fmt.Println("top three most calories sum: ", sumTop3)
}
