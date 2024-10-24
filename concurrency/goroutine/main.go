package main

import (
	"fmt"
	"time"
)

func printer(msg string) {
	fmt.Println(msg)
}

func main() {
	fmt.Println("Init")
	go printer("hello concurrency")
	fmt.Println("End")
	time.Sleep(1 * time.Second)
}
