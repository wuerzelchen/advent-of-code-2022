package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// a && y = x
// b && x = x
// c && z = x
// a && z = y
// b && y = y
func getOutcome(opponent_move string, my_move string) string {
	if opponent_move == "a" && my_move == "y" {
		return "x"
	} else if opponent_move == "b" && my_move == "x" {
		return "x"
	} else if opponent_move == "c" && my_move == "z" {
		return "x"
	} else if opponent_move == "a" && my_move == "z" {
		return "y"
	} else if opponent_move == "b" && my_move == "y" {
		return "y"
	} else if opponent_move == "c" && my_move == "x" {
		return "y"
	} else if opponent_move == "a" && my_move == "x" {
		return "z"
	} else if opponent_move == "b" && my_move == "z" {
		return "z"
	} else if opponent_move == "c" && my_move == "y" {
		return "z"
	} else {
		return "error"
	}
}

// a = rock
// b = paper
// c = scissors
// x = rock
// y = paper
// z = scissors
func checkWinCondition(opponent_move string, my_move string) bool {
	if opponent_move == "a" && my_move == "y" {
		return true
	} else if opponent_move == "b" && my_move == "z" {
		return true
	} else if opponent_move == "c" && my_move == "x" {
		return true
	} else {
		return false
	}
}

// a = rock
// b = paper
// c = scissors
// x = rock
// y = paper
// z = scissors
func checkDrawCondition(opponent_move string, my_move string) bool {
	if opponent_move == "a" && my_move == "x" {
		return true
	} else if opponent_move == "b" && my_move == "y" {
		return true
	} else if opponent_move == "c" && my_move == "z" {
		return true
	} else {
		return false
	}
}

func day2() {
	score := 0
	// read contents of file from disk
	data, err := ioutil.ReadFile("day2input.txt")
	if err != nil {
		panic(err)
	}
	// read data line by line
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		_ = i
		//try catch
		if line == "" {
			continue
		}
		opponent_move := strings.Split(line, " ")[0]
		my_move := strings.Split(line, " ")[1]
		opponent_move = strings.ToLower(opponent_move)
		my_move = strings.ToLower(my_move)
		if checkWinCondition(opponent_move, getOutcome(opponent_move, my_move)) {
			score += 6
		} else if checkDrawCondition(opponent_move, getOutcome(opponent_move, my_move)) {
			score += 3
		} else {
			score += 0
		}
		my_move = getOutcome(opponent_move, my_move)
		if my_move == "x" {
			score += 1
		} else if my_move == "y" {
			score += 2
		} else if my_move == "z" {
			score += 3
		}
	}
	fmt.Println("total score: ", score)

}
