package main

import (
	"fmt"
	"time"
)

func doSomething(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing something...")
		}
	}
}

func main() {
	done := make(chan bool)

	go doSomething(done)

	time.Sleep(time.Second * 3)
	close(done)
}
