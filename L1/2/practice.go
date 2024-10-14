package main

import (
	"fmt"
	"sync"
)

func squareWorker(num int, wg *sync.WaitGroup, resultChan chan<- int) {
	defer wg.Done()
	result := num * num
	resultChan <- result
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	resultChan := make(chan int, len(numbers))

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go squareWorker(num, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		fmt.Println(result)
	}
}
