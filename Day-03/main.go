package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, err := os.Open("rucks.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	Found := false
	TotalPriority := 0
	PriorityScores := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5,
		"f": 6, "g": 7, "h": 8, "i": 9, "j": 10,
		"k": 11, "l": 12, "m": 13, "n": 14, "o": 15,
		"p": 16, "q": 17, "r": 18, "s": 19, "t": 20,
		"u": 21, "v": 22, "w": 23, "x": 24, "y": 25,
		"z": 26, "A": 27, "B": 28, "C": 29, "D": 30,
		"E": 31, "F": 32, "G": 33, "H": 34, "I": 35,
		"J": 36, "K": 37, "L": 38, "M": 39, "N": 40,
		"O": 41, "P": 42, "Q": 43, "R": 44, "S": 45,
		"T": 46, "U": 47, "V": 48, "W": 49, "X": 50,
		"Y": 51, "Z": 52,
	}

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	input.Close()

	linePointer := 0
	for linePointer < len(fileLines) {
		first := fileLines[linePointer]
		linePointer += 1
		second := fileLines[linePointer]
		linePointer += 1
		third := fileLines[linePointer]
		linePointer += 1

		for _, v := range first {
			for _, x := range second {
				if v == x {
					for _, z := range third {
						if v == z {
							TotalPriority += PriorityScores[string(v)]
							Found = true
							break
						}
					}
				}
				if Found {
					break
				}
			}
			if Found {
				Found = false
				break
			}
		}
	}
	fmt.Printf("The total priority is %v", TotalPriority)
}
