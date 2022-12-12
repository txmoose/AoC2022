package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("stacks.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	var (
		stacks        = make(map[int][]string)
		fileLines     []string
		tops          []string
		upDownPtr     int
		MovesStartPtr = 0
		stackPtr      = 1
	)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	input.Close()

	// find the empty line that denotes the bottom of the stacks
	for i, v := range fileLines {
		if v == "" {
			upDownPtr = i - 1
			MovesStartPtr = i + 1
			break
		}
	}

	// populate the stacks
	for _, v := range strings.Fields(fileLines[upDownPtr]) {
		name, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		stacks[name] = nil
	}
	upDownPtr -= 1

	for ; upDownPtr >= 0; upDownPtr -= 1 {
		for i := 1; i < len(fileLines[upDownPtr]); i += 4 {
			line := []rune(fileLines[upDownPtr])
			if string(line[i]) == " " {
				stackPtr += 1
				continue
			}
			stacks[stackPtr] = append(stacks[stackPtr], string(line[i]))
			stackPtr += 1
		}
		stackPtr = 1
	}

	// finally start doing the push/pops
	//    move 1 from 2 to 1
	// strings.Fields()
	// 1 - count for the for loop
	// 3 - pop
	// 5 - push
	moveLines := fileLines[MovesStartPtr:]
	for i := 0; i < len(moveLines); i += 1 {
		move := strings.Fields(moveLines[i])
		numMoves, err := strconv.Atoi(move[1])
		if err != nil {
			panic(err)
		}
		popStack, err := strconv.Atoi(move[3])
		if err != nil {
			panic(err)
		}
		pushStack, err := strconv.Atoi(move[5])
		if err != nil {
			panic(err)
		}

		n := len(stacks[popStack]) - numMoves
		hold := stacks[popStack][n:]
		stacks[popStack] = stacks[popStack][:n]

		for x := 0; x < len(hold); x += 1 {
			stacks[pushStack] = append(stacks[pushStack], hold[x])
		}
	}

	for i := 1; i <= len(stacks); i += 1 {
		fmt.Printf("%d => %v\n", i, stacks[i])
		n := len(stacks[i]) - 1
		tops = append(tops, stacks[i][n])
	}
	for i := 0; i < len(tops); i += 1 {
		fmt.Printf("%s", tops[i])
	}
}
