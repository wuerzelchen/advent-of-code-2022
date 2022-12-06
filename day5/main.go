package main

import (
	"container/list"
	"io/ioutil"
	"strconv"
	"strings"
)

func loadFile(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	return lines
}

type Stack []rune

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(value rune) {
	*s = append(*s, value) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return -1, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

type EvictionAlgo interface {
	evict(c *Cache, count int, fromColumn int, toColumn int)
}
type Fifo struct {
}

func (l *Fifo) evict(c *Cache, count int, fromColumn int, toColumn int) {
	toRemoveCount := count
	queue := list.New()

	if len(c.storage) == 0 {
		return
	}
	// get the values to move into a stack
	for i := len(c.storage) - 1; i >= 0; i-- {
		if len(c.storage[i]) == 0 && count == 0 {
			continue
		}
		if c.storage[i][fromColumn] != ' ' && len(c.storage[i]) > 0 {
			queue.PushBack(c.storage[i][fromColumn])
			count--
		}
		if count == 0 {
			break
		}
	}

	// remove the values from the fromColumn
	c.delete(toRemoveCount, fromColumn)
	// add the values to the toColumn
	for i := 0; i < (len(c.storage) + queue.Len()); i++ {
		if queue.Len() == 0 {
			break
		}
		lastRow := false
		if i == len(c.storage)-1 {
			lastRow = true
		}
		if lastRow && c.storage[i][toColumn] != ' ' {
			c.createAndAddEmptyRow()
			lastRow = false
		}
		if c.storage[i][toColumn] == ' ' {
			if rn, ok := queue.Front().Value.(rune); ok {
				q := queue.Front()
				c.storage[i][toColumn] = rune(rn)
				queue.Remove(q)
				if queue.Len() == 0 {
					break
				}
			} else {
			}
			if lastRow {
				c.createAndAddEmptyRow()
			}
		}
	}
}

type Lifo struct {
}

func (c *Cache) createAndAddEmptyRow() {
	row := make([]rune, len(c.storage[0]))
	for i := range row {
		row[i] = ' '
	}
	c.addRow(row)
}
func (l *Lifo) evict(c *Cache, count int, fromColumn int, toColumn int) {
	//create a stack array of integers
	toRemoveCount := count
	var stack Stack
	if len(c.storage) == 0 {
		return
	}
	// get the values to move into a stack
	for i := len(c.storage) - 1; i >= 0; i-- {
		if len(c.storage[i]) == 0 && count == 0 {
			continue
		}
		if c.storage[i][fromColumn] != ' ' && len(c.storage[i]) > 0 {
			stack.Push(c.storage[i][fromColumn])
			count--
		}
		if count == 0 {
			break
		}
	}
	// remove the values from the fromColumn
	c.delete(toRemoveCount, fromColumn)

	// add the values to the toColumn
	success := true
	for i := 0; i < (len(c.storage) + len(stack)); i++ {
		if len(stack) == 0 {
			break
		}
		lastRow := false
		if i == len(c.storage)-1 {
			lastRow = true
		}
		if lastRow && c.storage[i][toColumn] != ' ' {
			c.createAndAddEmptyRow()
			lastRow = false
		}
		if c.storage[i][toColumn] == ' ' {
			c.storage[i][toColumn], success = stack.Pop()
			if !success || len(stack) == 0 {
				break
			}
			if lastRow {
				c.createAndAddEmptyRow()
			}
		}
	}
}

type Cache struct {
	storage      [][]rune
	evictionAlgo EvictionAlgo
}

func initCache(e EvictionAlgo) *Cache {
	storage := make([][]rune, 0)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) addRow(value []rune) {
	c.storage = append(c.storage, value)
}

func (c *Cache) delete(count int, fromColumn int) {
	if len(c.storage) == 0 {
		return
	}
	for i := len(c.storage) - 1; i >= 0; i-- {
		if len(c.storage[i]) == 0 {
			continue
		}
		if c.storage[i][fromColumn] != ' ' {
			c.storage[i][fromColumn] = ' '
			count--
		}
		if count == 0 {
			break
		}
	}
}

func (c *Cache) evict(count int, fromColumn int, toColumn int) {
	c.evictionAlgo.evict(c, count, fromColumn, toColumn)
}

// move 2 from 8 to 1
// move 4 from 9 to 8
// move 2 from 1 to 6
// move 7 from 4 to 2
// move 10 from 2 to 7
func cleanUpStrings(input []string) []string {
	moves := make([]string, 0)
	for _, line := range input {
		line = strings.Replace(line, "move ", "", -1)
		line = strings.Replace(line, " from ", ",", -1)
		line = strings.Replace(line, " to ", ",", -1)
		if len(line) <= 6 && line != "" {
			moves = append(moves, line)
		}
	}
	return moves
}

func printTopOfCrate(c *Cache) {
	for i := 0; i < len(c.storage[0]); i++ {
		for j := len(c.storage) - 1; j >= 0; j-- {
			if c.storage[j][i] != ' ' {
				println(string(c.storage[j][i]))
				break
			}
		}
	}
	println("---")
}

func main() {
	fileName := "input.txt"
	fifo := &Fifo{}
	lifo := &Lifo{}
	lines := loadFile(fileName)
	moves := cleanUpStrings(lines)
	//save lines into rune matrix till a number is found
	crates := initCache(fifo)
	tmpCrates := make([][]rune, 0)
	matrixDone := false
	spaceCount := 0

	for _, line := range lines {
		if line == "" {
			break
		}
		row := make([]rune, 0)
		for _, char := range line {
			//if a number is found, continue
			if char >= '0' && char <= '9' {
				matrixDone = true
				spaceCount = 0
				break
				// else if not [ or ] then add to row
			} else if char != '[' && char != ']' && char != ' ' {
				row = append(row, char)
				spaceCount = 0
				//else if space, add to space count
			} else if char == ' ' {
				spaceCount++
				if spaceCount == 4 {
					row = append(row, char)
					spaceCount = 0
				}
			}
		}
		if len(row) > 0 {
			tmpCrates = append(tmpCrates, row)
		}
		if matrixDone {
			for i := len(tmpCrates) - 1; i >= 0; i-- {
				crates.addRow(tmpCrates[i])
			}
			break
		}
	}
	crates.setEvictionAlgo(lifo)
	// execute moves
	moveCounts := 0
	for _, move := range moves {
		splitMove := strings.Split(move, ",")
		if len(splitMove) > 0 {
			count, _ := strconv.Atoi(splitMove[0])
			fromColumn, _ := strconv.Atoi(splitMove[1])
			toColumn, _ := strconv.Atoi(splitMove[2])
			crates.evict(count, fromColumn-1, toColumn-1)
			println(move)
		}
		moveCounts++
	}
	println("---")
	println(moveCounts)
	println("---")
	printTopOfCrate(crates)
}
