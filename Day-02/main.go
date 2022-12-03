package main

import (
	"bufio"
	"fmt"
	"os"
)

// A - Rock     - X = Scissors, Y = Rock, Z = Paper
// B - Paper    - X = Rock, Y = Paper, Z = Scissors
// C - Scissors - X = Paper, Y = Scissors, Z = Rock
// A[X] = 3, A[Y] = 4, A[Z] = 8
// B[X] = 1, B[Y] = 5, B[Z] = 9
// C[X] = 2, C[Y] = 6, C[Z] = 7

func main() {
	input, err := os.Open("guide.txt")
	defer input.Close()
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	Scores := make(map[string]map[string]int)
	OppoRock := map[string]int{"X": 3, "Y": 4, "Z": 8}
	OppoPaper := map[string]int{"X": 1, "Y": 5, "Z": 9}
	OppoScissors := map[string]int{"X": 2, "Y": 6, "Z": 7}

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
