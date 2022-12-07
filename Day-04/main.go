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
		leftBegin, leftEnd, LeftLength, err := GetLength(leftRange)
		if err != nil {
			panic(err)
		}
		righBegin, righEnd, RighLength, err := GetLength(rightRange)
		if err != nil {
			panic(err)
		}
		if leftBegin <= righBegin && righBegin <= leftEnd {
			runningCount += 1
			continue
		} else if leftBegin <= righEnd && righEnd <= leftEnd {
			runningCount += 1
			continue
		} else if righBegin <= leftBegin && leftBegin <= righEnd {
			runningCount += 1
			continue
		} else if righBegin <= leftEnd && leftEnd <= righEnd {
			runningCount += 1
			continue
		}
		fmt.Printf("%d %d\n", LeftLength, RighLength)
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
