package main

import (
	"fmt"
	"time"
)

func main() {
	sec := 2
	ch := make(chan int)
	q := make(chan bool)

	go func() {
		i := 0
		for {
			select {
			case <-q:
				fmt.Println("Stopping sender!")
				return
			case ch <- i:
				i++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	go func() {
		for val := range ch {
			fmt.Println("Received:", val)
		}
	}()

	time.AfterFunc(time.Duration(sec)*time.Second, func() {
		q <- true
		close(ch)
		fmt.Println("Time is up!")
	})

	time.Sleep(time.Duration(sec+1) * time.Second)
	fmt.Println("Program finished")
}
