package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	Position int
	Calories int
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	var Elves []Elf
	position := 0
	CurrentElf := Elf{position, 0}

	for fileScanner.Scan() {
		if len(fileScanner.Text()) > 0 {
			snack, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				panic(err)
			}
			CurrentElf.Calories += snack
		} else {
			Elves = append(Elves, CurrentElf)
			position += 1
			CurrentElf.Calories = 0
		}
	}
	Elves = append(Elves, CurrentElf)

	sort.Slice(Elves[:], func(i, j int) bool {
		return Elves[i].Calories > Elves[j].Calories
	})

	fmt.Println(Elves[:3])
	fmt.Println(Elves[0].Calories + Elves[1].Calories + Elves[2].Calories)

	input.Close()
}
