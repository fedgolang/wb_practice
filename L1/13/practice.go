package main

import "fmt"

func main() {
	var firstNum, secondNum int
	fmt.Scanf("%d %d", &firstNum, &secondNum)

	firstNum, secondNum = secondNum, firstNum

	fmt.Printf("first: %d\nsecond: %d", firstNum, secondNum)
}
