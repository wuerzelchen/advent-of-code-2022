package main

import (
	"io/ioutil"
	"strings"
)

// a through z = 1 through 26
// A through Z = 27 through 52
func getPriorityValue(char rune) int {
	if char >= 97 && char <= 122 {
		return int(char - 96)
	} else if char >= 65 && char <= 90 {
		return int(char - 38)
	} else {
		return 0
	}
}

func day3() {
	sum := 0
	//read contents of file from disk
	data, err := ioutil.ReadFile("day3input.txt")
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
		//split line in half
		half := len(line) / 2
		//split line into two halves
		first_half := line[:half]
		second_half := line[half:]
		//check which character is in both halves
		for j, char := range first_half {
			_ = j
			//check if character is in second half
			if strings.Contains(second_half, string(char)) {
				//if character is in both halves, add it to the sum
				sum += getPriorityValue(char)
				break
			}
		}
	}
	println(sum)
}
