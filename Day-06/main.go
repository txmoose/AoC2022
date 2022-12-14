package main

import (
	"bufio"
	"fmt"
	"os"
)

// borrowed from https://www.geeksforgeeks.org/how-to-remove-duplicate-values-from-slice-in-golang/
func removeDuplicateValues(intSlice []rune) []rune {
	keys := make(map[rune]bool)
	list := []rune{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// Evaluate 4 characters at a time, sliding window
func main() {
	var (
		fileLines   []string
		inputStream []rune
		pktStart    int
	)

	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	err = input.Close()
	if err != nil {
		panic(err)
	}

	inputStream = []rune(fileLines[0])

	for i := 4; i < len(inputStream); i += 1 {
		//fmt.Printf("Testing => %+v\n", inputStream[i-4:i])
		test := removeDuplicateValues(inputStream[i-4 : i])
		//fmt.Printf("Got back => %+v\n", test)
		if len(test) == 4 {
			fmt.Printf("The packet start is %d\n", i)
			pktStart = i
			break
		}
	}

	msgStream := inputStream[pktStart:]
	for j := 14; j < len(msgStream); j += 1 {
		msgTest := removeDuplicateValues(msgStream[j-14 : j])
		if len(msgTest) == 14 {
			fmt.Printf("The message start is %d\n", j+pktStart)
			break
		}
	}

}
