package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		count := 1
		for {
			ch1 <- fmt.Sprintf("Data from ch1 %d", count)
			time.Sleep(time.Second)
			count++
		}
	}()

	go func() {
		count := 1
		for {
			ch2 <- fmt.Sprintf("Data from ch2 %d", count)
			time.Sleep(time.Second * 2)
			count++
		}
	}()

	for {
		select {
		case msg := <-ch1:
			fmt.Println("Received from ch1: ", msg)
		case msg := <-ch2:
			fmt.Println("Received from ch2: ", msg)
		}
	}
}
