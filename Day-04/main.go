package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("sections.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)
	runningCount := 0

	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), ",")
		//fmt.Printf("%s || %s\n", line[0], line[1])
		leftRange := strings.Split(line[0], "-")
		rightRange := strings.Split(line[1], "-")
		leftBegin, leftEnd, leftLength, err := GetLength(leftRange)
		if err != nil {
			panic(err)
		}
		righBegin, righEnd, righLength, err := GetLength(rightRange)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Checking %s || %s\n", line[0], line[1])
		if righLength > leftLength {
			if righBegin <= leftBegin && righEnd >= leftEnd {
				//fmt.Println(" => Right Contains Left")
				runningCount += 1
				continue
			}
		} else if righLength < leftLength {
			if righBegin >= leftBegin && righEnd <= leftEnd {
				//fmt.Println(" => Left Contains Right")
				runningCount += 1
				continue
			}
		} else if righLength == leftLength {
			if righBegin == leftBegin && righEnd == leftEnd {
				//fmt.Println(" => Ranges are the same")
				runningCount += 1
				continue
			}
		}
	}
	fmt.Printf("runningCount => %d", runningCount)
}

func GetLength(r []string) (b, e, l int, err error) {
	b, er := strconv.Atoi(r[0])
	if er != nil {
		return 0, 0, 0, er
	}
	e, er = strconv.Atoi(r[1])
	if er != nil {
		return 0, 0, 0, er
	}
	l = e - b + 1
	return b, e, l, nil
}
