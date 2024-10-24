package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		count := 1
		for {
			ch <- fmt.Sprintf("msg %d", count)
			time.Sleep(time.Second)
			count++
		}
	}()

	for msg := range ch {
		fmt.Printf("Message: %s\n", msg)
	}
}
