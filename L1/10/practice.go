package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func splitter(s string) []string {
	splitString := strings.Split(s, ", ")
	return splitString
}

func main() {

	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		fmt.Printf("Read error: %v", err)
		return
	}

	splited := splitter(strings.Trim(line, "\r\n"))

	temp := make(map[int][]float64)
	for _, split := range splited {
		numSplit, err := strconv.ParseFloat(split, 64)
		if err != nil {
			fmt.Printf("Convert to float error: %v", err)
			return
		}

		group := 0
		if numSplit < 0 {
			group = int(math.Ceil(numSplit/10) * 10)
		} else {
			group = int(math.Floor(numSplit/10) * 10)
		}

		temp[group] = append(temp[group], numSplit)
	}
	fmt.Printf("Groups: \n%v", temp)
}
