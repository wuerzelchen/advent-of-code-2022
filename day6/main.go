package main

import (
	"io/ioutil"
	"strings"
)

const messageLength = 14

func loadFile(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	return lines
}

// bvwbjplbgvbhsrlpgdmjqwftvncz = first marker after character 5
// nppdvjthqldpwncqszvftbrmjlhg = first marker after character 6
// nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg = first marker after character 10
// zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw = first marker after character 11
// find the first marker after 4 unequal characters
func findMarker(line string) int {
	count := messageLength
	returnValue := 0
	currentPacket := ""
	for i := 0; i < len(line); i++ {
		if !strings.Contains(currentPacket, string(line[i])) {
			count--
			currentPacket = currentPacket + string(line[i])

		} else {
			i -= (messageLength - count)
			count = messageLength
			currentPacket = ""
		}
		if count == 0 {
			returnValue = i + 1
			break
		}
	}
	return returnValue
}

func main() {
	filename := "input.txt"
	lines := loadFile(filename)
	for _, line := range lines {
		if len(line) > 0 {
			println("first marker after character: ", findMarker(line))
		}
	}
}
