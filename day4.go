package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// 1-6,2-5 = contained
// 1-5,3-6 = not contained
// 2-3,1-6 = contained
func isContained(a1 int, a2 int, b1 int, b2 int) bool {
	if a1 >= b1 && a2 <= b2 {
		return true
	} else if a1 <= b1 && a2 >= b2 {
		return true
	} else {
		return false
	}
}

// 1-5,3-7 = overlapping
// 1-3,4-6 = not overlapping
// 1-3,3-8 = overlapping
// 8-9,2-9 = overlapping
// 5-9,6-8 = overlapping
// 1-6,2-4 = overlapping
// 7-8,1-3 = not overlapping
// 1-7,4-9 = overlapping
func isOverlapping(a1 int, a2 int, b1 int, b2 int) bool {
	if a1 > b2 || a2 < b1 {
		return false
	} else {
		return true
	}
}

func day4() {
	sum := 0
	overlappingSum := 0
	//read contents of file from disk
	data, err := ioutil.ReadFile("day4input.txt")
	if err != nil {
		panic(err)
	}
	//read data line by line
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		_ = i
		//check if line is empty
		if line == "" {
			continue
		}
		//split line into sections
		sections := strings.Split(line, ",")
		//split first section into two numbers
		section1 := strings.Split(sections[0], "-")
		//convert first number to int
		a1 := 0
		fmt.Sscanf(section1[0], "%d", &a1)
		//convert second number to int
		a2 := 0
		fmt.Sscanf(section1[1], "%d", &a2)
		//split second section into two numbers
		section2 := strings.Split(sections[1], "-")
		//convert first number to int
		b1 := 0
		fmt.Sscanf(section2[0], "%d", &b1)
		//convert second number to int
		b2 := 0
		fmt.Sscanf(section2[1], "%d", &b2)
		//check if a is contained in b
		if isContained(a1, a2, b1, b2) {
			sum++
		}
		if isOverlapping(a1, a2, b1, b2) {
			overlappingSum++
		}

	}
	fmt.Println("containing sum: ", sum)
	fmt.Println("overlapping sum: ", overlappingSum)
}
