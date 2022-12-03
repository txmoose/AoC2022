package main

import (
	"bufio"
	"fmt"
	"os"
)

// A/X - Rock
// B/Y - Paper
// C/Z - Scissors
// A[X] = 4, A[Y] = 8, A[Z] = 3
// B[X] = 1, B[Y] = 5, B[Z] = 9
// C[X] = 7, C[Y] = 2, C[Z] = 6

func main() {
	input, err := os.Open("guide.txt")
	defer input.Close()
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	Scores := make(map[string]map[string]int)
	OppoRock := map[string]int{"X": 4, "Y": 8, "Z": 3}
	OppoPaper := map[string]int{"X": 1, "Y": 5, "Z": 9}
	OppoScissors := map[string]int{"X": 7, "Y": 2, "Z": 6}

	Scores["A"] = OppoRock
	Scores["B"] = OppoPaper
	Scores["C"] = OppoScissors

	TotalScore := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		TotalScore += Scores[string(line[0])][string(line[2])]
	}

	fmt.Printf("The total score is %d", TotalScore)
}
